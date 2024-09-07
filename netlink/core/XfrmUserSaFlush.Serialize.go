package core

import (
	"bytes"
	"encoding/binary"
)

// Serialize safely serializes the XfrmUserSaFlush structure into a byte slice.
// It returns the serialized byte slice or an error if serialization fails.
func (msg *XfrmUserSaFlush) Serialize() ([]byte, error) {
	// Create a buffer to hold the serialized data.
	buf := new(bytes.Buffer)

	// Serialize the Proto field (uint8)
	if err := binary.Write(buf, binary.BigEndian, msg.Proto); err != nil {
		return nil, err
	}

	// Return the serialized byte slice.
	return buf.Bytes(), nil
}
