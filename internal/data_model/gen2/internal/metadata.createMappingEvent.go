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

var _ = True{}

type MetadataCreateMappingEvent struct {
	FieldMask uint32
	Id        int32
	Key       string
	Metric    string
	Badget    int64
	// Create True // Conditional: item.FieldMask.0
	UpdatedAt uint32
}

func (MetadataCreateMappingEvent) TLName() string { return "metadata.createMappingEvent" }
func (MetadataCreateMappingEvent) TLTag() uint32  { return 0x12345678 }

func (item *MetadataCreateMappingEvent) SetCreate(v bool) {
	if v {
		item.FieldMask |= 1 << 0
	} else {
		item.FieldMask &^= 1 << 0
	}
}
func (item *MetadataCreateMappingEvent) IsSetCreate() bool { return item.FieldMask&(1<<0) != 0 }

func (item *MetadataCreateMappingEvent) Reset() {
	item.FieldMask = 0
	item.Id = 0
	item.Key = ""
	item.Metric = ""
	item.Badget = 0
	item.UpdatedAt = 0
}

func (item *MetadataCreateMappingEvent) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldMask); err != nil {
		return w, err
	}
	if w, err = basictl.IntRead(w, &item.Id); err != nil {
		return w, err
	}
	if w, err = basictl.StringRead(w, &item.Key); err != nil {
		return w, err
	}
	if w, err = basictl.StringRead(w, &item.Metric); err != nil {
		return w, err
	}
	if w, err = basictl.LongRead(w, &item.Badget); err != nil {
		return w, err
	}
	return basictl.NatRead(w, &item.UpdatedAt)
}

func (item *MetadataCreateMappingEvent) Write(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, item.FieldMask)
	w = basictl.IntWrite(w, item.Id)
	if w, err = basictl.StringWrite(w, item.Key); err != nil {
		return w, err
	}
	if w, err = basictl.StringWrite(w, item.Metric); err != nil {
		return w, err
	}
	w = basictl.LongWrite(w, item.Badget)
	return basictl.NatWrite(w, item.UpdatedAt), nil
}

func (item *MetadataCreateMappingEvent) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x12345678); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *MetadataCreateMappingEvent) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x12345678)
	return item.Write(w)
}

func (item MetadataCreateMappingEvent) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func MetadataCreateMappingEvent__ReadJSON(item *MetadataCreateMappingEvent, j interface{}) error {
	return item.readJSON(j)
}
func (item *MetadataCreateMappingEvent) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("metadata.createMappingEvent", "expected json object")
	}
	_jFieldMask := _jm["field_mask"]
	delete(_jm, "field_mask")
	if err := JsonReadUint32(_jFieldMask, &item.FieldMask); err != nil {
		return err
	}
	_jId := _jm["id"]
	delete(_jm, "id")
	if err := JsonReadInt32(_jId, &item.Id); err != nil {
		return err
	}
	_jKey := _jm["key"]
	delete(_jm, "key")
	if err := JsonReadString(_jKey, &item.Key); err != nil {
		return err
	}
	_jMetric := _jm["metric"]
	delete(_jm, "metric")
	if err := JsonReadString(_jMetric, &item.Metric); err != nil {
		return err
	}
	_jBadget := _jm["badget"]
	delete(_jm, "badget")
	if err := JsonReadInt64(_jBadget, &item.Badget); err != nil {
		return err
	}
	_jCreate := _jm["create"]
	delete(_jm, "create")
	_jUpdatedAt := _jm["updated_at"]
	delete(_jm, "updated_at")
	if err := JsonReadUint32(_jUpdatedAt, &item.UpdatedAt); err != nil {
		return err
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("metadata.createMappingEvent", k)
	}
	if _jCreate != nil {
		_bit := false
		if err := JsonReadBool(_jCreate, &_bit); err != nil {
			return err
		}
		if _bit {
			item.FieldMask |= 1 << 0
		} else {
			item.FieldMask &^= 1 << 0
		}
	}
	return nil
}

func (item *MetadataCreateMappingEvent) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if item.FieldMask != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"field_mask":`...)
		w = basictl.JSONWriteUint32(w, item.FieldMask)
	}
	if item.Id != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"id":`...)
		w = basictl.JSONWriteInt32(w, item.Id)
	}
	if len(item.Key) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"key":`...)
		w = basictl.JSONWriteString(w, item.Key)
	}
	if len(item.Metric) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"metric":`...)
		w = basictl.JSONWriteString(w, item.Metric)
	}
	if item.Badget != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"badget":`...)
		w = basictl.JSONWriteInt64(w, item.Badget)
	}
	if item.FieldMask&(1<<0) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"create":true`...)
	}
	if item.UpdatedAt != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"updated_at":`...)
		w = basictl.JSONWriteUint32(w, item.UpdatedAt)
	}
	return append(w, '}'), nil
}

func (item *MetadataCreateMappingEvent) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *MetadataCreateMappingEvent) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("metadata.createMappingEvent", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("metadata.createMappingEvent", err.Error())
	}
	return nil
}