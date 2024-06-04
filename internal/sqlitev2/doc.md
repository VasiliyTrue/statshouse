# sqlitev2

Обертка над sqlite версии 2. 

Цели:

- Интеграция с бинлогами (на данный момент с файловыми, но в будущем и с барсиковыми)
- Предоставление удобного go интерфейса над sqlite
- Простота внутренностей относительно версии 1

Заметки:

- Используется одно RW соединение и много RO
- Работа с каждым соедининием должна выполняться строго под блокировкой
- Коммит в базу делается в конце каждого Do
- `View` - видит только то, что было закомиченно в базу
- Бинлоги используются для удобства репликации

Открытые вопросы:

- Дизайн внешнего интерфейса
- Требуемый дополнительный функционал 
- Как правильно тестировать


Для ознакомления

- `OpenEngine`- конструктор
- `engine.Run` - запуск потока бинлога
- `engine.Do` - обработка запросов на запись
- `engine.View` - обработка запросов на чтение
- `engine.Backup` - создание бэкапа
- `engine.Close` - закрывает бинлог и все соединения
- публичные методы `binlogEngine`