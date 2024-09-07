package core

import (
	"bytes"
	"encoding/binary"
)

// Serialize safely serializes the XfrmUserSpiInfo structure into a byte slice.
// It returns the byte slice that represents the serialized structure, or an error if serialization fails.
func (msg *XfrmUserSpiInfo) Serialize() ([]byte, error) {
	// Create a buffer to hold the serialized data.
	buf := new(bytes.Buffer)

	// Serialize the embedded XfrmUsersaInfo
	usersaBytes, err := msg.XfrmUsersaInfo.Serialize()
	if err != nil {
		return nil, err
	}
	if _, err := buf.Write(usersaBytes); err != nil {
		return nil, err
	}

	// Serialize the Min (uint32)
	if err := binary.Write(buf, binary.BigEndian, msg.Min); err != nil {
		return nil, err
	}

	// Serialize the Max (uint32)
	if err := binary.Write(buf, binary.BigEndian, msg.Max); err != nil {
		return nil, err
	}

	// Return the serialized byte slice.
	return buf.Bytes(), nil
}
