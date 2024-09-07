package core

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// DeserializeXfrmUserTmpl safely deserializes a byte slice into an XfrmUserTmpl structure.
// It uses DeserializeXfrmId and DeserializeXfrmAddress for respective fields.
// Returns an error if the byte slice is too small or deserialization fails.
func DeserializeXfrmUserTmpl(b []byte) (*XfrmUserTmpl, error) {
	if len(b) < SizeOfXfrmUserTmpl {
		return nil, errors.New("byte slice too small to deserialize XfrmUserTmpl")
	}

	// Create a reader for the byte slice.
	reader := bytes.NewReader(b)

	// Create a new XfrmUserTmpl struct.
	tmpl := &XfrmUserTmpl{}

	// Deserialize the XfrmId using DeserializeXfrmId
	xfrmId, err := DeserializeXfrmId(b)
	if err != nil {
		return nil, err
	}
	tmpl.XfrmId = *xfrmId
	offset := SizeOfXfrmId

	// Move the reader to the next field (Family)
	reader.Seek(int64(offset), 0)

	// Deserialize Family
	if err := binary.Read(reader, binary.BigEndian, &tmpl.Family); err != nil {
		return nil, err
	}
	offset += 2 // Size of Family (uint16)

	// Deserialize Pad1
	if err := binary.Read(reader, binary.BigEndian, &tmpl.Pad1); err != nil {
		return nil, err
	}
	offset += 2 // Size of Pad1 (2 bytes)

	// Deserialize the Saddr using DeserializeXfrmAddress
	saddr, err := DeserializeXfrmAddress(b[offset:])
	if err != nil {
		return nil, err
	}
	tmpl.Saddr = *saddr
	offset += SizeOfXfrmAddress

	// Move the reader to the next field (Reqid)
	reader.Seek(int64(offset), 0)

	// Deserialize Reqid
	if err := binary.Read(reader, binary.BigEndian, &tmpl.Reqid); err != nil {
		return nil, err
	}
	offset += 4 // Size of Reqid (uint32)

	// Deserialize Mode
	if err := binary.Read(reader, binary.BigEndian, &tmpl.Mode); err != nil {
		return nil, err
	}

	// Deserialize Share
	if err := binary.Read(reader, binary.BigEndian, &tmpl.Share); err != nil {
		return nil, err
	}

	// Deserialize Optional
	if err := binary.Read(reader, binary.BigEndian, &tmpl.Optional); err != nil {
		return nil, err
	}

	// Deserialize Pad2
	if err := binary.Read(reader, binary.BigEndian, &tmpl.Pad2); err != nil {
		return nil, err
	}

	// Deserialize Aalgos
	if err := binary.Read(reader, binary.BigEndian, &tmpl.Aalgos); err != nil {
		return nil, err
	}

	// Deserialize Ealgos
	if err := binary.Read(reader, binary.BigEndian, &tmpl.Ealgos); err != nil {
		return nil, err
	}

	// Deserialize Calgos
	if err := binary.Read(reader, binary.BigEndian, &tmpl.Calgos); err != nil {
		return nil, err
	}

	// Return the deserialized XfrmUserTmpl
	return tmpl, nil
}
