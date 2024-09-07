package core

import (
	"bytes"
	"encoding/binary"
)

// Serialize safely serializes the XfrmStats structure into a byte slice.
// It returns the byte slice that represents the serialized structure, or an error if serialization fails.
func (msg *XfrmStats) Serialize() ([]byte, error) {
	// Create a buffer to hold the serialized data.
	buf := new(bytes.Buffer)

	// Serialize the ReplayWindow (uint32)
	if err := binary.Write(buf, binary.BigEndian, msg.ReplayWindow); err != nil {
		return nil, err
	}

	// Serialize the Replay (uint32)
	if err := binary.Write(buf, binary.BigEndian, msg.Replay); err != nil {
		return nil, err
	}

	// Serialize the IntegrityFailed (uint32)
	if err := binary.Write(buf, binary.BigEndian, msg.IntegrityFailed); err != nil {
		return nil, err
	}

	// Return the serialized byte slice.
	return buf.Bytes(), nil
}
