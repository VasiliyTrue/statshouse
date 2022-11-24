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

type StatshouseApiSeriesMeta struct {
	FieldsMask uint32
	TimeShift  int64
	Tags       map[string]string
	What       StatshouseApiFunction // Conditional: item.FieldsMask.1
}

func (StatshouseApiSeriesMeta) TLName() string { return "statshouseApi.seriesMeta" }
func (StatshouseApiSeriesMeta) TLTag() uint32  { return 0x5c2bf286 }

func (item *StatshouseApiSeriesMeta) SetWhat(v StatshouseApiFunction) {
	item.What = v
	item.FieldsMask |= 1 << 1
}
func (item *StatshouseApiSeriesMeta) ClearWhat() {
	item.What.Reset()
	item.FieldsMask &^= 1 << 1
}
func (item *StatshouseApiSeriesMeta) IsSetWhat() bool { return item.FieldsMask&(1<<1) != 0 }

func (item *StatshouseApiSeriesMeta) Reset() {
	item.FieldsMask = 0
	item.TimeShift = 0
	VectorDictionaryFieldString0Reset(item.Tags)
	item.What.Reset()
}

func (item *StatshouseApiSeriesMeta) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldsMask); err != nil {
		return w, err
	}
	if w, err = basictl.LongRead(w, &item.TimeShift); err != nil {
		return w, err
	}
	if w, err = VectorDictionaryFieldString0Read(w, &item.Tags); err != nil {
		return w, err
	}
	if item.FieldsMask&(1<<1) != 0 {
		if w, err = item.What.ReadBoxed(w); err != nil {
			return w, err
		}
	} else {
		item.What.Reset()
	}
	return w, nil
}

func (item *StatshouseApiSeriesMeta) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.FieldsMask)
	w = basictl.LongWrite(w, item.TimeShift)
	if w, err = VectorDictionaryFieldString0Write(w, item.Tags); err != nil {
		return w, err
	}
	if item.FieldsMask&(1<<1) != 0 {
		if w, err = item.What.WriteBoxed(w); err != nil {
			return w, err
		}
	}
	return w, nil
}

func (item *StatshouseApiSeriesMeta) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x5c2bf286); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *StatshouseApiSeriesMeta) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x5c2bf286)
	return item.Write(w)
}

func (item StatshouseApiSeriesMeta) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func StatshouseApiSeriesMeta__ReadJSON(item *StatshouseApiSeriesMeta, j interface{}) error {
	return item.readJSON(j)
}
func (item *StatshouseApiSeriesMeta) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("statshouseApi.seriesMeta", "expected json object")
	}
	_jFieldsMask := _jm["fields_mask"]
	delete(_jm, "fields_mask")
	if err := JsonReadUint32(_jFieldsMask, &item.FieldsMask); err != nil {
		return err
	}
	_jTimeShift := _jm["time_shift"]
	delete(_jm, "time_shift")
	if err := JsonReadInt64(_jTimeShift, &item.TimeShift); err != nil {
		return err
	}
	_jTags := _jm["tags"]
	delete(_jm, "tags")
	_jWhat := _jm["what"]
	delete(_jm, "what")
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("statshouseApi.seriesMeta", k)
	}
	if _jWhat != nil {
		item.FieldsMask |= 1 << 1
	}
	if err := VectorDictionaryFieldString0ReadJSON(_jTags, &item.Tags); err != nil {
		return err
	}
	if _jWhat != nil {
		if err := StatshouseApiFunction__ReadJSON(&item.What, _jWhat); err != nil {
			return err
		}
	} else {
		item.What.Reset()
	}
	return nil
}

func (item *StatshouseApiSeriesMeta) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if item.FieldsMask != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"fields_mask":`...)
		w = basictl.JSONWriteUint32(w, item.FieldsMask)
	}
	if item.TimeShift != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"time_shift":`...)
		w = basictl.JSONWriteInt64(w, item.TimeShift)
	}
	if len(item.Tags) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"tags":`...)
		if w, err = VectorDictionaryFieldString0WriteJSON(w, item.Tags); err != nil {
			return w, err
		}
	}
	if item.FieldsMask&(1<<1) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"what":`...)
		if w, err = item.What.WriteJSON(w); err != nil {
			return w, err
		}
	}
	return append(w, '}'), nil
}

func (item *StatshouseApiSeriesMeta) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *StatshouseApiSeriesMeta) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("statshouseApi.seriesMeta", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("statshouseApi.seriesMeta", err.Error())
	}
	return nil
}

func VectorStatshouseApiSeriesMeta0Read(w []byte, vec *[]StatshouseApiSeriesMeta) (_ []byte, err error) {
	var l uint32
	if w, err = basictl.NatRead(w, &l); err != nil {
		return w, err
	}
	if err = basictl.CheckLengthSanity(w, l, 4); err != nil {
		return w, err
	}
	if uint32(cap(*vec)) < l {
		*vec = make([]StatshouseApiSeriesMeta, l)
	} else {
		*vec = (*vec)[:l]
	}
	for i := range *vec {
		if w, err = (*vec)[i].Read(w); err != nil {
			return w, err
		}
	}
	return w, nil
}

func VectorStatshouseApiSeriesMeta0Write(w []byte, vec []StatshouseApiSeriesMeta) (_ []byte, err error) {
	w = basictl.NatWrite(w, uint32(len(vec)))
	for _, elem := range vec {
		if w, err = elem.Write(w); err != nil {
			return w, err
		}
	}
	return w, nil
}

func VectorStatshouseApiSeriesMeta0ReadJSON(j interface{}, vec *[]StatshouseApiSeriesMeta) error {
	l, _arr, err := JsonReadArray("[]StatshouseApiSeriesMeta", j)
	if err != nil {
		return err
	}
	if cap(*vec) < l {
		*vec = make([]StatshouseApiSeriesMeta, l)
	} else {
		*vec = (*vec)[:l]
	}
	for i := range *vec {
		if err := StatshouseApiSeriesMeta__ReadJSON(&(*vec)[i], _arr[i]); err != nil {
			return err
		}
	}
	return nil
}

func VectorStatshouseApiSeriesMeta0WriteJSON(w []byte, vec []StatshouseApiSeriesMeta) (_ []byte, err error) {
	w = append(w, '[')
	for _, elem := range vec {
		w = basictl.JSONAddCommaIfNeeded(w)
		if w, err = elem.WriteJSON(w); err != nil {
			return w, err
		}
	}
	return append(w, ']'), nil
}