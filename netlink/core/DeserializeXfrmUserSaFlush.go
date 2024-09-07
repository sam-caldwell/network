package core

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// DeserializeXfrmUserSaFlush safely deserializes a byte slice into an XfrmUserSaFlush structure.
// It returns the deserialized structure or an error if deserialization fails.
func DeserializeXfrmUserSaFlush(b []byte) (*XfrmUserSaFlush, error) {
	if len(b) < SizeOfXfrmUserSaFlush {
		return nil, errors.New("byte slice too small to deserialize XfrmUserSaFlush")
	}

	// Create a new XfrmUserSaFlush structure
	msg := &XfrmUserSaFlush{}

	// Create a reader for the byte slice
	reader := bytes.NewReader(b)

	// Deserialize Proto (uint8)
	if err := binary.Read(reader, binary.BigEndian, &msg.Proto); err != nil {
		return nil, err
	}

	// Return the deserialized XfrmUserSaFlush structure
	return msg, nil
}
