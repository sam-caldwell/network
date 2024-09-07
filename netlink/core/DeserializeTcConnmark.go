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
	if err := binary.Read(buf, NativeEndian, &result.Index); err != nil {
		return nil, err
	}
	if err := binary.Read(buf, NativeEndian, &result.Capab); err != nil {
		return nil, err
	}
	if err := binary.Read(buf, NativeEndian, &result.Action); err != nil {
		return nil, err
	}
	if err := binary.Read(buf, NativeEndian, &result.Refcnt); err != nil {
		return nil, err
	}
	if err := binary.Read(buf, NativeEndian, &result.Bindcnt); err != nil {
		return nil, err
	}

	// Deserialize TcConnmark-specific part
	if err := binary.Read(buf, NativeEndian, &result.Zone); err != nil {
		return nil, err
	}

	return result, nil
}
