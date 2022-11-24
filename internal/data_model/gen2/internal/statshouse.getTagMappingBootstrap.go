// Copyright 2022 V Kontakte LLC
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

type StatshouseGetTagMappingBootstrap struct {
	FieldsMask uint32
	Header     StatshouseCommonProxyHeader
}

func (StatshouseGetTagMappingBootstrap) TLName() string { return "statshouse.getTagMappingBootstrap" }
func (StatshouseGetTagMappingBootstrap) TLTag() uint32  { return 0x75a7f68e }

func (item *StatshouseGetTagMappingBootstrap) Reset() {
	item.FieldsMask = 0
	item.Header.Reset()
}

func (item *StatshouseGetTagMappingBootstrap) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldsMask); err != nil {
		return w, err
	}
	return item.Header.Read(w, item.FieldsMask)
}

func (item *StatshouseGetTagMappingBootstrap) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.FieldsMask)
	return item.Header.Write(w, item.FieldsMask)
}

func (item *StatshouseGetTagMappingBootstrap) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x75a7f68e); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *StatshouseGetTagMappingBootstrap) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x75a7f68e)
	return item.Write(w)
}

func (item *StatshouseGetTagMappingBootstrap) ReadResult(w []byte, ret *StatshouseGetTagMappingBootstrapResult) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *StatshouseGetTagMappingBootstrap) WriteResult(w []byte, ret StatshouseGetTagMappingBootstrapResult) (_ []byte, err error) {
	return ret.WriteBoxed(w)
}

func (item *StatshouseGetTagMappingBootstrap) ReadResultJSON(j interface{}, ret *StatshouseGetTagMappingBootstrapResult) error {
	if err := StatshouseGetTagMappingBootstrapResult__ReadJSON(ret, j); err != nil {
		return err
	}
	return nil
}

func (item *StatshouseGetTagMappingBootstrap) WriteResultJSON(w []byte, ret StatshouseGetTagMappingBootstrapResult) (_ []byte, err error) {
	if w, err = ret.WriteJSON(w); err != nil {
		return w, err
	}
	return w, nil
}

func (item *StatshouseGetTagMappingBootstrap) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret StatshouseGetTagMappingBootstrapResult
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *StatshouseGetTagMappingBootstrap) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	j, err := JsonBytesToInterface(r)
	if err != nil {
		return r, w, ErrorInvalidJSON("statshouse.getTagMappingBootstrap", err.Error())
	}
	var ret StatshouseGetTagMappingBootstrapResult
	if err = item.ReadResultJSON(j, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item StatshouseGetTagMappingBootstrap) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func StatshouseGetTagMappingBootstrap__ReadJSON(item *StatshouseGetTagMappingBootstrap, j interface{}) error {
	return item.readJSON(j)
}
func (item *StatshouseGetTagMappingBootstrap) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("statshouse.getTagMappingBootstrap", "expected json object")
	}
	_jFieldsMask := _jm["fields_mask"]
	delete(_jm, "fields_mask")
	if err := JsonReadUint32(_jFieldsMask, &item.FieldsMask); err != nil {
		return err
	}
	_jHeader := _jm["header"]
	delete(_jm, "header")
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("statshouse.getTagMappingBootstrap", k)
	}
	if err := StatshouseCommonProxyHeader__ReadJSON(&item.Header, _jHeader, item.FieldsMask); err != nil {
		return err
	}
	return nil
}

func (item *StatshouseGetTagMappingBootstrap) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if item.FieldsMask != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"fields_mask":`...)
		w = basictl.JSONWriteUint32(w, item.FieldsMask)
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"header":`...)
	if w, err = item.Header.WriteJSON(w, item.FieldsMask); err != nil {
		return w, err
	}
	return append(w, '}'), nil
}

func (item *StatshouseGetTagMappingBootstrap) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *StatshouseGetTagMappingBootstrap) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("statshouse.getTagMappingBootstrap", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("statshouse.getTagMappingBootstrap", err.Error())
	}
	return nil
}

type StatshouseGetTagMappingBootstrapBytes struct {
	FieldsMask uint32
	Header     StatshouseCommonProxyHeaderBytes
}

func (StatshouseGetTagMappingBootstrapBytes) TLName() string {
	return "statshouse.getTagMappingBootstrap"
}
func (StatshouseGetTagMappingBootstrapBytes) TLTag() uint32 { return 0x75a7f68e }

func (item *StatshouseGetTagMappingBootstrapBytes) Reset() {
	item.FieldsMask = 0
	item.Header.Reset()
}

func (item *StatshouseGetTagMappingBootstrapBytes) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldsMask); err != nil {
		return w, err
	}
	return item.Header.Read(w, item.FieldsMask)
}

func (item *StatshouseGetTagMappingBootstrapBytes) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.FieldsMask)
	return item.Header.Write(w, item.FieldsMask)
}

func (item *StatshouseGetTagMappingBootstrapBytes) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x75a7f68e); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *StatshouseGetTagMappingBootstrapBytes) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x75a7f68e)
	return item.Write(w)
}

func (item *StatshouseGetTagMappingBootstrapBytes) ReadResult(w []byte, ret *StatshouseGetTagMappingBootstrapResultBytes) (_ []byte, err error) {
	return ret.ReadBoxed(w)
}

func (item *StatshouseGetTagMappingBootstrapBytes) WriteResult(w []byte, ret StatshouseGetTagMappingBootstrapResultBytes) (_ []byte, err error) {
	return ret.WriteBoxed(w)
}

func (item *StatshouseGetTagMappingBootstrapBytes) ReadResultJSON(j interface{}, ret *StatshouseGetTagMappingBootstrapResultBytes) error {
	if err := StatshouseGetTagMappingBootstrapResultBytes__ReadJSON(ret, j); err != nil {
		return err
	}
	return nil
}

func (item *StatshouseGetTagMappingBootstrapBytes) WriteResultJSON(w []byte, ret StatshouseGetTagMappingBootstrapResultBytes) (_ []byte, err error) {
	if w, err = ret.WriteJSON(w); err != nil {
		return w, err
	}
	return w, nil
}

func (item *StatshouseGetTagMappingBootstrapBytes) ReadResultWriteResultJSON(r []byte, w []byte) (_ []byte, _ []byte, err error) {
	var ret StatshouseGetTagMappingBootstrapResultBytes
	if r, err = item.ReadResult(r, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResultJSON(w, ret)
	return r, w, err
}

func (item *StatshouseGetTagMappingBootstrapBytes) ReadResultJSONWriteResult(r []byte, w []byte) ([]byte, []byte, error) {
	j, err := JsonBytesToInterface(r)
	if err != nil {
		return r, w, ErrorInvalidJSON("statshouse.getTagMappingBootstrap", err.Error())
	}
	var ret StatshouseGetTagMappingBootstrapResultBytes
	if err = item.ReadResultJSON(j, &ret); err != nil {
		return r, w, err
	}
	w, err = item.WriteResult(w, ret)
	return r, w, err
}

func (item StatshouseGetTagMappingBootstrapBytes) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func StatshouseGetTagMappingBootstrapBytes__ReadJSON(item *StatshouseGetTagMappingBootstrapBytes, j interface{}) error {
	return item.readJSON(j)
}
func (item *StatshouseGetTagMappingBootstrapBytes) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("statshouse.getTagMappingBootstrap", "expected json object")
	}
	_jFieldsMask := _jm["fields_mask"]
	delete(_jm, "fields_mask")
	if err := JsonReadUint32(_jFieldsMask, &item.FieldsMask); err != nil {
		return err
	}
	_jHeader := _jm["header"]
	delete(_jm, "header")
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("statshouse.getTagMappingBootstrap", k)
	}
	if err := StatshouseCommonProxyHeaderBytes__ReadJSON(&item.Header, _jHeader, item.FieldsMask); err != nil {
		return err
	}
	return nil
}

func (item *StatshouseGetTagMappingBootstrapBytes) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if item.FieldsMask != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"fields_mask":`...)
		w = basictl.JSONWriteUint32(w, item.FieldsMask)
	}
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"header":`...)
	if w, err = item.Header.WriteJSON(w, item.FieldsMask); err != nil {
		return w, err
	}
	return append(w, '}'), nil
}

func (item *StatshouseGetTagMappingBootstrapBytes) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *StatshouseGetTagMappingBootstrapBytes) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("statshouse.getTagMappingBootstrap", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("statshouse.getTagMappingBootstrap", err.Error())
	}
	return nil
}