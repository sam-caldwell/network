package core

import (
	"encoding/binary"
	"errors"
)

// DeserializeTcU32Sel converts a byte slice into a TcU32Sel object.
// It carefully parses the byte slice by extracting each field manually
// and handles keys in the Keys array sequentially.
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
	resultObject.Off = NativeEndian.Uint16(b[6:8])
	resultObject.Offoff = int16(NativeEndian.Uint16(b[8:10]))
	resultObject.Hoff = int16(NativeEndian.Uint16(b[10:12]))
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
