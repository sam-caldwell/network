package core

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// DeserializeXfrmLifetimeCfg safely deserializes a byte slice into an XfrmLifetimeCfg struct.
// The byte slice must be of the correct length to match the size of XfrmLifetimeCfg.
// This function ensures portability and avoids using unsafe pointers.
//
// Returns a pointer to XfrmLifetimeCfg or an error if the byte slice size is invalid.
func DeserializeXfrmLifetimeCfg(b []byte) (*XfrmLifetimeCfg, error) {
	// Check if the byte slice is the correct size to avoid out-of-bounds access.
	expectedSize := 8 * 8 // Each field is uint64, and there are 8 fields, so 64 bytes total.
	if len(b) != expectedSize {
		return nil, fmt.Errorf("invalid byte slice size: expected %d, got %d", expectedSize, len(b))
	}

	// Create a buffer to read from the byte slice
	buf := bytes.NewReader(b)

	// Create a new XfrmLifetimeCfg struct to hold the deserialized data
	var cfg XfrmLifetimeCfg

	// Deserialize each field in the correct order, using binary.Read
	if err := binary.Read(buf, binary.BigEndian, &cfg.SoftByteLimit); err != nil {
		return nil, fmt.Errorf("failed to deserialize SoftByteLimit: %w", err)
	}
	if err := binary.Read(buf, binary.BigEndian, &cfg.HardByteLimit); err != nil {
		return nil, fmt.Errorf("failed to deserialize HardByteLimit: %w", err)
	}
	if err := binary.Read(buf, binary.BigEndian, &cfg.SoftPacketLimit); err != nil {
		return nil, fmt.Errorf("failed to deserialize SoftPacketLimit: %w", err)
	}
	if err := binary.Read(buf, binary.BigEndian, &cfg.HardPacketLimit); err != nil {
		return nil, fmt.Errorf("failed to deserialize HardPacketLimit: %w", err)
	}
	if err := binary.Read(buf, binary.BigEndian, &cfg.SoftAddExpiresSeconds); err != nil {
		return nil, fmt.Errorf("failed to deserialize SoftAddExpiresSeconds: %w", err)
	}
	if err := binary.Read(buf, binary.BigEndian, &cfg.HardAddExpiresSeconds); err != nil {
		return nil, fmt.Errorf("failed to deserialize HardAddExpiresSeconds: %w", err)
	}
	if err := binary.Read(buf, binary.BigEndian, &cfg.SoftUseExpiresSeconds); err != nil {
		return nil, fmt.Errorf("failed to deserialize SoftUseExpiresSeconds: %w", err)
	}
	if err := binary.Read(buf, binary.BigEndian, &cfg.HardUseExpiresSeconds); err != nil {
		return nil, fmt.Errorf("failed to deserialize HardUseExpiresSeconds: %w", err)
	}

	// Return the deserialized struct
	return &cfg, nil
}
