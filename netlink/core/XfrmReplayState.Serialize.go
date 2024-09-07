package core

import (
	"bytes"
	"encoding/binary"
)

// Serialize safely serializes the XfrmReplayState structure into a byte slice.
// It returns the serialized byte slice or an error if serialization fails.
func (msg *XfrmReplayState) Serialize() ([]byte, error) {
	// Create a buffer to hold the serialized data.
	buf := new(bytes.Buffer)

	// Serialize OSeq (uint32)
	if err := binary.Write(buf, binary.BigEndian, msg.OSeq); err != nil {
		return nil, err
	}

	// Serialize Seq (uint32)
	if err := binary.Write(buf, binary.BigEndian, msg.Seq); err != nil {
		return nil, err
	}

	// Serialize BitMap (uint32)
	if err := binary.Write(buf, binary.BigEndian, msg.BitMap); err != nil {
		return nil, err
	}

	// Return the serialized byte slice
	return buf.Bytes(), nil
}
