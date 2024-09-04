package core

import (
	"unsafe"
)

// Len - return the size of the RtMsg struct in bytes
func (msg *RtMsg) Len() int {

	return int(unsafe.Sizeof(RtMsg{}))

}
