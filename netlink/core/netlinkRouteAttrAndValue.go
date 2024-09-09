package core

import (
	"golang.org/x/sys/unix"
)

// netlinkRouteAttrAndValue parses a byte slice into an RtAttr struct
// and returns the attribute, its value, the aligned length of the attribute,
// and an error if the byte slice is malformed or too short.
//
// Parameters:
// - b []byte: The byte slice to be parsed.
//
// Returns:
// - *RtAttr: A pointer to the parsed RtAttr struct.
// - []byte: The value part of the attribute (excluding the header).
// - int: The aligned length of the RtAttr (after alignment for word boundaries).
// - error: Returns an error if the input slice is too short or contains invalid data.
func netlinkRouteAttrAndValue(b []byte) (*RtAttr, []byte, int, error) {
	// Initialize the RtAttr struct
	var attr RtAttr

	// Deserialize the byte slice into the RtAttr struct.
	// The Deserialize method should be defined to extract the unix.RtAttr fields (Len, Type) from the byte slice.
	if err := attr.Deserialize(b); err != nil {
		return nil, nil, 0, err
	}

	// Ensure that the length of the attribute is valid.
	// It must be at least the size of the RtAttr header (SizeOfUnixRtAttr).
	// It must also not exceed the size of the provided byte slice.
	if int(attr.Len()) < SizeOfUnixRtAttr || int(attr.Len()) > len(b) {
		return nil, nil, 0, unix.EINVAL
	}

	// Return the parsed attribute, the value portion of the byte slice (after the header),
	// and the aligned size of the attribute using rtaAlignOf, which aligns lengths to the next multiple of 4 bytes.
	// The aligned size ensures proper memory alignment for network protocols.
	return &attr, b[SizeOfUnixRtAttr:], rtaAlignOf(int(attr.Len())), nil
}
