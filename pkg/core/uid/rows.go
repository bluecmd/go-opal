// Copyright (c) 2021 by library authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uid

var (
	Admin_TPerInfoObj   RowUID = Admin_TPerInfoTable.Row([4]byte{0x00, 0x03, 0x00, 0x01})
	Admin_C_PIN_MSIDRow RowUID = Admin_C_PINTable.Row([4]byte{0x00, 0x00, 0x84, 0x02})
	Admin_C_PIN_SIDRow  RowUID = Admin_C_PINTable.Row([4]byte{0x00, 0x00, 0x00, 0x01})

	LockingInfoObj           RowUID = [8]byte{0x00, 0x00, 0x08, 0x01, 0x00, 0x00, 0x00, 0x01}
	EnterpriseLockingInfoObj RowUID = [8]byte{0x00, 0x00, 0x08, 0x01, 0x00, 0x00, 0x00, 0x00}
	MBRControlObj            RowUID = [8]byte{0x00, 0x00, 0x08, 0x03, 0x00, 0x00, 0x00, 0x01}

	LockingRange1 RowUID = [8]byte{0x00, 0x00, 0x08, 0x02, 0x00, 0x03, 0x00, 0x01}
)
