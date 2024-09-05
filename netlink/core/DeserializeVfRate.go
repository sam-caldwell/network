package core

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// DeserializeVfRate - create a VfRate struct from a byte slice. It returns an error if the byte slice is not
// the correct length.
func DeserializeVfRate(b []byte) (*VfRate, error) {
	var v VfRate

	// Each uint32 is 4 bytes, so 3 uint32 fields require 12 bytes
	if len(b) != 12 {
		return nil, errors.New("invalid length for VfRate")
	}

	reader := bytes.NewReader(b)
	if err := binary.Read(reader, NativeEndian, &v); err != nil {
		return nil, err
	}

	return &v, nil
}
