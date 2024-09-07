package core

import (
	"bytes"
	"encoding/binary"
	"errors"
	"unsafe"
)

// DeserializeXfrmEncapTmpl safely deserializes a byte slice into an XfrmEncapTmpl structure.
// It returns the deserialized structure or an error if deserialization fails.
func DeserializeXfrmEncapTmpl(b []byte) (*XfrmEncapTmpl, error) {
	if len(b) < SizeOfXfrmEncapTmpl {
		return nil, errors.New("byte slice too small to deserialize XfrmEncapTmpl")
	}

	// Create a new XfrmEncapTmpl structure
	msg := &XfrmEncapTmpl{}

	// Create a reader for the byte slice
	reader := bytes.NewReader(b)

	// Deserialize EncapType (2 bytes, little-endian)
	if err := binary.Read(reader, binary.LittleEndian, &msg.EncapType); err != nil {
		return nil, err
	}

	// Deserialize EncapSport (2 bytes, big-endian)
	if err := binary.Read(reader, binary.BigEndian, &msg.EncapSport); err != nil {
		return nil, err
	}

	// Deserialize EncapDport (2 bytes, big-endian)
	if err := binary.Read(reader, binary.BigEndian, &msg.EncapDport); err != nil {
		return nil, err
	}

	// Deserialize Pad (2 bytes)
	if _, err := reader.Read(msg.Pad[:]); err != nil {
		return nil, err
	}

	// Deserialize EncapOa (XfrmAddress) using its own deserialization method
	if temp, err := DeserializeXfrmAddress(b[unsafe.Sizeof(XfrmEncapTmpl{}):]); err != nil {
		return nil, err
	} else {
		msg.EncapOa = *temp
	}

	// Return the deserialized XfrmEncapTmpl structure
	return msg, nil
}
