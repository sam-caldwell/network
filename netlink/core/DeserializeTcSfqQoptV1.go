package core

import (
	"errors"
)

// DeserializeTcSfqQoptV1 - Safely converts a byte slice into a TcSfqQoptV1 struct.
// This method ensures that the length of the input byte slice matches the expected size of TcSfqQoptV1.
//
// If the length of the input byte slice is less than the size of TcSfqQoptV1, it returns an error.
func DeserializeTcSfqQoptV1(b []byte) (*TcSfqQoptV1, error) {
	if len(b) < SizeOfTcSfqQoptV1 {
		return nil, errors.New("DeserializeTcSfqQoptV1: byte slice too short")
	}

	msg := &TcSfqQoptV1{}

	// Deserialize the base TcSfqQopt structure
	err := msg.TcSfqQopt.Deserialize(b[:SizeOfTcSfqQopt])
	if err != nil {
		return nil, err
	}

	// Manually copy the additional fields in TcSfqQoptV1
	msg.Depth = NativeEndian.Uint32(b[SizeOfTcSfqQopt:])
	msg.HeadDrop = NativeEndian.Uint32(b[SizeOfTcSfqQopt+4:])
	msg.Limit = NativeEndian.Uint32(b[SizeOfTcSfqQopt+8:])
	msg.QthMin = NativeEndian.Uint32(b[SizeOfTcSfqQopt+12:])
	msg.QthMax = NativeEndian.Uint32(b[SizeOfTcSfqQopt+16:])
	msg.Wlog = b[SizeOfTcSfqQopt+20]
	msg.Plog = b[SizeOfTcSfqQopt+21]
	msg.ScellLog = b[SizeOfTcSfqQopt+22]
	msg.Flags = b[SizeOfTcSfqQopt+23]
	msg.MaxP = NativeEndian.Uint32(b[SizeOfTcSfqQopt+24:])

	// Deserialize the TcSfqRedStats structure
	t, err := DeserializeTcSfqRedStats(b[SizeOfTcSfqQopt+28:])
	if err != nil {
		return nil, err
	}
	msg.TcSfqRedStats = *t
	return msg, err
}
