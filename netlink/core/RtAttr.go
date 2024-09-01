package core

import (
	"golang.org/x/sys/unix"
)

// RtAttr - Extend unix.RtAttr to handle data and children
type RtAttr struct {
	unix.RtAttr
	Data     []byte
	children []NetlinkRequestData
}
