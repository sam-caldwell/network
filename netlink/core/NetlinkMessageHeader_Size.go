package core

import (
	"golang.org/x/sys/unix"
	"unsafe"
)

// NetlinkMessageHeaderSize - Size of unix.NlMsghdr struct
const NetlinkMessageHeaderSize = int(unsafe.Sizeof(unix.NlMsghdr{}))
