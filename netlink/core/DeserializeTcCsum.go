package core

import (
	"encoding/binary"
	"errors"
)

// DeserializeTcCsum - Converts a byte slice into a TcCsum struct.
// This function validates the size of the byte slice and safely deserializes it into a TcCsum object.
// Returns the TcCsum object or an error if the input size is incorrect.
//
// The TcCsum struct is typically used in conjunction with netlink messages for checksum action in Linux traffic control (tc).
//
// References:
// - Linux Kernel Source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_csum.h
func DeserializeTcCsum(b []byte) (*TcCsum, error) {
	if len(b) < SizeOfTcCsum {
		return nil, errors.New("DeserializeTcCsum: byte slice too short")
	}

	// Create an empty TcCsum structure
	msg := &TcCsum{}

	// Deserialize TcGen fields
	msg.TcGen.Index = binary.LittleEndian.Uint32(b[0:4])
	msg.TcGen.Capab = binary.LittleEndian.Uint32(b[4:8])
	msg.TcGen.Action = int32(binary.LittleEndian.Uint32(b[8:12]))
	msg.TcGen.Refcnt = int32(binary.LittleEndian.Uint32(b[12:16]))
	msg.TcGen.Bindcnt = int32(binary.LittleEndian.Uint32(b[16:20]))

	// Deserialize UpdateFlags field
	msg.UpdateFlags = binary.LittleEndian.Uint32(b[20:24])

	return msg, nil
}
