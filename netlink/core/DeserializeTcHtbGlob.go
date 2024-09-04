package core

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// DeserializeTcHtbGlob - Safely deserializes a byte slice into a TcHtbGlob struct.
func DeserializeTcHtbGlob(b []byte) (*TcHtbGlob, error) {
	if len(b) < SizeofTcHtbGlob {
		return nil, errors.New("byte slice is too short to contain a TcHtbGlob")
	}

	buf := bytes.NewReader(b)
	var msg TcHtbGlob

	// Using LittleEndian decoding, as Linux generally uses this on most platforms.
	err := binary.Read(buf, binary.LittleEndian, &msg)
	if err != nil {
		return nil, err
	}

	return &msg, nil
}
