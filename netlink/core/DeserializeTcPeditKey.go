package core

import (
	"errors"
)

// DeserializeTcPedit safely deserializes a byte slice into a TcPeditSel structure and a slice of TcPeditKey.
// The function ensures that the byte slice is large enough to hold both the TcPeditSel structure and all keys.
//
// If the byte slice is too small or invalid, it returns nil slices and an error.
func DeserializeTcPedit(b []byte) (*TcPeditSel, []TcPeditKey, error) {
	// Check if the byte slice is large enough to contain TcPeditSel.
	if len(b) < SizeOfTcPeditSel {
		return nil, nil, errors.New("input too small for TcPeditSel")
	}

	// Deserialize the TcPeditSel structure.
	tcPeditSel, err := DeserializeTcPeditSel(b)
	if err != nil {
		return nil, nil, err
	}

	// Prepare to deserialize keys, ensure there is enough space for NKeys.
	keys := make([]TcPeditKey, 0, tcPeditSel.NKeys)
	next := SizeOfTcPeditSel

	// Loop over NKeys and deserialize each key.
	for i := uint8(0); i < tcPeditSel.NKeys; i++ {
		// Ensure there's enough data left in the byte slice for the next key.
		if len(b[next:]) < SizeOfTcPeditKey {
			return nil, nil, errors.New("not enough bytes for the key")
		}

		// Deserialize each key.
		key, err := DeserializeTcPeditKey(b[next:])
		if err != nil {
			return nil, nil, err
		}
		keys = append(keys, *key)
		next += SizeOfTcPeditKey
	}

	// Return the deserialized TcPeditSel and the list of keys.
	return tcPeditSel, keys, nil
}
