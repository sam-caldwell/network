package core

import (
	"errors"
)

// DeserializeSkbEdit safely deserializes a byte slice into a TcSkbEdit struct.
// The function ensures the byte slice is of the correct size and uses binary deserialization.
func DeserializeSkbEdit(b []byte) (*TcSkbEdit, error) {
	if len(b) < SizeofTcSkbEdit {
		return nil, errors.New("DeserializeSkbEdit: byte slice too short")
	}

	msg := &TcSkbEdit{}

	// Deserialize each field in the TcSkbEdit struct.
	msg.Index = NativeEndian.Uint32(b[0:4])
	msg.Capab = NativeEndian.Uint32(b[4:8])
	msg.Action = int32(NativeEndian.Uint32(b[8:12]))
	msg.Refcnt = int32(NativeEndian.Uint32(b[12:16]))
	msg.Bindcnt = int32(NativeEndian.Uint32(b[16:20]))

	return msg, nil
}
