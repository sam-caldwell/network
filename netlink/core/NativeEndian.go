package core

import (
	"encoding/binary"
	"unsafe"
)

var nativeEndian binary.ByteOrder

// initialize the nativeEndian variable at startup
func init() {
	var x uint32 = 0x01020304
	if *(*byte)(unsafe.Pointer(&x)) == 0x01 {
		nativeEndian = binary.BigEndian
	} else {
		nativeEndian = binary.LittleEndian
	}
}

// NativeEndian returns the native endianness of the system.
func NativeEndian() binary.ByteOrder {
	return nativeEndian
}
