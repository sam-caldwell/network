package network

import "unsafe"

// isLittleEndian checks if the system is little-endian.
func isLittleEndian() bool {
	var i uint16 = 0x0102
	u8 := (*[2]byte)(unsafe.Pointer(&i))
	return u8[0] != 0x01
}
