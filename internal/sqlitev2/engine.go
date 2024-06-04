package sqlitev2

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	restart2 "github.com/vkcom/statshouse/internal/sqlitev2/checkpoint"
	"github.com/vkcom/statshouse/internal/sqlitev2/waitpool"
	"github.com/vkcom/statshouse/internal/vkgo/binlog"
	"go.uber.org/multierr"
	"pgregory.net/rand"
)

/*
TODO
-- !!! Прерывание долгих транзакций
- Создание индекса атомарная операция?
- Если упадет во время долгого бэкапа, то все сделанные записи за это время будут откачены. Решение: если писать оффсет в рестарт файл каждый раз при коммите, то при рестарте оба вала выживут
- Надо следить когда завершаются рид транзакции и если была завершена последняя из тех которые держит вал, то делать чекпоинт
- "PRAGMA journal_size_limit". поигаться с конфигурацией
- Убивать долгие read транзакции чтобы избежать разростания вал файла (или хотя бы писать метрику на такие)
- WAL switch лучше контролировать самому, чтобы сразу после смены файла инициировать коммит барсика

- sqlite3_db_cacheflush

- Репортить в метрику размер вала

- Чекпоинт из другой транзакции

- https://www.sqlite.org/lang_attach.html - использовать чтобы получить стандартные транзакционные гарантии

NOTES:
- Если стартуем с бэкапа у него должно быть тоже имя что у основной базы, при этом надо не забыть потереть wal и wal2 файл. Поэтому последний бэкап рекомендуется хранить рядом на том же диске
- Гарантии на транзакции нет. Все операции могут быть откачены.
- В некоторых крайних случаях откатываться может бесконечно, поэтому крайне не рекомендуется использовать без бинлога. Если требуется сделать изменения без бинлога:
 1. Сделать изменения после OpenEngine
 2. Вызвать Run

- Не рекомендуются долгие Read транзакции
*/
type (
	Engine struct {
		opt               Options
		rw                *sqliteBinlogConn
		binlog            binlog.Binlog
		userEngine        UserEngine
		binlogEngine      *binlogEngine
		finishBinlogRunCh chan struct{}
		readyCh           chan error
		readOnly          bool
		re                *restart2.RestartFile

		roConnPool  *connPool
		readyNotify sync.Once

		//testOptions *testOptions

		logMx       sync.Mutex
		logger      *log.Logger
		nextLogTime time.Time

		waitDbOffsetPool *waitpool.WaitPool
	}

	Options struct {
		// Path to db file
		Path string

		// Use this to specify your SQLITE db format
		APPID uint32

		// User table scheme
		Scheme string

		// Open db in readonly mode. Don't use binlog in this mode
		ReadOnly bool

		// ReadOnly connection pool max size
		MaxROConn int

		// Prepared statement's cache max size (soft)
		CacheApproxMaxSizePerConnect int

		// SQLite page size (fill 0 to use default)
		PageSize int

		StatsOptions StatsOptions

		BinlogOptions

		IntegrityCheckBeforeStart bool
		// Advanced RO mode, DONT USE
		notUseWALROMMode bool
	}
	BinlogOptions struct {
		// Set true if binlog created in replica mode
		Replica bool
	}

	ApplyEventFunction func(conn Conn, payload []byte) (int, error)

	ViewTxOptions struct {
		QueryName  string
		WaitOffset int64
	}

	DoTxResult struct {
		DBOffset int64 // sqlite snapshot offset, after applying current Do
	}
	ViewTxResult struct {
		DBOffset int64 // sqlite snapshot offset for current View transaction
	}
)

const (
	initOffsetTable       = "CREATE TABLE IF NOT EXISTS __binlog_offset (offset INTEGER);"
	initCommitOffsetTable = "CREATE TABLE IF NOT EXISTS __binlog_commit_offset (offset INTEGER);"
	snapshotMetaTable     = "CREATE TABLE IF NOT EXISTS __snapshot_meta (meta BLOB);"
	internalQueryPrefix   = "__"
	logPrefix             = "[sqlite-engine]"
	debugFlag             = true
)

func openRO(opt Options) (*Engine, error) {
	logger := log.New(os.Stdout, logPrefix, log.LstdFlags)
	e := &Engine{
		opt:      opt,
		readOnly: true,
		roConnPool: newConnPool(opt.MaxROConn, func() (*sqliteConn, error) {
			if !opt.notUseWALROMMode {
				return newSqliteROWALConn(opt.Path, opt.CacheApproxMaxSizePerConnect, opt.StatsOptions, logger)
			} else {
				return newSqliteROConn(opt.Path, opt.StatsOptions, logger)
			}
		}, logger),
		logger:           logger,
		waitDbOffsetPool: waitpool.NewPool(),
	}
	return e, nil
}

/*
TODO
Функция хелпер которая позволяет протестировать соотносимость ду операций и apply операций. Ождиается что в конце будет сравниваться 2 таблицы
*/
// func Helper(scheme string, do func(c Conn, cache []byte) ([]byte, error), apply ApplyEventFunction) error

/*
OpenEngine open or create SQLite db file.

	engine := OpenEngine(...)
	can use engine as sqlite wrapper

	go engine.Run(...)
	can use View, can't use Do

	err := <-engine.ReadyCh()
	can use engine as sqlite + binlog wrapper
*/
func OpenEngine(opt Options) (*Engine, error) {
	if opt.ReadOnly {
		return openRO(opt)
	}
	logger := log.New(os.Stdout, logPrefix, log.LstdFlags)
	re, err := restart2.OpenAndLock(opt.Path)
	if err != nil {
		return nil, fmt.Errorf("failed to open runRestart file: %w", err)
	}
	logger.Println("Running runRestart script")
	err = runRestart(re, opt, logger)
	if err != nil {
		return nil, multierr.Append(err, re.Close())
	}
	stat, _ := os.Stat(opt.Path)
	var size int64
	if stat != nil {
		size = stat.Size()
	}
	waitDbOffsetPool := waitpool.NewPool()
	logger.Printf("OPEN DB path: %s size(only db file): %d", opt.Path, size)
	rw, err := newSqliteBinlogConn(opt.Path, opt.APPID, opt.CacheApproxMaxSizePerConnect, opt.PageSize, opt.Replica, opt.StatsOptions, waitDbOffsetPool, logger)
	if err != nil {
		return nil, multierr.Append(err, re.Close())
	}

	if opt.IntegrityCheckBeforeStart {
		err = rw.conn.integrityCheck()
		if err != nil {
			err = multierr.Append(err, re.Close())
			errClose := rw.Close()
			return nil, fmt.Errorf("failed to intergrity check: %w", multierr.Append(err, errClose))
		}
		logger.Println("integrity check: ok")
	}
	err = rw.conn.applyScheme(initOffsetTable, snapshotMetaTable, initCommitOffsetTable, opt.Scheme)
	if err != nil {
		err = multierr.Append(err, re.Close())
		errClose := rw.Close()
		return nil, fmt.Errorf("failed to apply acheme: %w", multierr.Append(err, errClose))
	}

	e := &Engine{
		opt:               opt,
		rw:                rw,
		re:                re,
		finishBinlogRunCh: make(chan struct{}),
		readyCh:           make(chan error, 1),
		roConnPool: newConnPool(opt.MaxROConn, func() (*sqliteConn, error) {
			return newSqliteROWALConn(opt.Path, opt.CacheApproxMaxSizePerConnect, opt.StatsOptions, logger)
		}, logger),
		logger:           logger,
		waitDbOffsetPool: waitDbOffsetPool,
	}

	dbOffset, err := e.binlogLoadOrCreatePosition()
	e.rw.setDBOffsetLocked(dbOffset)
	e.waitDbOffsetPool.Notify(e.rw.getDBOffsetLocked())
	if err != nil {
		err = fmt.Errorf("failed to load binlog position during to start run: %w", err)
		e.readyNotify.Do(func() {
			e.readyCh <- err
			close(e.readyCh)
		})
		return nil, multierr.Append(err, re.Close())

	}
	e.logger.Printf("load binlog position: %d", e.rw.getDBOffsetLocked())

	_, err = e.binlogLoadOrCreateMeta()
	if err != nil {
		err = fmt.Errorf("failed to load binlog meta durint to start run: %w", err)
		return nil, multierr.Append(err, re.Close())
	}

	err = e.rw.enableWALSwitchCallbackLocked()
	if err != nil {
		err = fmt.Errorf("failed to set wal switch callback: %w", err)
		return nil, multierr.Append(err, re.Close())
	}
	return e, nil
}

/*
binlog - will be closed during to Engin	e.Close
*/
func (e *Engine) Run(binlog binlog.Binlog, userEngine UserEngine, applyEventFunction ApplyEventFunction) (err error) {
	if e.readOnly {
		return fmt.Errorf("can't use binlog in readonly mode")
	}
	e.rw.mu.Lock()
	e.rw.registerWALSwitchCallbackLocked(e.switchCallBack)
	e.binlog = binlog
	e.userEngine = userEngine
	e.binlogEngine = newBinlogEngine(e, applyEventFunction)
	e.rw.mu.Unlock()
	defer func() { close(e.finishBinlogRunCh) }()

	go e.binlogEngine.RunCheckpointer()
	defer e.binlogEngine.StopCheckpointer()
	meta, err := e.binlogLoadOrCreateMeta()
	if err != nil {
		err = fmt.Errorf("failed to load binlog meta durint to start run: %w", err)
		e.readyNotify.Do(func() {
			e.readyCh <- err
			close(e.readyCh)
		})
		return err
	}
	e.logger.Printf("load snapshot meta: %s", hex.EncodeToString(meta))

	e.logger.Printf("running binlog")
	err = e.rw.setError(e.binlog.Run2(e.rw.getDBOffsetLocked(), meta, meta, false, e.binlogEngine))
	e.rw.mu.Lock()
	defer e.rw.mu.Unlock()
	e.binlog = nil
	if err != nil {
		e.readyNotify.Do(func() {
			e.readyCh <- err
			close(e.readyCh)
		})
		return err
	}
	e.readyNotify.Do(func() {
		close(e.readyCh)
	})
	return nil
}

func (e *Engine) ReadyCh() <-chan error {
	return e.readyCh
}

func (e *Engine) switchCallBack(iApp int, maxFrame uint) {
	e.opt.StatsOptions.walSwitchSize(iApp, maxFrame)
	e.binlogEngine.checkpointer.setWaitCheckpointOffsetLocked()
}

// TODO better interface?
func (e *Engine) Backup(ctx context.Context, prefix string) (string, int64, error) {
	if prefix == "" {
		return "", 0, fmt.Errorf("backup prefix is Empty")
	}
	e.logger.Printf("starting backup")
	startTime := time.Now()
	defer e.opt.StatsOptions.measureActionDurationSince("backup", startTime)
	conn, err := newSqliteROWALConn(e.opt.Path, 10, e.opt.StatsOptions, e.logger)
	if err != nil {
		return "", 0, fmt.Errorf("failed to open RO connection to backup: %w", err)
	}
	defer func() {
		_ = conn.Close()
	}()
	c := newUserConn(conn, ctx)
	path := prefix + "." + strconv.FormatUint(rand.Uint64(), 10) + ".tmp"
	defer func() {
		_ = os.Remove(path)
	}()
	if _, err := os.Stat(path); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err := c.Exec("__vacuum", "VACUUM INTO $to", TextString("$to", path))
			if err != nil {
				return path, 0, err
			}
		} else {
			return path, 0, fmt.Errorf("os.Stats failed: %w", err)
		}
	}

	conn1, err := newSqliteROConn(path, e.opt.StatsOptions, e.logger)
	if err != nil {
		return "", 0, fmt.Errorf("failed to open RO connection to rename backup: %w", err)
	}
	defer func() {
		_ = conn1.Close()
	}()
	c1 := newInternalConn(conn1)
	expectedPath, binlogPos, err := getBackupPath(c1, prefix)
	if err != nil {
		return "", 0, fmt.Errorf("failed to get backup path: %w", err)
	}
	e.logger.Printf("backup is loaded to temp file. Starting wait binlog commit")
	// TODO при реверте без рестарта требуется таймаут
	e.binlogEngine.binlogWait(binlogPos, true)
	stat, _ := os.Stat(path)
	e.logger.Printf("finish backup successfully in %f seconds, path: %s, pos: %d, size: %d", time.Since(startTime).Seconds(), expectedPath, binlogPos, stat.Size())
	return expectedPath, binlogPos, os.Rename(path, expectedPath)
}

func getBackupPath(conn internalConn, prefix string) (string, int64, error) {
	pos, isExists, err := binlogLoadPosition(conn)
	if err != nil {
		return "", pos, fmt.Errorf("failed to load binlog position from backup: %w", err)
	}
	if !isExists {
		return "", pos, fmt.Errorf("failed to load binlog position: db is Empty")
	}

	copyPos := pos
	numLen := -4
	for copyPos > 0 {
		numLen++
		copyPos /= 10
	}
	if numLen < 0 {
		numLen = 0
	}

	posStr := fmt.Sprintf(`%04d`, pos)
	prefix = fmt.Sprintf(`%s.%02d`, prefix, numLen)

	for l := 4; l <= len(posStr); l++ {
		filename := prefix + posStr[:l]
		if _, err = os.Stat(filename); os.IsNotExist(err) {
			return filename, pos, nil
		}
	}

	return "", pos, fmt.Errorf("can not create backup with pos=%d, probably backup already exist", pos)
}

func (e *Engine) View(ctx context.Context, queryName string, fn func(Conn) error) (res ViewTxResult, err error) {
	return e.ViewOpts(ctx, ViewTxOptions{QueryName: queryName}, fn)
}

func (e *Engine) ViewOpts(ctx context.Context, opt ViewTxOptions, fn func(Conn) error) (res ViewTxResult, err error) {
	if err = checkUserQueryName(opt.QueryName); err != nil {
		return res, err
	}
	err = e.waitDbOffsetPool.Wait(ctx, opt.WaitOffset)
	if err != nil {
		return res, err
	}
	startTimeBeforeLock := time.Now()
	conn, err := e.roConnPool.Get()
	if err != nil {
		return res, fmt.Errorf("faield to get RO conn: %w", err)
	}
	defer e.roConnPool.Put(conn)

	e.opt.StatsOptions.measureWaitDurationSince(waitView, startTimeBeforeLock)
	defer e.opt.StatsOptions.measureSqliteTxDurationSince(txView, opt.QueryName, time.Now())
	err = conn.beginTxLocked()
	if err != nil {
		return res, fmt.Errorf("failed to begin RO tx: %w", err)
	}
	defer func() {
		errRollback := conn.rollbackLocked()
		if errRollback != nil {
			err = multierr.Append(err, errRollback)
		}
	}()
	c := newUserConn(conn, ctx)
	offset, _, err := binlogLoadPosition(internalFromUser(c))
	if err != nil {
		return res, fmt.Errorf("read binlog position error: %w", err)
	}
	res.DBOffset = offset
	err = fn(c)
	if err != nil {
		return res, fmt.Errorf("user error: %w", err)
	}
	err = conn.commitTxLocked()
	if err != nil {
		return res, fmt.Errorf("failed to commit RO tx: %w", err)
	}
	return res, err
}

func (e *Engine) DoTx(ctx context.Context, queryName string, do func(c Conn, cache []byte) ([]byte, error)) (res DoTxResult, err error) {
	if err := checkUserQueryName(queryName); err != nil {
		return res, err
	}
	startTimeBeforeLock := time.Now()
	e.rw.mu.Lock()
	defer e.rw.mu.Unlock()
	e.opt.StatsOptions.measureWaitDurationSince(waitDo, startTimeBeforeLock)
	defer e.opt.StatsOptions.measureSqliteTxDurationSince(txDo, queryName, time.Now())
	if e.readOnly || e.rw.isReplica {
		return res, errReadOnly
	}
	err = e.rw.beginTxLocked()
	if err != nil {
		e.opt.StatsOptions.engineBrokenEvent()
		return res, fmt.Errorf("failed to begin tx: %w", err)
	}
	defer func() {
		errRollback := e.rw.rollbackLocked()
		if errRollback != nil {
			err = multierr.Append(err, errRollback)
		}
	}()
	conn := newUserConn(e.rw.conn, ctx)
	bytes, err := do(conn, e.rw.binlogCache[:0])
	if err != nil {
		return res, fmt.Errorf("user error: %w", err)
	}
	if len(bytes) == 0 {
		if e.binlog != nil {
			return res, fmt.Errorf("do without binlog event")
		}
		return res, e.rw.nonBinlogCommitTxLocked()
	}
	if e.binlog == nil {
		return res, fmt.Errorf("can't write binlog event: binlog is nil")
	}
	offsetAfterWrite, err := e.binlog.Append(e.rw.getDBOffsetLocked(), bytes)
	if err != nil {
		return res, fmt.Errorf("binlog Append return error: %w", err)
	}
	err = e.rw.binlogCommitTxLocked(offsetAfterWrite)
	return DoTxResult{
		DBOffset: offsetAfterWrite,
	}, err
}

// В случае возникновения ошибки движок считается сломаным
func (e *Engine) internalDoBinlog(queryName string, do func(c internalConn) (int64, error)) error {
	if err := checkInternalQueryName(queryName); err != nil {
		return err
	}
	startTimeBeforeLock := time.Now()
	e.rw.mu.Lock()
	defer e.rw.mu.Unlock()
	defer func() {
		err := recover()
		if err != nil {
			_ = e.rw.setErrorLocked(errEnginePanic)
			panic(err)
		}
	}()

	e.opt.StatsOptions.measureWaitDurationSince(waitDo, startTimeBeforeLock)
	defer e.opt.StatsOptions.measureSqliteTxDurationSince(txDo, queryName, time.Now())
	err := e.internalDoLocked(do)
	return e.rw.setErrorLocked(err)
}

// В случае возникновения ошибки движок считается сломанным
func (e *Engine) internalDo(queryName string, do func(c internalConn) error) error {
	return e.internalDoBinlog(queryName, func(c internalConn) (int64, error) {
		return 0, do(c)
	})
}

func (e *Engine) internalDoLocked(do func(c internalConn) (int64, error)) (err error) {
	err = e.rw.beginTxLocked()
	if err != nil {
		e.opt.StatsOptions.engineBrokenEvent()
		return fmt.Errorf("failed to begin tx: %w", err)
	}
	defer func() {
		errRollback := e.rw.rollbackLocked()
		if errRollback != nil {
			err = multierr.Append(err, errRollback)
		}
	}()
	conn := newInternalConn(e.rw.conn)
	offset, err := do(conn)
	if err != nil {
		return fmt.Errorf("user logic error: %w", err)
	}
	if offset > 0 {
		return e.rw.binlogCommitTxLocked(offset)
	}
	return e.rw.nonBinlogCommitTxLocked()
}

func (e *Engine) binlogLoadOrCreateMeta() ([]byte, error) {
	var meta []byte
	err := e.internalDo("__load_binlog", func(conn internalConn) error {
		rows := conn.Query("__select_meta", "SELECT meta from __snapshot_meta")
		if rows.err != nil {
			return rows.err
		}
		for rows.Next() {
			meta, _ = rows.ColumnBlob(0, meta)
		}
		if meta != nil {
			return nil
		}
		err := conn.Exec("__insert_meta", "INSERT INTO __snapshot_meta(meta) VALUES($meta)", Blob("$meta", meta))
		return err
	})
	return meta, err
}

func (e *Engine) binlogLoadOrCreatePosition() (int64, error) {
	var offset int64
	err := e.internalDo("__load_binlog", func(conn internalConn) error {
		var isExists bool
		var err error
		offset, isExists, err = binlogLoadPosition(conn)
		if err != nil {
			return err
		}
		if isExists {
			return nil
		}
		err = conn.Exec("__insert_binlog_pos", "INSERT INTO __binlog_offset(offset) VALUES(0)")
		return err
	})
	return offset, err
}

func (e *Engine) Close() error {
	return e.close(e.binlog != nil)
}

func (e *Engine) close(waitCommitBinlog bool) error {
	e.logger.Printf("starting close, waitCommitBinlog: %t", waitCommitBinlog)
	defer e.opt.StatsOptions.measureActionDurationSince(closeEngine, time.Now())
	readOnly := e.readOnly
	if !readOnly {
		e.rw.mu.Lock()
		e.logger.Println("set readOnly")
		e.readOnly = true
		e.rw.mu.Unlock()
	}
	var error error
	if waitCommitBinlog {
		e.logger.Println("calling binlog.Shutdown")
		err := e.binlog.Shutdown()
		if err != nil {
			multierr.AppendInto(&error, err)
		}
		<-e.finishBinlogRunCh
		e.binlogEngine.checkpointer.doCheckpointIfCan()
	}
	if !readOnly {
		e.logger.Println("closing RW connection")
		err := e.rw.Close()
		if err != nil {
			multierr.AppendInto(&error, fmt.Errorf("failed to close RW connection: %w", err))
		}
	}
	e.logger.Println("closing RO connection pool")
	e.roConnPool.Close(&error)
	if !readOnly {
		error = multierr.Append(error, e.re.Close())
	}

	return error
}

func binlogLoadPosition(conn internalConn) (offset int64, isExists bool, err error) {
	rows := conn.Query("__select_binlog_pos", "SELECT offset from __binlog_offset")
	if rows.err != nil {
		return 0, false, rows.err
	}
	for rows.Next() {
		offset := rows.ColumnInteger(0)
		return offset, true, nil
	}
	return 0, false, nil
}

func checkUserQueryName(qn string) error {
	if len(qn) > 2 && qn[0:2] == internalQueryPrefix {
		return fmt.Errorf("query prefix %q is reserved, got: %s", internalQueryPrefix, qn)
	}
	return nil
}

func checkInternalQueryName(qn string) error {
	if len(qn) > 2 && qn[0:2] == internalQueryPrefix {
		return nil
	}
	return fmt.Errorf("use prefix %s for internal query, got: %s", internalQueryPrefix, qn)
}

func (e *Engine) rareLog(format string, v ...any) {
	e.logMx.Lock()
	defer e.logMx.Unlock()
	now := time.Now()
	if now.After(e.nextLogTime) {
		e.logger.Printf(format, v...)
		e.nextLogTime = now.Add(time.Second * 10)
	}
}