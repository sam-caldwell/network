package core

import (
	"bytes"
	"encoding/binary"
)

// Serialize - Converts the TcMirred structure into a byte slice safely.
// This method serializes each field of the structure individually using
// the binary encoding package to ensure portability across different architectures.
//
// The TcMirred structure is used in the Linux Traffic Control system
// to mirror or redirect traffic to specific network interfaces.
//
// References:
// - https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_mirred.h
func (msg *TcMirred) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	// Safely serialize each field using binary encoding in little endian format
	if err := binary.Write(buf, NativeEndian, msg.TcGen.Index); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, NativeEndian, msg.TcGen.Capab); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, NativeEndian, msg.TcGen.Action); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, NativeEndian, msg.TcGen.Refcnt); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, NativeEndian, msg.TcGen.Bindcnt); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, NativeEndian, msg.Eaction); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, NativeEndian, msg.Ifindex); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
