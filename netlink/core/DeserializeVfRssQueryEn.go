package core

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// DeserializeVfRssQueryEn - create a VfRssQueryEn structure from a byte slice. Return error if the slice not the
// correct length.
func DeserializeVfRssQueryEn(b []byte) (*VfRssQueryEn, error) {
	if len(b) != SizeOfVfRssQueryEn {
		return nil, errors.New("input too short")
	}

	var msg VfRssQueryEn
	reader := bytes.NewReader(b)
	if err := binary.Read(reader, NativeEndian, &msg); err != nil {
		return nil, err
	}
	return &msg, nil
}
