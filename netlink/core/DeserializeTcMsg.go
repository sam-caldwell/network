package core

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// DeserializeTcMsg - Given []byte input, deserialize to TcMsg object and return by reference.
// This version uses binary.Read for safer deserialization.
func DeserializeTcMsg(b []byte) (*TcMsg, error) {
	if len(b) < SizeofTcMsg {
		return nil, errors.New("DeserializeTcMsg: input byte slice is too small")
	}

	buf := bytes.NewReader(b)
	var msg TcMsg

	err := binary.Read(buf, NativeEndian, &msg)
	if err != nil {
		return nil, err
	}

	return &msg, nil
}
