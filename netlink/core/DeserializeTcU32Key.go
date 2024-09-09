package core

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// DeserializeTcU32Key - Deserializes a byte slice into a TcU32Key struct.
//
//	type TcU32Key struct {
//	   Mask    uint32
//	   Val     uint32
//	   Off     int32
//	   OffMask int32
//	}
func DeserializeTcU32Key(b []byte) (*TcU32Key, error) {
	if len(b) < SizeOfTcU32Key {
		return nil, fmt.Errorf("byte slice too short: expected %d, got %d", SizeOfTcU32Key, len(b))
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
