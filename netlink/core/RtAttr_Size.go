package core

import (
	"golang.org/x/sys/unix"
	"math"
	"unsafe"
)

const (
	// SizeOfUnixRtAttr - The size of a unix.RtAttr struct
	SizeOfUnixRtAttr = int(unsafe.Sizeof(unix.RtAttr{}))
	MaxRtAttrSize    = int(math.MaxUint16)
)
