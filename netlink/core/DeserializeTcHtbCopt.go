package core

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// DeserializeTcHtbCopt safely deserializes a byte slice into a TcHtbCopt structure.
func DeserializeTcHtbCopt(b []byte) (*TcHtbCopt, error) {
	if len(b) < SizeOfTcHtbCopt {
		return nil, errors.New("byte slice too short to deserialize TcHtbCopt")
	}

	msg := &TcHtbCopt{}
	reader := bytes.NewReader(b)

	// Deserialize fields using binary.Read
	if err := binary.Read(reader, binary.LittleEndian, &msg.Rate); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.LittleEndian, &msg.Ceil); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.LittleEndian, &msg.Buffer); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.LittleEndian, &msg.Cbuffer); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.LittleEndian, &msg.Quantum); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.LittleEndian, &msg.Level); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.LittleEndian, &msg.Prio); err != nil {
		return nil, err
	}

	return msg, nil
}
