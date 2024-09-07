package core

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// DeserializeVfLinkState - create a VfLinkState structure from a byte slice.
// It returns an error if the byte slice is not the correct length.
func DeserializeVfLinkState(b []byte) (*VfLinkState, error) {

	var msg VfLinkState

	if len(b) != SizeOfVfLinkState {
		return nil, errors.New("invalid length for VfLinkState")
	}

	reader := bytes.NewReader(b)

	if err := binary.Read(reader, NativeEndian, &msg); err != nil {
		return nil, err
	}

	return &msg, nil

}
