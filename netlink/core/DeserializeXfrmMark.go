package core

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// DeserializeXfrmMark - safely deserialize a byte slice into an XfrmMark struct.
// The byte slice must be of the correct size to match the XfrmMark struct.
// This function ensures portability and avoids using unsafe pointers.
//
// Returns a pointer to XfrmMark or an error if the byte slice size is invalid.
func DeserializeXfrmMark(b []byte) (*XfrmMark, error) {
	// Calculate the expected size of the XfrmMark struct.
	// XfrmMark has two uint32 fields, so the expected size is 8 bytes.
	expectedSize := 8

	// Check if the byte slice is the correct size.
	if len(b) != expectedSize {
		return nil, fmt.Errorf("invalid byte slice size: expected %d, got %d", expectedSize, len(b))
	}

	// Create a buffer to read from the byte slice.
	buf := bytes.NewReader(b)

	// Create a new XfrmMark struct to hold the deserialized data.
	var mark XfrmMark

	// Deserialize the Value field (uint32).
	if err := binary.Read(buf, binary.BigEndian, &mark.Value); err != nil {
		return nil, fmt.Errorf("failed to deserialize Value: %w", err)
	}

	// Deserialize the Mask field (uint32).
	if err := binary.Read(buf, binary.BigEndian, &mark.Mask); err != nil {
		return nil, fmt.Errorf("failed to deserialize Mask: %w", err)
	}

	// Return the deserialized struct.
	return &mark, nil
}
