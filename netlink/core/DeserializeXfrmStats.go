package core

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// DeserializeXfrmStats safely deserializes a byte slice into an XfrmStats structure.
// Returns an error if the byte slice is too small or if deserialization fails.
func DeserializeXfrmStats(b []byte) (*XfrmStats, error) {
	// Check if the byte slice is large enough to hold an XfrmStats structure.
	if len(b) < SizeOfXfrmStats {
		return nil, errors.New("byte slice too small to deserialize XfrmStats")
	}

	// Create a reader for the byte slice.
	reader := bytes.NewReader(b)

	// Create a new XfrmStats struct to store the deserialized data.
	stats := &XfrmStats{}

	// Deserialize ReplayWindow (uint32)
	if err := binary.Read(reader, binary.BigEndian, &stats.ReplayWindow); err != nil {
		return nil, err
	}

	// Deserialize Replay (uint32)
	if err := binary.Read(reader, binary.BigEndian, &stats.Replay); err != nil {
		return nil, err
	}

	// Deserialize IntegrityFailed (uint32)
	if err := binary.Read(reader, binary.BigEndian, &stats.IntegrityFailed); err != nil {
		return nil, err
	}

	// Return the deserialized XfrmStats structure.
	return stats, nil
}
