package core

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// DeserializeTcPeditSel safely deserializes a byte slice into a TcPeditSel structure.
// It checks the length of the byte slice to ensure that it is at least the size of TcPeditSel
// and uses the encoding/binary package to read the fields safely.
//
// Returns an error if the byte slice is too small or cannot be properly deserialized.
func DeserializeTcPeditSel(b []byte) (*TcPeditSel, error) {
	// Check if the byte slice is large enough to hold a TcPeditSel
	if len(b) < SizeOfTcPeditSel {
		return nil, errors.New("byte slice too small to deserialize TcPeditSel")
	}

	// Create a new TcPeditSel struct
	sel := TcPeditSel{}

	// Create a reader for the byte slice
	reader := bytes.NewReader(b)

	// Safely deserialize the TcGen fields
	if err := binary.Read(reader, binary.BigEndian, &sel.TcGen.Index); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.BigEndian, &sel.TcGen.Capab); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.BigEndian, &sel.TcGen.Action); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.BigEndian, &sel.TcGen.Refcnt); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.BigEndian, &sel.TcGen.Bindcnt); err != nil {
		return nil, err
	}

	// Safely deserialize the NKeys and Flags
	if err := binary.Read(reader, binary.BigEndian, &sel.NKeys); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.BigEndian, &sel.Flags); err != nil {
		return nil, err
	}

	// Return the deserialized TcPeditSel
	return &sel, nil
}
