package core

import (
	"bytes"
	"encoding/binary"
)

// Serialize safely serializes the XfrmUsersaId structure into a byte slice.
// It returns the byte slice that represents the serialized structure, or an error if serialization fails.
func (msg *XfrmUsersaId) Serialize() ([]byte, error) {
	// Create a buffer to hold the serialized data.
	buf := new(bytes.Buffer)

	// Serialize the fields in the order they appear in the structure.

	// Serialize the Daddr (XfrmAddress)
	data, err := msg.Daddr.Serialize()
	if err != nil {
		return nil, err
	}
	if _, err := buf.Write(data); err != nil {
		return nil, err
	}

	// Serialize the Spi (Security Parameter Index)
	if err := binary.Write(buf, binary.BigEndian, msg.Spi); err != nil {
		return nil, err
	}

	// Serialize the Family (uint16)
	if err := binary.Write(buf, binary.BigEndian, msg.Family); err != nil {
		return nil, err
	}

	// Serialize the Proto (uint8)
	if err := binary.Write(buf, binary.BigEndian, msg.Proto); err != nil {
		return nil, err
	}

	// Serialize the Pad (byte)
	if err := binary.Write(buf, binary.BigEndian, msg.Pad); err != nil {
		return nil, err
	}

	// Return the serialized byte slice.
	return buf.Bytes(), nil
}
