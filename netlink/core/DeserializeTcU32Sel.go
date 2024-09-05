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
func DeserializeTcU32Sel(b []byte) (*TcU32Sel, error) {
	if len(b) < SizeofTcU32Sel {
		return nil, errors.New("DeserializeTcU32Sel: byte slice too short")
	}

	resultObject := &TcU32Sel{}

	// Deserialize fixed-size fields
	resultObject.Flags = b[0]
	resultObject.Offshift = b[1]
	resultObject.Nkeys = b[2]
	resultObject.Pad = b[3]
	resultObject.Offmask = binary.BigEndian.Uint16(b[4:])
	resultObject.Off = binary.LittleEndian.Uint16(b[6:])
	resultObject.Offoff = int16(binary.LittleEndian.Uint16(b[8:]))
	resultObject.Hoff = int16(binary.LittleEndian.Uint16(b[10:]))
	resultObject.Hmask = binary.BigEndian.Uint32(b[12:])

	// Start reading after the fixed-size part of TcU32Sel
	next := SizeofTcU32Sel

	// Deserialize keys
	for i := uint8(0); i < resultObject.Nkeys; i++ {
		if len(b) < next+SizeofTcU32Key {
			return nil, errors.New("DeserializeTcU32Sel: byte slice too short for keys")
		}

		// Deserialize each key
		currentKey, err := DeserializeTcU32Key(b[next:])
		if err != nil {
			return nil, err
		}

		resultObject.Keys = append(resultObject.Keys, *currentKey)
		next += SizeofTcU32Key
	}

	return resultObject, nil
}
