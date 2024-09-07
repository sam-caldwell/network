package core

import (
	"bytes"
	"encoding/binary"
)

// Serialize - safely serializes the XfrmUserPolicyId structure into a byte slice. The function returns the byte
// slice that represents the serialized structure, or an error if serialization fails.
func (msg *XfrmUserPolicyId) Serialize() ([]byte, error) {
	// Create a buffer to hold the serialized data.
	buf := new(bytes.Buffer)

	// Serialize the fields in the order they appear in the structure.

	// Serialize the Sel (XfrmSelector)
	if err := binary.Write(buf, binary.BigEndian, &msg.Sel); err != nil {
		return nil, err
	}

	// Serialize the Index
	if err := binary.Write(buf, binary.BigEndian, msg.Index); err != nil {
		return nil, err
	}

	// Serialize the Dir
	if err := binary.Write(buf, binary.BigEndian, msg.Dir); err != nil {
		return nil, err
	}

	// Serialize the Pad (padding)
	if err := binary.Write(buf, binary.BigEndian, &msg.Pad); err != nil {
		return nil, err
	}

	// Return the serialized byte slice.
	return buf.Bytes(), nil
}
