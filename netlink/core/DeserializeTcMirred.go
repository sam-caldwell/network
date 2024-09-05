package core

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// DeserializeTcMirred - Safely deserialize a byte slice into a TcMirred structure.
//
// The function checks the size of the byte slice and uses binary encoding to
// populate each field of the TcMirred struct. This ensures safety and portability.
//
// The TcMirred structure is used in the Linux Traffic Control system for
// mirroring and redirecting traffic to network interfaces.
//
// References:
// - https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_mirred.h
func DeserializeTcMirred(b []byte) (*TcMirred, error) {
	if len(b) < SizeofTcMirred {
		return nil, errors.New("DeserializeTcMirred: byte slice too short")
	}

	buf := bytes.NewReader(b)
	msg := &TcMirred{}

	// Safely deserialize each field using binary.Read in little endian format
	if err := binary.Read(buf, NativeEndian, &msg.TcGen.Index); err != nil {
		return nil, err
	}
	if err := binary.Read(buf, NativeEndian, &msg.TcGen.Capab); err != nil {
		return nil, err
	}
	if err := binary.Read(buf, NativeEndian, &msg.TcGen.Action); err != nil {
		return nil, err
	}
	if err := binary.Read(buf, NativeEndian, &msg.TcGen.Refcnt); err != nil {
		return nil, err
	}
	if err := binary.Read(buf, NativeEndian, &msg.TcGen.Bindcnt); err != nil {
		return nil, err
	}
	if err := binary.Read(buf, NativeEndian, &msg.Eaction); err != nil {
		return nil, err
	}
	if err := binary.Read(buf, NativeEndian, &msg.Ifindex); err != nil {
		return nil, err
	}

	return msg, nil
}
