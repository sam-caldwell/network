package core

import (
	"encoding/binary"
	"testing"
	"unsafe"
)

func TestNativeEndian(t *testing.T) {
	// Determine the expected endianness
	var expected binary.ByteOrder
	var x uint32 = 0x01020304
	if *(*byte)(unsafe.Pointer(&x)) == 0x01 {
		expected = binary.BigEndian
	} else {
		expected = binary.LittleEndian
	}

	// Get the native endianness from the function

	// Check if the actual endianness matches the expected endianness
	if nativeEndian != expected {
		t.Errorf("NativeEndian() = %v; want %v", nativeEndian, expected)
	}
}
