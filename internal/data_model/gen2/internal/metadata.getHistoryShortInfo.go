// Copyright 2023 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package internal

import (
	"github.com/vkcom/statshouse/internal/vkgo/basictl"
)

var _ = basictl.NatWrite

type MetadataGetHistoryShortInfo struct {
	FieldsMask uint32
	Id         int64
}

func (MetadataGetHistoryShortInfo) TLName() string { return "metadata.getHistoryShortInfo" }
func (MetadataGetHistoryShortInfo) TLTag() uint32  { return 0x22ff6a79 }

func (item *MetadataGetHistoryShortInfo) Reset() {
	item.FieldsMask = 0
	item.Id = 0
}

func (item *MetadataGetHistoryShortInfo) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldsMask); err != nil {
		return w, err
	}
	return basictl.LongRead(w, &item.Id)
}

func (item *MetadataGetHistoryShortInfo) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.FieldsMask)
	return basictl.LongWrite(w, item.Id), nil
}

func (item *MetadataGetHistoryShortInfo) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x22ff6a79); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *MetadataGetHistoryShortInfo) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x22ff6a79)
	return item.Write(w)
}

func (item *MetadataGetHistoryShortInfo) ReadResult(w []byte, ret *MetadataHistoryShortResponse) (_ []byte, err error) {
	return ret.ReadBoxed(w, item.FieldsMask)
}

func (item *MetadataGetHistoryShortInfo) WriteResult(w []byte, ret MetadataHistoryShortResponse) (_ []byte, err error) {
	return ret.WriteBoxed(w, item.FieldsMask)
}

func (item *MetadataGetHistoryShortInfo) ReadResultJSON(j interface{}, ret *MetadataHistoryShortResponse) error {
	if err := MetadataHistoryShortResponse__ReadJSON(ret, j, item.FieldsMask); err != nil {
		return err
	}
	return nil
}

func (item *MetadataGetHistoryShortInfo) WriteResultJSON(w []byte, ret MetadataHistoryShortResponse) (_ []byte, err error) {
	return item.writeResultJSON(false, w, ret)
}

func (item *MetadataGetHistoryShortInfo) writeResultJSON(short bool, w []byte, ret MetadataHistoryShortResponse) (_ []byte, err error) {
	if w, err = ret.WriteJSONOpt(short, w, item.FieldsMask); err != nil {
		return w, err
	}
	return w, nil
}

func (item *MetadataGetHistoryShortInfo) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret MetadataHistoryShortResponse
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *MetadataGetHistoryShortInfo) ReadResultWriteResultJSONShort(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret MetadataHistoryShortResponse
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.writeResultJSON(true, w, ret)
	return r, w, err
}

func (item *MetadataGetHistoryShortInfo) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	j, err := JsonBytesToInterface(r)
	if err != nil {
		return r, w, ErrorInvalidJSON("metadata.getHistoryShortInfo", err.Error())
	}
	var ret MetadataHistoryShortResponse
	if err = item.ReadResultJSON(j, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item MetadataGetHistoryShortInfo) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func MetadataGetHistoryShortInfo__ReadJSON(item *MetadataGetHistoryShortInfo, j interface{}) error {
	return item.readJSON(j)
}
func (item *MetadataGetHistoryShortInfo) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("metadata.getHistoryShortInfo", "expected json object")
	}
	_jFieldsMask := _jm["fields_mask"]
	delete(_jm, "fields_mask")
	if err := JsonReadUint32(_jFieldsMask, &item.FieldsMask); err != nil {
		return err
	}
	_jId := _jm["id"]
	delete(_jm, "id")
	if err := JsonReadInt64(_jId, &item.Id); err != nil {
		return err
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("metadata.getHistoryShortInfo", k)
	}
	return nil
}

func (item *MetadataGetHistoryShortInfo) WriteJSON(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(false, w)
}
func (item *MetadataGetHistoryShortInfo) WriteJSONOpt(short bool, w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if item.FieldsMask != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"fields_mask":`...)
		w = basictl.JSONWriteUint32(w, item.FieldsMask)
	}
	if item.Id != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"id":`...)
		w = basictl.JSONWriteInt64(w, item.Id)
	}
	return append(w, '}'), nil
}

func (item *MetadataGetHistoryShortInfo) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *MetadataGetHistoryShortInfo) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("metadata.getHistoryShortInfo", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("metadata.getHistoryShortInfo", err.Error())
	}
	return nil
}