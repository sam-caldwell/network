package core

import "unsafe"

// DeserializeTcPedit safely deserializes a byte slice into a TcPeditSel structure and a slice of TcPeditKey.
// The function ensures that the byte slice is large enough to hold both the TcPeditSel structure and all keys.
//
// If the byte slice is too small or invalid, it returns nil for both TcPeditSel and TcPeditKey slices.
func DeserializeTcPedit(b []byte) (*TcPeditSel, []TcPeditKey) {
	// Check if the byte slice is large enough to contain TcPeditSel.
	if len(b) < SizeOfTcPeditSel {
		return nil, nil // Return nil if the byte slice is too small
	}

	// Deserialize TcPeditSel
	x := &TcPeditSel{}
	copy((*(*[SizeOfTcPeditSel]byte)(unsafe.Pointer(x)))[:SizeOfTcPeditSel], b)

	// Prepare to deserialize keys, check if there are enough bytes for the keys
	keys := make([]TcPeditKey, 0, x.NKeys)
	next := SizeOfTcPeditSel

	for i := uint8(0); i < x.NKeys; i++ {
		// Ensure there's enough data left in the byte slice for the next key
		if len(b[next:]) < SizeOfTcPeditKey {
			return nil, nil // Return nil if the byte slice doesn't contain enough bytes for the key
		}

		// Deserialize each key
		key := DeserializeTcPeditKey(b[next:])
		if key == nil {
			return nil, nil // Return nil if deserialization of the key fails
		}
		keys = append(keys, *key)
		next += SizeOfTcPeditKey
	}

	return x, keys
}
