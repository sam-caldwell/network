package core

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// DeserializeTcConnmark - Safely deserialize a byte slice into a TcConnmark object.
// This method reads each field manually and ensures proper byte order and alignment.
func DeserializeTcConnmark(b []byte) (*TcConnmark, error) {
	if len(b) < SizeOfTcConnmark {
		return nil, errors.New("byte slice too short to deserialize TcConnmark")
	}

	buf := bytes.NewReader(b)
	result := &TcConnmark{}

	// Deserialize TcGen part
	if err := binary.Read(buf, binary.LittleEndian, &result.Index); err != nil {
		return nil, err
	}
	if err := binary.Read(buf, binary.LittleEndian, &result.Capab); err != nil {
		return nil, err
	}
	if err := binary.Read(buf, binary.LittleEndian, &result.Action); err != nil {
		return nil, err
	}
	if err := binary.Read(buf, binary.LittleEndian, &result.Refcnt); err != nil {
		return nil, err
	}
	if err := binary.Read(buf, binary.LittleEndian, &result.Bindcnt); err != nil {
		return nil, err
	}

	// Deserialize TcConnmark-specific part
	if err := binary.Read(buf, binary.LittleEndian, &result.Zone); err != nil {
		return nil, err
	}

	return result, nil
}
