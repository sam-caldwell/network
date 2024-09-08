package core

import (
	"encoding/binary"
	"errors"
)

// DeserializeSkbEdit safely deserializes a byte slice into a TcSkbEdit struct.
// The function ensures the byte slice is of the correct size and uses binary deserialization.
func DeserializeSkbEdit(b []byte) (*TcSkbEdit, error) {
	if len(b) < SizeOfTcSkbEdit {
		return nil, errors.New("byte slice too short")
	}

	msg := &TcSkbEdit{}

	// Deserialize each field in the TcSkbEdit struct.
	msg.Index = binary.LittleEndian.Uint32(b[0:4])
	msg.Capab = binary.LittleEndian.Uint32(b[4:8])
	msg.Action = int32(binary.LittleEndian.Uint32(b[8:12]))
	msg.Refcnt = int32(binary.LittleEndian.Uint32(b[12:16]))
	msg.Bindcnt = int32(binary.LittleEndian.Uint32(b[16:20]))

	return msg, nil
}
