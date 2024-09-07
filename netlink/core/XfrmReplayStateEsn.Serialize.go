package core

import (
	"bytes"
	"encoding/binary"
)

// Serialize safely serializes the XfrmReplayStateEsn structure into a byte slice.
// The Bmp field is not serialized, as it is set by the kernel.
func (msg *XfrmReplayStateEsn) Serialize() ([]byte, error) {
	// Create a buffer to hold the serialized data.
	buf := new(bytes.Buffer)

	// Serialize BmpLen (uint32)
	if err := binary.Write(buf, binary.BigEndian, msg.BmpLen); err != nil {
		return nil, err
	}

	// Serialize OSeq (uint32)
	if err := binary.Write(buf, binary.BigEndian, msg.OSeq); err != nil {
		return nil, err
	}

	// Serialize Seq (uint32)
	if err := binary.Write(buf, binary.BigEndian, msg.Seq); err != nil {
		return nil, err
	}

	// Serialize OSeqHi (uint32)
	if err := binary.Write(buf, binary.BigEndian, msg.OSeqHi); err != nil {
		return nil, err
	}

	// Serialize SeqHi (uint32)
	if err := binary.Write(buf, binary.BigEndian, msg.SeqHi); err != nil {
		return nil, err
	}

	// Serialize ReplayWindow (uint32)
	if err := binary.Write(buf, binary.BigEndian, msg.ReplayWindow); err != nil {
		return nil, err
	}

	// Note: We skip serializing Bmp, as per the requirements

	// Return the serialized byte slice
	return buf.Bytes(), nil
}
