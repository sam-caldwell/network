package core

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// DeserializeXfrmLifetimeCur safely deserializes a byte slice into an XfrmLifetimeCur struct.
// The byte slice must be of the correct size to match the size of the XfrmLifetimeCur struct.
// This function ensures portability and avoids using unsafe pointers.
//
// Returns a pointer to XfrmLifetimeCur or an error if the byte slice size is invalid.
func DeserializeXfrmLifetimeCur(b []byte) (*XfrmLifetimeCur, error) {
	// Check if the byte slice is the correct size to avoid out-of-bounds access.
	expectedSize := 4 * 8 // Each field is a uint64, so 4 fields * 8 bytes per uint64 = 32 bytes total.
	if len(b) != expectedSize {
		return nil, fmt.Errorf("invalid byte slice size: expected %d, got %d", expectedSize, len(b))
	}

	// Create a buffer to read from the byte slice
	buf := bytes.NewReader(b)

	// Create a new XfrmLifetimeCur struct to hold the deserialized data
	var cur XfrmLifetimeCur

	// Deserialize each field in the correct order, using binary.Read
	if err := binary.Read(buf, binary.BigEndian, &cur.Bytes); err != nil {
		return nil, fmt.Errorf("failed to deserialize Bytes: %w", err)
	}
	if err := binary.Read(buf, binary.BigEndian, &cur.Packets); err != nil {
		return nil, fmt.Errorf("failed to deserialize Packets: %w", err)
	}
	if err := binary.Read(buf, binary.BigEndian, &cur.AddTime); err != nil {
		return nil, fmt.Errorf("failed to deserialize AddTime: %w", err)
	}
	if err := binary.Read(buf, binary.BigEndian, &cur.UseTime); err != nil {
		return nil, fmt.Errorf("failed to deserialize UseTime: %w", err)
	}

	// Return the deserialized struct
	return &cur, nil
}
