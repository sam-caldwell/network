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
//	// result will be "hello" (ignoring anything after the null terminator).
//
// If the input byte slice is nil, an empty string is returned.
//
// Returns:
// - A Go string representing the byte slice up to the null byte or the entire byte slice if no null byte is found.
func BytesToString(b []byte) string {
	if b == nil {
		return ""
	}

	// Find the index of the first null byte (0).
	// If no null byte is found, `bytes.Index` returns -1, and we use the entire slice length.
	n := bytes.Index(b, []byte{0})
	if n == -1 {
		n = len(b)
	}

	// Convert the slice of bytes up to the null byte (or the entire slice) into a string.
	return string(b[:n])
}
