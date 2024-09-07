package core

import (
	"bytes"
	"encoding/binary"
)

// Serialize safely serializes the XfrmUserTmpl structure into a byte slice.
// It returns the byte slice that represents the serialized structure, or an error if serialization fails.
func (msg *XfrmUserTmpl) Serialize() ([]byte, error) {
	// Create a buffer to hold the serialized data.
	buf := new(bytes.Buffer)

	// Serialize the fields in the order they appear in the structure.

	// Serialize the XfrmId
	if temp, err := msg.XfrmId.Serialize(); err != nil {
		return nil, err
	} else {
		if err := binary.Write(buf, binary.BigEndian, temp); err != nil {
			return nil, err
		}
	}

	// Serialize the Family
	if err := binary.Write(buf, binary.BigEndian, msg.Family); err != nil {
		return nil, err
	}

	// Serialize Pad1
	if err := binary.Write(buf, binary.BigEndian, msg.Pad1); err != nil {
		return nil, err
	}

	// Serialize Saddr (XfrmAddress)
	sAddr := msg.Saddr.Serialize()
	if err := binary.Write(buf, binary.BigEndian, &sAddr); err != nil {
		return nil, err
	}

	// Serialize Reqid
	if err := binary.Write(buf, binary.BigEndian, msg.Reqid); err != nil {
		return nil, err
	}

	// Serialize Mode
	if err := binary.Write(buf, binary.BigEndian, msg.Mode); err != nil {
		return nil, err
	}

	// Serialize Share
	if err := binary.Write(buf, binary.BigEndian, msg.Share); err != nil {
		return nil, err
	}

	// Serialize Optional
	if err := binary.Write(buf, binary.BigEndian, msg.Optional); err != nil {
		return nil, err
	}

	// Serialize Pad2
	if err := binary.Write(buf, binary.BigEndian, msg.Pad2); err != nil {
		return nil, err
	}

	// Serialize Aalgos
	if err := binary.Write(buf, binary.BigEndian, msg.Aalgos); err != nil {
		return nil, err
	}

	// Serialize Ealgos
	if err := binary.Write(buf, binary.BigEndian, msg.Ealgos); err != nil {
		return nil, err
	}

	// Serialize Calgos
	if err := binary.Write(buf, binary.BigEndian, msg.Calgos); err != nil {
		return nil, err
	}

	// Return the serialized byte slice.
	return buf.Bytes(), nil
}
