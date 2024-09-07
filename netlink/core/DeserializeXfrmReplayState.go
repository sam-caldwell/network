package core

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// DeserializeXfrmReplayState safely deserializes a byte slice into an XfrmReplayState structure.
// It returns the deserialized structure or an error if deserialization fails.
func DeserializeXfrmReplayState(b []byte) (*XfrmReplayState, error) {
	// Check if the byte slice is large enough to hold the XfrmReplayState structure.
	if len(b) < SizeOfXfrmReplayState {
		return nil, errors.New("byte slice too small to deserialize XfrmReplayState")
	}

	// Create a new XfrmReplayState structure
	msg := &XfrmReplayState{}

	// Create a reader for the byte slice
	reader := bytes.NewReader(b)

	// Deserialize OSeq (uint32)
	if err := binary.Read(reader, binary.BigEndian, &msg.OSeq); err != nil {
		return nil, err
	}

	// Deserialize Seq (uint32)
	if err := binary.Read(reader, binary.BigEndian, &msg.Seq); err != nil {
		return nil, err
	}

	// Deserialize BitMap (uint32)
	if err := binary.Read(reader, binary.BigEndian, &msg.BitMap); err != nil {
		return nil, err
	}

	// Return the deserialized XfrmReplayState structure
	return msg, nil
}
