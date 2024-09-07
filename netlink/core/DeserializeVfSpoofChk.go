package core

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// DeserializeVfSpoofchk - create VfSpoofchk structure from a byte slice. It returns an error if the byte slice is
// not the correct length.
func DeserializeVfSpoofchk(b []byte) (*VfSpoofchk, error) {

	var msg VfSpoofchk

	if len(b) != SizeOfVfSpoofchk {
		return nil, errors.New("invalid length for VfSpoofchk")
	}

	reader := bytes.NewReader(b)

	if err := binary.Read(reader, NativeEndian, &msg); err != nil {
		return nil, err
	}

	return &msg, nil

}
