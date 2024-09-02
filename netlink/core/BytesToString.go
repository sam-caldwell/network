//go:build linux

package core

import (
	"bytes"
)

// BytesToString converts a null-terminated byte slice into a Go string.
//
// This function takes a byte slice `b` and looks for the first occurrence of a null byte (`0`).
// It then converts the portion of the byte slice before the null byte into a Go string.
// This is particularly useful when working with byte slices that represent C-style strings,
// where the end of the string is marked by a null terminator.
//
// If no null byte is found, the entire byte slice is converted into a string.
//
// Example:
//
//	b := []byte{'h', 'e', 'l', 'l', 'o', 0, 'x', 'y', 'z'}
//	result := BytesToString(b)
//	// result will be "hello" (ignoring anything after the null terminator)
func BytesToString(b []byte) string {
	// Find the index of the first null byte (0).
	n := bytes.Index(b, []byte{0})

	// Convert the slice of bytes before the null byte into a string.
	return string(b[:n])
}
