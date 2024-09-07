package core

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// DeserializeXfrmUsersaId safely deserializes a byte slice into an XfrmUserSaId structure.
// It uses DeserializeXfrmAddress for the Daddr field and manually deserializes the other fields.
func DeserializeXfrmUsersaId(b []byte) (*XfrmUserSaId, error) {
	if len(b) < SizeOfXfrmUsersaId {
		return nil, errors.New("byte slice too small to deserialize XfrmUserSaId")
	}

	// Create a new XfrmUserSaId struct
	id := &XfrmUserSaId{}

	// Create a reader for the byte slice
	reader := bytes.NewReader(b)

	// Deserialize the Daddr (XfrmAddress) using DeserializeXfrmAddress
	daddr, err := DeserializeXfrmAddress(b)
	if err != nil {
		return nil, err
	}
	id.Daddr = *daddr
	offset := SizeOfXfrmAddress // Adjust the offset after deserializing the XfrmAddress

	// Seek the reader to the next field (after Daddr)
	if _, err = reader.Seek(int64(offset), 0); err != nil {
		return nil, err
	}

	// Deserialize Spi (Security Parameter Index)
	if err := binary.Read(reader, binary.BigEndian, &id.Spi); err != nil {
		return nil, err
	}

	// Deserialize Family (uint16)
	if err := binary.Read(reader, binary.BigEndian, &id.Family); err != nil {
		return nil, err
	}

	// Deserialize Proto (uint8)
	if err := binary.Read(reader, binary.BigEndian, &id.Proto); err != nil {
		return nil, err
	}

	// Deserialize Pad (byte)
	if err := binary.Read(reader, binary.BigEndian, &id.Pad); err != nil {
		return nil, err
	}

	// Return the deserialized XfrmUserSaId structure
	return id, nil
}
