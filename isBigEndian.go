package network

import "unsafe"

// isBigEndian checks if the system is big-endian.
func isBigEndian() bool {
	var i uint16 = 0x0102
	u8 := (*[2]byte)(unsafe.Pointer(&i))
	return u8[0] == 0x01
}
