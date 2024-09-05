package core

import (
	"errors"
	"unsafe"
)

// DeserializeTcCsum - Converts a byte slice into a TcCsum struct.
// The function checks the length of the byte slice to ensure it matches the size of the TcCsum struct.
// Returns the TcCsum object or an error if the input size is incorrect.
//
// The TcCsum struct is typically used in conjunction with netlink messages for checksum action in Linux traffic control (tc).
//
// References:
// - Linux Kernel Source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_csum.h
func DeserializeTcCsum(b []byte) (*TcCsum, error) {
	if len(b) < SizeofTcCsum {
		return nil, errors.New("DeserializeTcCsum: byte slice too short")
	}

	// Safely copy bytes into the TcCsum structure
	var msg TcCsum
	copy((*(*[SizeofTcCsum]byte)(unsafe.Pointer(&msg)))[:], b[:SizeofTcCsum])

	return &msg, nil
}
