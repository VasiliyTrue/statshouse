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

func (item EngineAlreadyInMasterMode) AsUnion() EngineSwitchMasterReplicaModeResultUnion {
	var ret EngineSwitchMasterReplicaModeResultUnion
	ret.SetAlreadyInMasterMode()
	return ret
}

// AsUnion will be here
type EngineAlreadyInMasterMode struct {
}

func (EngineAlreadyInMasterMode) TLName() string { return "engine.alreadyInMasterMode" }
func (EngineAlreadyInMasterMode) TLTag() uint32  { return 0x402409cb }

func (item *EngineAlreadyInMasterMode) Reset()                         {}
func (item *EngineAlreadyInMasterMode) Read(w []byte) ([]byte, error)  { return w, nil }
func (item *EngineAlreadyInMasterMode) Write(w []byte) ([]byte, error) { return w, nil }
func (item *EngineAlreadyInMasterMode) ReadBoxed(w []byte) ([]byte, error) {
	return basictl.NatReadExactTag(w, 0x402409cb)
}
func (item *EngineAlreadyInMasterMode) WriteBoxed(w []byte) ([]byte, error) {
	return basictl.NatWrite(w, 0x402409cb), nil
}

func (item EngineAlreadyInMasterMode) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func EngineAlreadyInMasterMode__ReadJSON(item *EngineAlreadyInMasterMode, j interface{}) error {
	return item.readJSON(j)
}
func (item *EngineAlreadyInMasterMode) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("engine.alreadyInMasterMode", "expected json object")
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("engine.alreadyInMasterMode", k)
	}
	return nil
}

func (item *EngineAlreadyInMasterMode) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '{')
	return append(w, '}'), nil
}

func (item *EngineAlreadyInMasterMode) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *EngineAlreadyInMasterMode) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("engine.alreadyInMasterMode", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("engine.alreadyInMasterMode", err.Error())
	}
	return nil
}

func (item EngineAlreadyInReplicaMode) AsUnion() EngineSwitchMasterReplicaModeResultUnion {
	var ret EngineSwitchMasterReplicaModeResultUnion
	ret.SetAlreadyInReplicaMode()
	return ret
}

// AsUnion will be here
type EngineAlreadyInReplicaMode struct {
}

func (EngineAlreadyInReplicaMode) TLName() string { return "engine.alreadyInReplicaMode" }
func (EngineAlreadyInReplicaMode) TLTag() uint32  { return 0xebd80142 }

func (item *EngineAlreadyInReplicaMode) Reset()                         {}
func (item *EngineAlreadyInReplicaMode) Read(w []byte) ([]byte, error)  { return w, nil }
func (item *EngineAlreadyInReplicaMode) Write(w []byte) ([]byte, error) { return w, nil }
func (item *EngineAlreadyInReplicaMode) ReadBoxed(w []byte) ([]byte, error) {
	return basictl.NatReadExactTag(w, 0xebd80142)
}
func (item *EngineAlreadyInReplicaMode) WriteBoxed(w []byte) ([]byte, error) {
	return basictl.NatWrite(w, 0xebd80142), nil
}

func (item EngineAlreadyInReplicaMode) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func EngineAlreadyInReplicaMode__ReadJSON(item *EngineAlreadyInReplicaMode, j interface{}) error {
	return item.readJSON(j)
}
func (item *EngineAlreadyInReplicaMode) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("engine.alreadyInReplicaMode", "expected json object")
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("engine.alreadyInReplicaMode", k)
	}
	return nil
}

func (item *EngineAlreadyInReplicaMode) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '{')
	return append(w, '}'), nil
}

func (item *EngineAlreadyInReplicaMode) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *EngineAlreadyInReplicaMode) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("engine.alreadyInReplicaMode", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("engine.alreadyInReplicaMode", err.Error())
	}
	return nil
}

func (item EngineFailedToSwitchMode) AsUnion() EngineSwitchMasterReplicaModeResultUnion {
	var ret EngineSwitchMasterReplicaModeResultUnion
	ret.SetFailedToSwitchMode(item)
	return ret
}

// AsUnion will be here
type EngineFailedToSwitchMode struct {
	Error string
}

func (EngineFailedToSwitchMode) TLName() string { return "engine.failedToSwitchMode" }
func (EngineFailedToSwitchMode) TLTag() uint32  { return 0x17418662 }

func (item *EngineFailedToSwitchMode) Reset() {
	item.Error = ""
}

func (item *EngineFailedToSwitchMode) Read(w []byte) (_ []byte, err error) {
	return basictl.StringRead(w, &item.Error)
}

func (item *EngineFailedToSwitchMode) Write(w []byte) (_ []byte, err error) {
	return basictl.StringWrite(w, item.Error)
}

func (item *EngineFailedToSwitchMode) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x17418662); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *EngineFailedToSwitchMode) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0x17418662)
	return item.Write(w)
}

func (item EngineFailedToSwitchMode) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func EngineFailedToSwitchMode__ReadJSON(item *EngineFailedToSwitchMode, j interface{}) error {
	return item.readJSON(j)
}
func (item *EngineFailedToSwitchMode) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("engine.failedToSwitchMode", "expected json object")
	}
	_jError := _jm["error"]
	delete(_jm, "error")
	if err := JsonReadString(_jError, &item.Error); err != nil {
		return err
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("engine.failedToSwitchMode", k)
	}
	return nil
}

func (item *EngineFailedToSwitchMode) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if len(item.Error) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"error":`...)
		w = basictl.JSONWriteString(w, item.Error)
	}
	return append(w, '}'), nil
}

func (item *EngineFailedToSwitchMode) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *EngineFailedToSwitchMode) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("engine.failedToSwitchMode", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("engine.failedToSwitchMode", err.Error())
	}
	return nil
}

var _EngineSwitchMasterReplicaModeResultUnion = [6]UnionElement{
	{TLTag: 0x95b13964, TLName: "engine.switchedToMasterMode", TLString: "engine.switchedToMasterMode#95b13964"},
	{TLTag: 0xad642a0b, TLName: "engine.switchedToReplicaMode", TLString: "engine.switchedToReplicaMode#ad642a0b"},
	{TLTag: 0x402409cb, TLName: "engine.alreadyInMasterMode", TLString: "engine.alreadyInMasterMode#402409cb"},
	{TLTag: 0xebd80142, TLName: "engine.alreadyInReplicaMode", TLString: "engine.alreadyInReplicaMode#ebd80142"},
	{TLTag: 0xec61b4be, TLName: "engine.switchedToMasterModeForcefully", TLString: "engine.switchedToMasterModeForcefully#ec61b4be"},
	{TLTag: 0x17418662, TLName: "engine.failedToSwitchMode", TLString: "engine.failedToSwitchMode#17418662"},
}

type EngineSwitchMasterReplicaModeResultUnion struct {
	valueSwitchedToMasterModeForcefully EngineSwitchedToMasterModeForcefully
	valueFailedToSwitchMode             EngineFailedToSwitchMode
	index                               int
}

func (item EngineSwitchMasterReplicaModeResultUnion) TLName() string {
	return _EngineSwitchMasterReplicaModeResultUnion[item.index].TLName
}
func (item EngineSwitchMasterReplicaModeResultUnion) TLTag() uint32 {
	return _EngineSwitchMasterReplicaModeResultUnion[item.index].TLTag
}

func (item *EngineSwitchMasterReplicaModeResultUnion) Reset() { item.index = 0 }

func (item *EngineSwitchMasterReplicaModeResultUnion) IsSwitchedToMasterMode() bool {
	return item.index == 0
}

func (item *EngineSwitchMasterReplicaModeResultUnion) AsSwitchedToMasterMode() (EngineSwitchedToMasterMode, bool) {
	var value EngineSwitchedToMasterMode
	return value, item.index == 0
}
func (item *EngineSwitchMasterReplicaModeResultUnion) ResetToSwitchedToMasterMode() { item.index = 0 }
func (item *EngineSwitchMasterReplicaModeResultUnion) SetSwitchedToMasterMode()     { item.index = 0 }

func (item *EngineSwitchMasterReplicaModeResultUnion) IsSwitchedToReplicaMode() bool {
	return item.index == 1
}

func (item *EngineSwitchMasterReplicaModeResultUnion) AsSwitchedToReplicaMode() (EngineSwitchedToReplicaMode, bool) {
	var value EngineSwitchedToReplicaMode
	return value, item.index == 1
}
func (item *EngineSwitchMasterReplicaModeResultUnion) ResetToSwitchedToReplicaMode() { item.index = 1 }
func (item *EngineSwitchMasterReplicaModeResultUnion) SetSwitchedToReplicaMode()     { item.index = 1 }

func (item *EngineSwitchMasterReplicaModeResultUnion) IsAlreadyInMasterMode() bool {
	return item.index == 2
}

func (item *EngineSwitchMasterReplicaModeResultUnion) AsAlreadyInMasterMode() (EngineAlreadyInMasterMode, bool) {
	var value EngineAlreadyInMasterMode
	return value, item.index == 2
}
func (item *EngineSwitchMasterReplicaModeResultUnion) ResetToAlreadyInMasterMode() { item.index = 2 }
func (item *EngineSwitchMasterReplicaModeResultUnion) SetAlreadyInMasterMode()     { item.index = 2 }

func (item *EngineSwitchMasterReplicaModeResultUnion) IsAlreadyInReplicaMode() bool {
	return item.index == 3
}

func (item *EngineSwitchMasterReplicaModeResultUnion) AsAlreadyInReplicaMode() (EngineAlreadyInReplicaMode, bool) {
	var value EngineAlreadyInReplicaMode
	return value, item.index == 3
}
func (item *EngineSwitchMasterReplicaModeResultUnion) ResetToAlreadyInReplicaMode() { item.index = 3 }
func (item *EngineSwitchMasterReplicaModeResultUnion) SetAlreadyInReplicaMode()     { item.index = 3 }

func (item *EngineSwitchMasterReplicaModeResultUnion) IsSwitchedToMasterModeForcefully() bool {
	return item.index == 4
}

func (item *EngineSwitchMasterReplicaModeResultUnion) AsSwitchedToMasterModeForcefully() (*EngineSwitchedToMasterModeForcefully, bool) {
	if item.index != 4 {
		return nil, false
	}
	return &item.valueSwitchedToMasterModeForcefully, true
}
func (item *EngineSwitchMasterReplicaModeResultUnion) ResetToSwitchedToMasterModeForcefully() *EngineSwitchedToMasterModeForcefully {
	item.index = 4
	item.valueSwitchedToMasterModeForcefully.Reset()
	return &item.valueSwitchedToMasterModeForcefully
}
func (item *EngineSwitchMasterReplicaModeResultUnion) SetSwitchedToMasterModeForcefully(value EngineSwitchedToMasterModeForcefully) {
	item.index = 4
	item.valueSwitchedToMasterModeForcefully = value
}

func (item *EngineSwitchMasterReplicaModeResultUnion) IsFailedToSwitchMode() bool {
	return item.index == 5
}

func (item *EngineSwitchMasterReplicaModeResultUnion) AsFailedToSwitchMode() (*EngineFailedToSwitchMode, bool) {
	if item.index != 5 {
		return nil, false
	}
	return &item.valueFailedToSwitchMode, true
}
func (item *EngineSwitchMasterReplicaModeResultUnion) ResetToFailedToSwitchMode() *EngineFailedToSwitchMode {
	item.index = 5
	item.valueFailedToSwitchMode.Reset()
	return &item.valueFailedToSwitchMode
}
func (item *EngineSwitchMasterReplicaModeResultUnion) SetFailedToSwitchMode(value EngineFailedToSwitchMode) {
	item.index = 5
	item.valueFailedToSwitchMode = value
}

func (item *EngineSwitchMasterReplicaModeResultUnion) ReadBoxed(w []byte) (_ []byte, err error) {
	var tag uint32
	if w, err = basictl.NatRead(w, &tag); err != nil {
		return w, err
	}
	switch tag {
	case 0x95b13964:
		item.index = 0
		return w, nil
	case 0xad642a0b:
		item.index = 1
		return w, nil
	case 0x402409cb:
		item.index = 2
		return w, nil
	case 0xebd80142:
		item.index = 3
		return w, nil
	case 0xec61b4be:
		item.index = 4
		return item.valueSwitchedToMasterModeForcefully.Read(w)
	case 0x17418662:
		item.index = 5
		return item.valueFailedToSwitchMode.Read(w)
	default:
		return w, ErrorInvalidUnionTag("engine.SwitchMasterReplicaModeResult", tag)
	}
}

func (item *EngineSwitchMasterReplicaModeResultUnion) WriteBoxed(w []byte) (_ []byte, err error) {
	w = basictl.NatWrite(w, _EngineSwitchMasterReplicaModeResultUnion[item.index].TLTag)
	switch item.index {
	case 0:
		return w, nil
	case 1:
		return w, nil
	case 2:
		return w, nil
	case 3:
		return w, nil
	case 4:
		return item.valueSwitchedToMasterModeForcefully.Write(w)
	case 5:
		return item.valueFailedToSwitchMode.Write(w)
	default: // Impossible due to panic above
		return w, nil
	}
}

func EngineSwitchMasterReplicaModeResultUnion__ReadJSON(item *EngineSwitchMasterReplicaModeResultUnion, j interface{}) error {
	return item.readJSON(j)
}
func (item *EngineSwitchMasterReplicaModeResultUnion) readJSON(j interface{}) error {
	_jm, _tag, err := JsonReadUnionType("engine.SwitchMasterReplicaModeResult", j)
	if err != nil {
		return err
	}
	jvalue := _jm["value"]
	switch _tag {
	case "engine.switchedToMasterMode#95b13964", "engine.switchedToMasterMode", "#95b13964":
		item.index = 0
	case "engine.switchedToReplicaMode#ad642a0b", "engine.switchedToReplicaMode", "#ad642a0b":
		item.index = 1
	case "engine.alreadyInMasterMode#402409cb", "engine.alreadyInMasterMode", "#402409cb":
		item.index = 2
	case "engine.alreadyInReplicaMode#ebd80142", "engine.alreadyInReplicaMode", "#ebd80142":
		item.index = 3
	case "engine.switchedToMasterModeForcefully#ec61b4be", "engine.switchedToMasterModeForcefully", "#ec61b4be":
		item.index = 4
		if err := EngineSwitchedToMasterModeForcefully__ReadJSON(&item.valueSwitchedToMasterModeForcefully, jvalue); err != nil {
			return err
		}
		delete(_jm, "value")
	case "engine.failedToSwitchMode#17418662", "engine.failedToSwitchMode", "#17418662":
		item.index = 5
		if err := EngineFailedToSwitchMode__ReadJSON(&item.valueFailedToSwitchMode, jvalue); err != nil {
			return err
		}
		delete(_jm, "value")
	default:
		return ErrorInvalidUnionTagJSON("engine.SwitchMasterReplicaModeResult", _tag)
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("engine.SwitchMasterReplicaModeResult", k)
	}
	return nil
}

func (item *EngineSwitchMasterReplicaModeResultUnion) WriteJSON(w []byte) (_ []byte, err error) {
	switch item.index {
	case 0:
		return append(w, `{"type":"engine.switchedToMasterMode#95b13964"}`...), nil
	case 1:
		return append(w, `{"type":"engine.switchedToReplicaMode#ad642a0b"}`...), nil
	case 2:
		return append(w, `{"type":"engine.alreadyInMasterMode#402409cb"}`...), nil
	case 3:
		return append(w, `{"type":"engine.alreadyInReplicaMode#ebd80142"}`...), nil
	case 4:
		w = append(w, `{"type":"engine.switchedToMasterModeForcefully#ec61b4be","value":`...)
		if w, err = item.valueSwitchedToMasterModeForcefully.WriteJSON(w); err != nil {
			return w, err
		}
		return append(w, '}'), nil
	case 5:
		w = append(w, `{"type":"engine.failedToSwitchMode#17418662","value":`...)
		if w, err = item.valueFailedToSwitchMode.WriteJSON(w); err != nil {
			return w, err
		}
		return append(w, '}'), nil
	default: // Impossible due to panic above
		return w, nil
	}
}

func (item EngineSwitchMasterReplicaModeResultUnion) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func (item EngineSwitchedToMasterMode) AsUnion() EngineSwitchMasterReplicaModeResultUnion {
	var ret EngineSwitchMasterReplicaModeResultUnion
	ret.SetSwitchedToMasterMode()
	return ret
}

// AsUnion will be here
type EngineSwitchedToMasterMode struct {
}

func (EngineSwitchedToMasterMode) TLName() string { return "engine.switchedToMasterMode" }
func (EngineSwitchedToMasterMode) TLTag() uint32  { return 0x95b13964 }

func (item *EngineSwitchedToMasterMode) Reset()                         {}
func (item *EngineSwitchedToMasterMode) Read(w []byte) ([]byte, error)  { return w, nil }
func (item *EngineSwitchedToMasterMode) Write(w []byte) ([]byte, error) { return w, nil }
func (item *EngineSwitchedToMasterMode) ReadBoxed(w []byte) ([]byte, error) {
	return basictl.NatReadExactTag(w, 0x95b13964)
}
func (item *EngineSwitchedToMasterMode) WriteBoxed(w []byte) ([]byte, error) {
	return basictl.NatWrite(w, 0x95b13964), nil
}

func (item EngineSwitchedToMasterMode) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func EngineSwitchedToMasterMode__ReadJSON(item *EngineSwitchedToMasterMode, j interface{}) error {
	return item.readJSON(j)
}
func (item *EngineSwitchedToMasterMode) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("engine.switchedToMasterMode", "expected json object")
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("engine.switchedToMasterMode", k)
	}
	return nil
}

func (item *EngineSwitchedToMasterMode) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '{')
	return append(w, '}'), nil
}

func (item *EngineSwitchedToMasterMode) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *EngineSwitchedToMasterMode) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("engine.switchedToMasterMode", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("engine.switchedToMasterMode", err.Error())
	}
	return nil
}

func (item EngineSwitchedToMasterModeForcefully) AsUnion() EngineSwitchMasterReplicaModeResultUnion {
	var ret EngineSwitchMasterReplicaModeResultUnion
	ret.SetSwitchedToMasterModeForcefully(item)
	return ret
}

// AsUnion will be here
type EngineSwitchedToMasterModeForcefully struct {
	BytesTruncated int64
}

func (EngineSwitchedToMasterModeForcefully) TLName() string {
	return "engine.switchedToMasterModeForcefully"
}
func (EngineSwitchedToMasterModeForcefully) TLTag() uint32 { return 0xec61b4be }

func (item *EngineSwitchedToMasterModeForcefully) Reset() {
	item.BytesTruncated = 0
}

func (item *EngineSwitchedToMasterModeForcefully) Read(w []byte) (_ []byte, err error) {
	return basictl.LongRead(w, &item.BytesTruncated)
}

func (item *EngineSwitchedToMasterModeForcefully) Write(w []byte) (_ []byte, err error) {
	return basictl.LongWrite(w, item.BytesTruncated), nil
}

func (item *EngineSwitchedToMasterModeForcefully) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xec61b4be); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *EngineSwitchedToMasterModeForcefully) WriteBoxed(w []byte) ([]byte, error) {
	w = basictl.NatWrite(w, 0xec61b4be)
	return item.Write(w)
}

func (item EngineSwitchedToMasterModeForcefully) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func EngineSwitchedToMasterModeForcefully__ReadJSON(item *EngineSwitchedToMasterModeForcefully, j interface{}) error {
	return item.readJSON(j)
}
func (item *EngineSwitchedToMasterModeForcefully) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("engine.switchedToMasterModeForcefully", "expected json object")
	}
	_jBytesTruncated := _jm["bytes_truncated"]
	delete(_jm, "bytes_truncated")
	if err := JsonReadInt64(_jBytesTruncated, &item.BytesTruncated); err != nil {
		return err
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("engine.switchedToMasterModeForcefully", k)
	}
	return nil
}

func (item *EngineSwitchedToMasterModeForcefully) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '{')
	if item.BytesTruncated != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"bytes_truncated":`...)
		w = basictl.JSONWriteInt64(w, item.BytesTruncated)
	}
	return append(w, '}'), nil
}

func (item *EngineSwitchedToMasterModeForcefully) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *EngineSwitchedToMasterModeForcefully) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("engine.switchedToMasterModeForcefully", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("engine.switchedToMasterModeForcefully", err.Error())
	}
	return nil
}

func (item EngineSwitchedToReplicaMode) AsUnion() EngineSwitchMasterReplicaModeResultUnion {
	var ret EngineSwitchMasterReplicaModeResultUnion
	ret.SetSwitchedToReplicaMode()
	return ret
}

// AsUnion will be here
type EngineSwitchedToReplicaMode struct {
}

func (EngineSwitchedToReplicaMode) TLName() string { return "engine.switchedToReplicaMode" }
func (EngineSwitchedToReplicaMode) TLTag() uint32  { return 0xad642a0b }

func (item *EngineSwitchedToReplicaMode) Reset()                         {}
func (item *EngineSwitchedToReplicaMode) Read(w []byte) ([]byte, error)  { return w, nil }
func (item *EngineSwitchedToReplicaMode) Write(w []byte) ([]byte, error) { return w, nil }
func (item *EngineSwitchedToReplicaMode) ReadBoxed(w []byte) ([]byte, error) {
	return basictl.NatReadExactTag(w, 0xad642a0b)
}
func (item *EngineSwitchedToReplicaMode) WriteBoxed(w []byte) ([]byte, error) {
	return basictl.NatWrite(w, 0xad642a0b), nil
}

func (item EngineSwitchedToReplicaMode) String() string {
	w, err := item.WriteJSON(nil)
	if err != nil {
		return err.Error()
	}
	return string(w)
}

func EngineSwitchedToReplicaMode__ReadJSON(item *EngineSwitchedToReplicaMode, j interface{}) error {
	return item.readJSON(j)
}
func (item *EngineSwitchedToReplicaMode) readJSON(j interface{}) error {
	_jm, _ok := j.(map[string]interface{})
	if j != nil && !_ok {
		return ErrorInvalidJSON("engine.switchedToReplicaMode", "expected json object")
	}
	for k := range _jm {
		return ErrorInvalidJSONExcessElement("engine.switchedToReplicaMode", k)
	}
	return nil
}

func (item *EngineSwitchedToReplicaMode) WriteJSON(w []byte) (_ []byte, err error) {
	w = append(w, '{')
	return append(w, '}'), nil
}

func (item *EngineSwitchedToReplicaMode) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil)
}

func (item *EngineSwitchedToReplicaMode) UnmarshalJSON(b []byte) error {
	j, err := JsonBytesToInterface(b)
	if err != nil {
		return ErrorInvalidJSON("engine.switchedToReplicaMode", err.Error())
	}
	if err = item.readJSON(j); err != nil {
		return ErrorInvalidJSON("engine.switchedToReplicaMode", err.Error())
	}
	return nil
}