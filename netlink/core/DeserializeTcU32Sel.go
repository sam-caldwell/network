package core

import (
	"encoding/binary"
	"errors"
)

// DeserializeTcU32Sel converts a byte slice into a TcU32Sel object.
// It carefully parses the byte slice by extracting each field manually
// and handles keys in the Keys array sequentially.
//
// Parameters:
// - b []byte: The byte slice to deserialize.
//
// Returns:
// - *TcU32Sel: A pointer to the deserialized TcU32Sel object.
// - error: An error if deserialization fails or the byte slice is too short.
//
//	type TcU32Sel struct {
//	   Flags    uint8
//	   Offshift uint8
//	   Nkeys    uint8
//	   Pad      uint8
//	   Offmask  uint16
//	   Off      uint16
//	   Offoff   int16
//	   Hoff     int16
//	   Hmask    uint32
//	   Keys     []TcU32Key
//	}
//
//	type TcU32Key struct {
//	   Mask    uint32
//	   Val     uint32
//	   Off     int32
//	   OffMask int32
//	}
func DeserializeTcU32Sel(b []byte) (*TcU32Sel, error) {
	if len(b) < SizeOfTcU32Sel {
		return nil, errors.New("input too short")
	}

	resultObject := &TcU32Sel{}

	// Deserialize fixed-size fields
	resultObject.Flags = b[0]
	resultObject.Offshift = b[1]
	resultObject.Nkeys = b[2]
	resultObject.Pad = b[3]
	resultObject.Offmask = binary.BigEndian.Uint16(b[4:6])
	resultObject.Off = binary.LittleEndian.Uint16(b[6:8])
	resultObject.Offoff = int16(binary.LittleEndian.Uint16(b[8:10]))
	resultObject.Hoff = int16(binary.LittleEndian.Uint16(b[10:12]))
	resultObject.Hmask = binary.BigEndian.Uint32(b[12:16])

	// Start reading after the fixed-size part of TcU32Sel
	next := SizeOfTcU32Sel

	// Deserialize keys
	for i := uint8(0); i < resultObject.Nkeys; i++ {
		if len(b) < next+SizeOfTcU32Key {
			return nil, errors.New("too short for keys")
		}

		// Deserialize each key
		currentKey, err := DeserializeTcU32Key(b[next:])
		if err != nil {
			return nil, err
		}

		resultObject.Keys = append(resultObject.Keys, *currentKey)
		next += SizeOfTcU32Key
	}

	return resultObject, nil
}
