package core

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// DeserializeXfrmUsersaFlush safely deserializes a byte slice into an XfrmUsersaFlush structure.
// It returns the deserialized structure or an error if deserialization fails.
func DeserializeXfrmUsersaFlush(b []byte) (*XfrmUsersaFlush, error) {
	if len(b) < SizeofXfrmUsersaFlush {
		return nil, errors.New("byte slice too small to deserialize XfrmUsersaFlush")
	}

	// Create a new XfrmUsersaFlush structure
	msg := &XfrmUsersaFlush{}

	// Create a reader for the byte slice
	reader := bytes.NewReader(b)

	// Deserialize Proto (uint8)
	if err := binary.Read(reader, binary.BigEndian, &msg.Proto); err != nil {
		return nil, err
	}

	// Return the deserialized XfrmUsersaFlush structure
	return msg, nil
}
