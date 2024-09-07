package core

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// DeserializeXfrmUserSpiInfo safely deserializes a byte slice into an XfrmUserSpiInfo structure.
// It uses appropriate deserialization functions for the embedded XfrmUserSaInfo.
func DeserializeXfrmUserSpiInfo(b []byte) (*XfrmUserSpiInfo, error) {
	if len(b) < SizeOfXfrmUserSpiInfo {
		return nil, errors.New("byte slice too small to deserialize XfrmUserSpiInfo")
	}

	// Create a reader for the byte slice
	reader := bytes.NewReader(b)

	// Create a new XfrmUserSpiInfo struct
	msg := &XfrmUserSpiInfo{}

	// Deserialize the XfrmUserSaInfo
	usersaInfo, err := DeserializeXfrmUserSaInfo(b)
	if err != nil {
		return nil, err
	}
	msg.XfrmUsersaInfo = *usersaInfo
	offset := SizeOfXfrmUserSaInfo

	// Move the reader past the deserialized XfrmUserSaInfo
	if _, err := reader.Seek(int64(offset), 0); err != nil {
		return nil, err
	}

	// Deserialize Min (uint32)
	if err := binary.Read(reader, binary.BigEndian, &msg.Min); err != nil {
		return nil, err
	}

	// Deserialize Max (uint32)
	if err := binary.Read(reader, binary.BigEndian, &msg.Max); err != nil {
		return nil, err
	}

	// Return the deserialized XfrmUserSpiInfo structure
	return msg, nil
}
