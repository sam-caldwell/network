package core

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// DeserializeXfrmId - safely deserialize a byte slice into an XfrmId struct.
// The byte slice must be of the correct length to match the size of the XfrmId struct.
// This function ensures portability and avoids using unsafe pointers.
//
// Returns a pointer to XfrmId or an error if the byte slice size is invalid.
func DeserializeXfrmId(b []byte) (*XfrmId, error) {
	// Calculate the expected size of the XfrmId struct.
	// Daddr size (assumed 16 bytes for IPv6), Spi (4 bytes), Proto (1 byte), Pad (3 bytes).
	expectedSize := 16 + 4 + 1 + 3

	// Check if the byte slice is the correct size.
	if len(b) != expectedSize {
		return nil, fmt.Errorf("invalid byte slice size: expected %d, got %d", expectedSize, len(b))
	}

	// Create a buffer to read from the byte slice.
	buf := bytes.NewReader(b)

	// Create a new XfrmId struct to hold the deserialized data.
	var id XfrmId

	// Deserialize the Daddr (assuming XfrmAddress is a byte array or struct of 16 bytes).
	// Adjust the size of Daddr if necessary (IPv4 vs IPv6).
	if err := binary.Read(buf, binary.BigEndian, &id.Daddr); err != nil {
		return nil, fmt.Errorf("failed to deserialize Daddr: %w", err)
	}

	// Deserialize the Spi (Security Parameter Index).
	if err := binary.Read(buf, binary.BigEndian, &id.Spi); err != nil {
		return nil, fmt.Errorf("failed to deserialize Spi: %w", err)
	}

	// Deserialize the Proto (protocol identifier).
	if err := binary.Read(buf, binary.BigEndian, &id.Proto); err != nil {
		return nil, fmt.Errorf("failed to deserialize Proto: %w", err)
	}

	// Deserialize the Pad (3-byte padding).
	if err := binary.Read(buf, binary.BigEndian, &id.Pad); err != nil {
		return nil, fmt.Errorf("failed to deserialize Pad: %w", err)
	}

	// Return the deserialized struct.
	return &id, nil
}
