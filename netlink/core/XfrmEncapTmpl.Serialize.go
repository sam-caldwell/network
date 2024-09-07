package core

import (
	"bytes"
	"encoding/binary"
)

// Serialize safely serializes the XfrmEncapTmpl structure into a byte slice.
// It returns the byte slice or an error if serialization fails.
func (msg *XfrmEncapTmpl) Serialize() ([]byte, error) {
	// Create a buffer to hold the serialized data.
	buf := new(bytes.Buffer)

	// Serialize EncapType (2 bytes, little endian by default)
	if err := binary.Write(buf, binary.LittleEndian, msg.EncapType); err != nil {
		return nil, err
	}

	// Serialize EncapSport (2 bytes, big endian)
	if err := binary.Write(buf, binary.BigEndian, msg.EncapSport); err != nil {
		return nil, err
	}

	// Serialize EncapDport (2 bytes, big endian)
	if err := binary.Write(buf, binary.BigEndian, msg.EncapDport); err != nil {
		return nil, err
	}

	// Serialize Pad (2 bytes)
	if _, err := buf.Write(msg.Pad[:]); err != nil {
		return nil, err
	}

	// Serialize EncapOa (XfrmAddress, variable size)
	if _, err := buf.Write(msg.EncapOa.Serialize()); err != nil {
		return nil, err
	}

	// Return the serialized byte slice.
	return buf.Bytes(), nil
}
