package core

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// DeserializeVfSpoofChk - create VfSpoofchk structure from a byte slice. It returns an error if the byte slice is
// not the correct length.
func DeserializeVfSpoofChk(b []byte) (*VfSpoofchk, error) {

	var msg VfSpoofchk

	if len(b) != SizeOfVfSpoofchk {
		return nil, errors.New("input too short")
	}

	reader := bytes.NewReader(b)

	if err := binary.Read(reader, NativeEndian, &msg); err != nil {
		return nil, err
	}

	return &msg, nil

}
