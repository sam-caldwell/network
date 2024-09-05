package core

import (
	"errors"
)

// DeserializeTcSfqQoptV1 - Safely converts a byte slice into a TcSfqQoptV1 struct.
// This method ensures that the length of the input byte slice matches the expected size of TcSfqQoptV1.
//
// If the length of the input byte slice is less than the size of TcSfqQoptV1, it returns an error.
func DeserializeTcSfqQoptV1(b []byte) (*TcSfqQoptV1, error) {
	if len(b) < SizeofTcSfqQoptV1 {
		return nil, errors.New("DeserializeTcSfqQoptV1: byte slice too short")
	}

	msg := &TcSfqQoptV1{}

	// Deserialize the base TcSfqQopt structure
	err := msg.TcSfqQopt.Deserialize(b[:SizeofTcSfqQopt])
	if err != nil {
		return nil, err
	}

	// Manually copy the additional fields in TcSfqQoptV1
	msg.Depth = NativeEndian.Uint32(b[SizeofTcSfqQopt:])
	msg.HeadDrop = NativeEndian.Uint32(b[SizeofTcSfqQopt+4:])
	msg.Limit = NativeEndian.Uint32(b[SizeofTcSfqQopt+8:])
	msg.QthMin = NativeEndian.Uint32(b[SizeofTcSfqQopt+12:])
	msg.QthMax = NativeEndian.Uint32(b[SizeofTcSfqQopt+16:])
	msg.Wlog = b[SizeofTcSfqQopt+20]
	msg.Plog = b[SizeofTcSfqQopt+21]
	msg.ScellLog = b[SizeofTcSfqQopt+22]
	msg.Flags = b[SizeofTcSfqQopt+23]
	msg.MaxP = NativeEndian.Uint32(b[SizeofTcSfqQopt+24:])

	// Deserialize the TcSfqRedStats structure
	t, err := DeserializeTcSfqRedStats(b[SizeofTcSfqQopt+28:])
	if err != nil {
		return nil, err
	}
	msg.TcSfqRedStats = *t
	return msg, err
}
