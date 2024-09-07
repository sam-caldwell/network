package core

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// DeserializeXfrmUserPolicyId safely deserializes a byte slice into a XfrmUserPolicyId struct.
// It checks the length of the byte slice to ensure it is at least the size of XfrmUserPolicyId.
func DeserializeXfrmUserPolicyId(b []byte) (*XfrmUserPolicyId, error) {

	// Check if the byte slice is large enough to hold a XfrmUserPolicyId.
	if len(b) < SizeofXfrmUserPolicyId {
		return nil, errors.New("byte slice too small to deserialize XfrmUserPolicyId")
	}

	// Create a new XfrmUserPolicyId struct.
	id := XfrmUserPolicyId{}

	if sel, err := DeserializeXfrmSelector(b); err != nil {
		return nil, err
	} else {
		id.Sel = *sel
	}

	// Create a reader for the byte slice.
	reader := bytes.NewReader(b)

	// Safely deserialize the fields using binary.Read.

	// Deserialize XfrmSelector
	if err := binary.Read(reader, binary.BigEndian, &id.Sel); err != nil {
		return nil, err
	}

	// Deserialize the Index
	if err := binary.Read(reader, binary.BigEndian, &id.Index); err != nil {
		return nil, err
	}

	// Deserialize the Dir
	if err := binary.Read(reader, binary.BigEndian, &id.Dir); err != nil {
		return nil, err
	}

	// Deserialize the Pad
	if err := binary.Read(reader, binary.BigEndian, &id.Pad); err != nil {
		return nil, err
	}

	// Return the deserialized XfrmUserPolicyId.
	return &id, nil
}
