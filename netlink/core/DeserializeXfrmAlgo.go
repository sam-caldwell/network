package core

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// DeserializeXfrmAlgo safely deserializes a byte slice into an XfrmAlgo structure.
// It returns the deserialized XfrmAlgo structure, or an error if deserialization fails.
func DeserializeXfrmAlgo(b []byte) (*XfrmAlgo, error) {
	// Check if the byte slice is large enough to contain the fixed-size part of the XfrmAlgo structure.
	if len(b) < 68 {
		return nil, errors.New("byte slice too small to deserialize XfrmAlgo")
	}

	// Create a new XfrmAlgo structure.
	ret := XfrmAlgo{}

	// Copy the AlgName (64 bytes).
	copy(ret.AlgName[:], b[0:64])

	// Deserialize AlgKeyLen (4 bytes).
	reader := bytes.NewReader(b[64:68])
	if err := binary.Read(reader, binary.BigEndian, &ret.AlgKeyLen); err != nil {
		return nil, err
	}

	// Check if the remaining bytes are enough for AlgKey.
	expectedLen := int(ret.AlgKeyLen / 8) // AlgKeyLen is in bits, divide by 8 for bytes.
	if len(b) < 68+expectedLen {
		return nil, errors.New("byte slice too small for AlgKey")
	}

	// Copy the AlgKey (variable length based on AlgKeyLen).
	ret.AlgKey = make([]byte, expectedLen)
	copy(ret.AlgKey, b[68:68+expectedLen])

	// Return the deserialized XfrmAlgo structure.
	return &ret, nil
}
