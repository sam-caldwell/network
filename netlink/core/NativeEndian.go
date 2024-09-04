package core

import (
	"encoding/binary"
	"unsafe"
)

// NativeEndian - implementation of NativeEndian
var NativeEndian binary.ByteOrder

// initialize the NativeEndian variable at startup
func init() {
	var x uint32 = 0x01020304
	if *(*byte)(unsafe.Pointer(&x)) == 0x01 {
		NativeEndian = binary.BigEndian
	} else {
		NativeEndian = binary.LittleEndian
	}
}
