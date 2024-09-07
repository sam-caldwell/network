package core

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// DeserializeXfrmSelector safely deserializes a byte slice into an XfrmSelector structure.
// It checks the length of the byte slice to ensure it is at least the size of XfrmSelector.
//
// Returns an error if the byte slice is too small or deserialization fails.
func DeserializeXfrmSelector(b []byte) (*XfrmSelector, error) {
	// Check if the byte slice is large enough to hold an XfrmSelector.
	if len(b) < SizeOfXfrmSelector {
		return nil, errors.New("byte slice too small to deserialize XfrmSelector")
	}

	// Create a new XfrmSelector struct.
	sel := XfrmSelector{}

	if addr, err := DeserializeXfrmAddress(b[0:SizeofXfrmAddress]); err != nil {
		return nil, err
	} else {
		sel.Daddr = *addr
	}

	if addr, err := DeserializeXfrmAddress(b[SizeofXfrmAddress : 2*SizeofXfrmAddress]); err != nil {
		return nil, err
	} else {
		sel.Saddr = *addr
	}

	// Create a reader for the byte slice.
	reader := bytes.NewReader(b[2*SizeofXfrmAddress:])

	if err := binary.Read(reader, binary.BigEndian, &sel.Dport); err != nil {
		return nil, err
	}

	if err := binary.Read(reader, binary.BigEndian, &sel.DportMask); err != nil {
		return nil, err
	}

	if err := binary.Read(reader, binary.BigEndian, &sel.Sport); err != nil {
		return nil, err
	}

	if err := binary.Read(reader, binary.BigEndian, &sel.SportMask); err != nil {
		return nil, err
	}

	if err := binary.Read(reader, binary.BigEndian, &sel.Family); err != nil {
		return nil, err
	}

	if err := binary.Read(reader, binary.BigEndian, &sel.PrefixlenD); err != nil {
		return nil, err
	}

	if err := binary.Read(reader, binary.BigEndian, &sel.PrefixlenS); err != nil {
		return nil, err
	}

	if err := binary.Read(reader, binary.BigEndian, &sel.Proto); err != nil {
		return nil, err
	}

	if err := binary.Read(reader, binary.BigEndian, &sel.Pad); err != nil {
		return nil, err
	}

	if err := binary.Read(reader, binary.BigEndian, &sel.Ifindex); err != nil {
		return nil, err
	}

	if err := binary.Read(reader, binary.BigEndian, &sel.User); err != nil {
		return nil, err
	}

	// Return the deserialized XfrmSelector.
	return &sel, nil
}
