// Copyright (c) 2021 by library authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Implements TCG Storage Core Table operations

package table

import (
	"fmt"

	"github.com/open-source-firmware/go-tcg-storage/pkg/core"
	"github.com/open-source-firmware/go-tcg-storage/pkg/core/stream"
	"github.com/open-source-firmware/go-tcg-storage/pkg/core/uid"
)

var (
	Admin_C_PIN_ColumnPIN         uint = 3
	Admin_SP_ColumnLifeCycleState uint = 6
)

func Admin_C_PIN_MSID_GetPIN(s *core.Session) ([]byte, error) {
	val, err := GetCell(s, uid.Admin_C_PIN_MSIDRow, Admin_C_PIN_ColumnPIN, "PIN")
	if err != nil {
		return nil, err
	}
	pin, ok := val.([]byte)
	if !ok {
		return nil, fmt.Errorf("malformed PIN column")
	}
	return pin, nil
}

type Admin_TPerInfoRow struct {
	UID                     uid.RowUID
	Bytes                   *uint64
	GUDID                   *[12]byte
	Generation              *uint32
	FirmwareVersion         *uint32
	ProtocolVersion         *uint32
	SpaceForIssuance        *uint64
	SSC                     []string
	ProgrammaticResetEnable *bool
}

func Admin_TPerInfo(s *core.Session) (map[uid.RowUID]Admin_TPerInfoRow, error) {
	res := map[uid.RowUID]Admin_TPerInfoRow{}
	val, err := GetFullRow(s, uid.Admin_TPerInfoObj)
	if err != nil {
		return nil, err
	}

	row := Admin_TPerInfoRow{}
	for col, val := range val {
		switch col {
		case "0", "UID":
			v, ok := val.([]byte)
			if !ok {
				return nil, core.ErrMalformedMethodResponse
			}
			copy(row.UID[:], v[:8])
		case "1":
			v, ok := val.(uint)
			if !ok {
				return nil, core.ErrMalformedMethodResponse
			}
			vv := uint64(v)
			row.Bytes = &vv
		case "2":
			v, ok := val.([]byte)
			if !ok {
				return nil, core.ErrMalformedMethodResponse
			}
			vv := [12]byte{}
			copy(vv[:], v)
			row.GUDID = &vv
		case "3":
			v, ok := val.(uint)
			if !ok {
				return nil, core.ErrMalformedMethodResponse
			}
			vv := uint32(v)
			row.Generation = &vv
		case "4":
			v, ok := val.(uint)
			if !ok {
				return nil, core.ErrMalformedMethodResponse
			}
			vv := uint32(v)
			row.FirmwareVersion = &vv
		case "5":
			v, ok := val.(uint)
			if !ok {
				return nil, core.ErrMalformedMethodResponse
			}
			vv := uint32(v)
			row.ProtocolVersion = &vv
		case "6":
			v, ok := val.(uint)
			if !ok {
				return nil, core.ErrMalformedMethodResponse
			}
			vv := uint64(v)
			row.SpaceForIssuance = &vv
		case "7":
			vl, ok := val.(stream.List)
			if !ok {
				vl = stream.List{val}
			}
			for _, val := range vl {
				v, ok := val.([]byte)
				if !ok {
					return nil, core.ErrMalformedMethodResponse
				}
				row.SSC = append(row.SSC, string(v))
			}
		case "8":
			v, ok := val.(uint)
			if !ok {
				return nil, core.ErrMalformedMethodResponse
			}
			var vv bool
			if v > 0 {
				vv = true
			}
			row.ProgrammaticResetEnable = &vv
		}
	}

	res[row.UID] = row
	return res, nil
}

type LifeCycleState int

func Admin_SP_GetLifeCycleState(s *core.Session, spid uid.SPID) (LifeCycleState, error) {
	val, err := GetCell(s, uid.RowUID(spid), Admin_SP_ColumnLifeCycleState, "LifeCycleState")
	if err != nil {
		return -1, err
	}
	v, ok := val.(uint)
	if !ok {
		return -1, fmt.Errorf("malformed LifeCycleState column")
	}
	return LifeCycleState(v), nil
}
