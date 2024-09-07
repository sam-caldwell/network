package core

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// DeserializeXfrmAlgoAEAD safely deserializes a byte slice into an XfrmAlgoAEAD structure.
// It returns the deserialized XfrmAlgoAEAD structure or an error if deserialization fails.
func DeserializeXfrmAlgoAEAD(b []byte) (*XfrmAlgoAEAD, error) {
	if len(b) < 72 { // 64 bytes for AlgName + 4 for AlgKeyLen + 4 for AlgICVLen
		return nil, errors.New("byte slice too small to deserialize XfrmAlgoAEAD")
	}

	// Create a reader for the byte slice
	reader := bytes.NewReader(b)

	// Create a new XfrmAlgoAEAD structure
	msg := &XfrmAlgoAEAD{}

	// Deserialize AlgName (64 bytes)
	if err := binary.Read(reader, binary.BigEndian, &msg.AlgName); err != nil {
		return nil, err
	}

	// Deserialize AlgKeyLen (4 bytes)
	if err := binary.Read(reader, binary.BigEndian, &msg.AlgKeyLen); err != nil {
		return nil, err
	}

	// Deserialize AlgICVLen (4 bytes)
	if err := binary.Read(reader, binary.BigEndian, &msg.AlgICVLen); err != nil {
		return nil, err
	}

	// Calculate the length of AlgKey based on AlgKeyLen
	algKeyLenBytes := int(msg.AlgKeyLen / 8) // Convert bit length to byte length

	// Check if the byte slice is large enough for AlgKey
	if len(b) < 72+algKeyLenBytes {
		return nil, errors.New("byte slice too small for AlgKey")
	}

	// Deserialize AlgKey (variable length)
	msg.AlgKey = make([]byte, algKeyLenBytes)
	if _, err := reader.Read(msg.AlgKey); err != nil {
		return nil, err
	}

	return msg, nil
}
