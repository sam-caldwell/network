package core

import "unsafe"

// DeserializeTcPeditKey safely deserializes a byte slice into a TcPeditKey structure.
// It checks the length of the byte slice to ensure that it is at least the size of TcPeditKey
// to prevent out-of-bounds memory access.
//
// Returns nil if the byte slice is too small.
func DeserializeTcPeditKey(b []byte) *TcPeditKey {
	if len(b) < SizeOfTcPeditKey {
		return nil // Return nil if the byte slice is too small
	}
	return (*TcPeditKey)(unsafe.Pointer(&b[0])) // Safely convert the byte slice to a TcPeditKey pointer
}
