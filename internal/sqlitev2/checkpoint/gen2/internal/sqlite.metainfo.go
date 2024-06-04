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

type SqliteMetainfo struct {
	FieldMask uint32
	Offset    int64
}

func (SqliteMetainfo) TLName() string { return "sqlite.metainfo" }
func (SqliteMetainfo) TLTag() uint32  { return 0x9286affa }

func (item *SqliteMetainfo) Reset() {
	item.FieldMask = 0
	item.Offset = 0
}

func (item *SqliteMetainfo) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldMask); err != nil {
		return w, err
	}
	return basictl.LongRead(w, &item.Offset)
}

// This method is general version of Write, use it instead!
func (item *SqliteMetainfo) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *SqliteMetainfo) Write(w []byte) []byte {
	w = basictl.NatWrite(w, item.FieldMask)
	w = basictl.LongWrite(w, item.Offset)
	return w
}

func (item *SqliteMetainfo) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0x9286affa); err != nil {
		return w, err
	}
	return item.Read(w)
}

// This method is general version of WriteBoxed, use it instead!
func (item *SqliteMetainfo) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *SqliteMetainfo) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0x9286affa)
	return item.Write(w)
}

func (item SqliteMetainfo) String() string {
	return string(item.WriteJSON(nil))
}

func (item *SqliteMetainfo) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propFieldMaskPresented bool
	var propOffsetPresented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "field_mask":
				if propFieldMaskPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("sqlite.metainfo", "field_mask")
				}
				if err := Json2ReadUint32(in, &item.FieldMask); err != nil {
					return err
				}
				propFieldMaskPresented = true
			case "offset":
				if propOffsetPresented {
					return ErrorInvalidJSONWithDuplicatingKeys("sqlite.metainfo", "offset")
				}
				if err := Json2ReadInt64(in, &item.Offset); err != nil {
					return err
				}
				propOffsetPresented = true
			default:
				return ErrorInvalidJSONExcessElement("sqlite.metainfo", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propFieldMaskPresented {
		item.FieldMask = 0
	}
	if !propOffsetPresented {
		item.Offset = 0
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *SqliteMetainfo) WriteJSONGeneral(w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(true, false, w), nil
}

func (item *SqliteMetainfo) WriteJSON(w []byte) []byte {
	return item.WriteJSONOpt(true, false, w)
}
func (item *SqliteMetainfo) WriteJSONOpt(newTypeNames bool, short bool, w []byte) []byte {
	w = append(w, '{')
	backupIndexFieldMask := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"field_mask":`...)
	w = basictl.JSONWriteUint32(w, item.FieldMask)
	if (item.FieldMask != 0) == false {
		w = w[:backupIndexFieldMask]
	}
	backupIndexOffset := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"offset":`...)
	w = basictl.JSONWriteInt64(w, item.Offset)
	if (item.Offset != 0) == false {
		w = w[:backupIndexOffset]
	}
	return append(w, '}')
}

func (item *SqliteMetainfo) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *SqliteMetainfo) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return ErrorInvalidJSON("sqlite.metainfo", err.Error())
	}
	return nil
}