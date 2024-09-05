package core

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// DeserializeVfRssQueryEn - create a VfRssQueryEn structure from a byte slice. Return error if the slice not the
// correct length.
func DeserializeVfRssQueryEn(b []byte) (*VfRssQueryEn, error) {
	if len(b) != SizeofVfRssQueryEn {
		return nil, errors.New("invalid length for VfRssQueryEn")
	}

	var msg VfRssQueryEn
	reader := bytes.NewReader(b)
	if err := binary.Read(reader, NativeEndian, &msg); err != nil {
		return nil, err
	}
	return &msg, nil
}
