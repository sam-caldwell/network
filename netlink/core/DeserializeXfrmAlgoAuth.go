package core

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// DeserializeXfrmAlgoAuth - safely deserializes a byte slice into an XfrmAlgoAuth structure.
// It returns the deserialized XfrmAlgoAuth structure or an error if deserialization fails.
func DeserializeXfrmAlgoAuth(b []byte) (*XfrmAlgoAuth, error) {
	if len(b) < 72 { // 64 bytes for AlgName + 4 for AlgKeyLen + 4 for AlgTruncLen
		return nil, errors.New("byte slice too small to deserialize XfrmAlgoAuth")
	}

	// Create a reader for the byte slice
	reader := bytes.NewReader(b)

	// Create a new XfrmAlgoAuth structure
	msg := &XfrmAlgoAuth{}

	// Deserialize AlgName (64 bytes)
	if err := binary.Read(reader, binary.BigEndian, &msg.AlgName); err != nil {
		return nil, err
	}

	// Deserialize AlgKeyLen (4 bytes)
	if err := binary.Read(reader, binary.BigEndian, &msg.AlgKeyLen); err != nil {
		return nil, err
	}

	// Deserialize AlgTruncLen (4 bytes)
	if err := binary.Read(reader, binary.BigEndian, &msg.AlgTruncLen); err != nil {
		return nil, err
	}

	// Determine the size of AlgKey based on AlgKeyLen
	algKeyLenBytes := int(msg.AlgKeyLen / 8) // Convert bit length to byte length
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
