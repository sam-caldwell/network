package core

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// DeserializeTcPeditKey safely deserializes a byte slice into a TcPeditKey structure.
// It checks the length of the byte slice to ensure that it is at least the size of TcPeditKey
// and uses the encoding/binary package to read the fields safely.
//
// Returns an error if the byte slice is too small or cannot be properly deserialized.
func DeserializeTcPeditKey(b []byte) (*TcPeditKey, error) {
	// Check if the byte slice is large enough to hold a TcPeditKey
	if len(b) < SizeOfTcPeditKey {
		return nil, errors.New("byte slice too small to deserialize TcPeditKey")
	}

	// Create a new TcPeditKey struct
	key := TcPeditKey{}

	// Create a reader for the byte slice
	reader := bytes.NewReader(b)

	// Safely deserialize the fields using binary.Read
	if err := binary.Read(reader, binary.BigEndian, &key.Mask); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.BigEndian, &key.Val); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.BigEndian, &key.Off); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.BigEndian, &key.At); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.BigEndian, &key.OffMask); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.BigEndian, &key.Shift); err != nil {
		return nil, err
	}

	// Return the deserialized TcPeditKey
	return &key, nil
}
