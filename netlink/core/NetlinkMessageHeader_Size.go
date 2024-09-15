package core

import (
	"unsafe"
)

// NetlinkMessageHeaderSize - Size of NlMsghdr struct
const NetlinkMessageHeaderSize = int(unsafe.Sizeof(NlMsghdr{}))

const (
	NlMsghdrSize = int(unsafe.Sizeof(NlMsghdr{}))
)
