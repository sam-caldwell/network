package core

import (
	"bytes"
	"encoding/binary"
)

// Serialize safely serializes the XfrmAlgoAuth structure into a byte slice.
// It returns the byte slice or an error if serialization fails.
func (msg *XfrmAlgoAuth) Serialize() ([]byte, error) {
	// Create a buffer to hold the serialized data
	buf := new(bytes.Buffer)

	// Serialize AlgName (64 bytes)
	if _, err := buf.Write(msg.AlgName[:]); err != nil {
		return nil, err
	}

	// Serialize AlgKeyLen (4 bytes)
	if err := binary.Write(buf, binary.BigEndian, msg.AlgKeyLen); err != nil {
		return nil, err
	}

	// Serialize AlgTruncLen (4 bytes)
	if err := binary.Write(buf, binary.BigEndian, msg.AlgTruncLen); err != nil {
		return nil, err
	}

	// Serialize AlgKey (variable length)
	if _, err := buf.Write(msg.AlgKey); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
