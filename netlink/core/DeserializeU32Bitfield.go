package core

import (
	"encoding/binary"
	"errors"
)

// DeserializeUint32Bitfield - deserialize a Uint32Bitfield from a byte slice.
func DeserializeUint32Bitfield(data []byte) (*Uint32Bitfield, error) {
	if len(data) < SizeofUint32Bitfield {
		return nil, errors.New("data slice too short")
	}

	// Create a new Uint32Bitfield and use binary package to handle endian conversion
	bitfield := &Uint32Bitfield{
		Value:    NativeEndian.Uint32(data[0:4]),
		Selector: NativeEndian.Uint32(data[4:8]),
	}

	return bitfield, nil
}
