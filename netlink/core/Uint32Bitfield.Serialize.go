package core

import (
	"errors"
)

// Serialize - serialize the Uint32Bitfield struct into a byte slice
func (a *Uint32Bitfield) Serialize() ([]byte, error) {
	// Allocate a byte slice of the required size
	data := make([]byte, SizeofUint32Bitfield)

	// Check if SizeofUint32Bitfield matches the expected size
	if len(data) < 8 {
		return nil, errors.New("invalid size for Uint32Bitfield serialization")
	}

	// Use binary package to handle endian conversion
	NativeEndian.PutUint32(data[0:4], a.Value)
	NativeEndian.PutUint32(data[4:8], a.Selector)

	return data, nil
}
