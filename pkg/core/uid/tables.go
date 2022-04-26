// Copyright (c) 2021 by library authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uid

// TableUID represents a UID for a specific table and is derived forom UID
type TableUID UID

var (
	Base_TableTable         = TableUID{0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00}
	Base_MethodIDTable      = TableUID{0x00, 0x00, 0x00, 0x06, 0x00, 0x00, 0x00, 0x00}
	Base_AccessControlTable = TableUID{0x00, 0x00, 0x00, 0x07, 0x00, 0x00, 0x00, 0x00}
	Admin_TPerInfoTable     = TableUID{0x00, 0x00, 0x02, 0x01, 0x00, 0x00, 0x00, 0x00}
	Admin_C_PINTable        = TableUID{0x00, 0x00, 0x00, 0x0B, 0x00, 0x00, 0x00, 0x00}
	Locking_LockingTable    = TableUID{0x00, 0x00, 0x08, 0x02, 0x00, 0x00, 0x00, 0x00}
	Locking_MBRTable        = TableUID{0x00, 0x00, 0x08, 0x04, 0x00, 0x00, 0x00, 0x00}
)

func (t *TableUID) Row(uid [4]byte) RowUID {
	return [8]byte{t[0], t[1], t[2], t[3], uid[0], uid[1], uid[2], uid[3]}
}

func Base_TableRowForTable(tid TableUID) RowUID {
	return Base_TableTable.Row([4]byte{tid[0], tid[1], tid[2], tid[3]})
}
