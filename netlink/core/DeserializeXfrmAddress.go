package core

import (
	"errors"
)

// DeserializeXfrmAddress safely converts a byte slice into an XfrmAddress structure.
// This function checks if the length of the input byte slice matches the expected size of the XfrmAddress structure.
//
// If the length is incorrect, it returns an error to ensure safety.
//
// Parameters:
// - `b []byte`: The byte slice to deserialize.
//
// Returns:
// - `*XfrmAddress`: A pointer to the deserialized XfrmAddress.
// - `error`: An error if the input byte slice is too short.
func DeserializeXfrmAddress(b []byte) (*XfrmAddress, error) {
	// Ensure that the byte slice is the correct length
	if len(b) < SizeOfXfrmAddress {
		return nil, errors.New("DeserializeXfrmAddress: byte slice too short")
	}

	// Create an XfrmAddress instance
	var addr XfrmAddress

	// Safely copy the byte slice into the XfrmAddress structure
	copy(addr[:], b[:SizeOfXfrmAddress])

	return &addr, nil
}
