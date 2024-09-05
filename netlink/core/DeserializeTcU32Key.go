package core

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// DeserializeTcU32Key - Deserializes a byte slice into a TcU32Key struct.
func DeserializeTcU32Key(b []byte) (*TcU32Key, error) {
	if len(b) < SizeofTcU32Key {
		return nil, fmt.Errorf("byte slice too short: expected %d, got %d", SizeofTcU32Key, len(b))
	}

	buf := bytes.NewReader(b)
	var key TcU32Key

	// Deserialize each field with proper endianness
	if err := binary.Read(buf, binary.BigEndian, &key.Mask); err != nil {
		return nil, err
	}
	if err := binary.Read(buf, binary.BigEndian, &key.Val); err != nil {
		return nil, err
	}
	if err := binary.Read(buf, binary.LittleEndian, &key.Off); err != nil {
		return nil, err
	}
	if err := binary.Read(buf, binary.LittleEndian, &key.OffMask); err != nil {
		return nil, err
	}

	return &key, nil
}
