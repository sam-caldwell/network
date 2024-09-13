package core

import "unsafe"

// IfInfoMsgSize -  Size of the IfInfoMsg struct
const IfInfoMsgSize = int(unsafe.Sizeof(IfInfoMsg{}))
