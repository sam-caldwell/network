package core

import "encoding/binary"

// SerializeHfscCurve - Serializes the Curve structure into a byte slice.
//
// This function converts each field of the `Curve` struct (`m1`, `d`, `m2`)
// into a little-endian byte representation and copies them directly into a pre-allocated byte slice.
//
// Parameters:
// - c: Pointer to the `Curve` struct to be serialized.
//
// Returns:
// - A byte slice containing the serialized fields of the `Curve` struct.
func SerializeHfscCurve(c *Curve) (b []byte) {
	// Pre-allocate a byte slice to hold the serialized result (12 bytes total for m1, d, and m2).
	b = make([]byte, 12)

	// Copy the serialized fields directly into the byte slice using binary.LittleEndian encoding.
	binary.LittleEndian.PutUint32(b[0:], c.m1) // First 4 bytes for m1
	binary.LittleEndian.PutUint32(b[4:], c.d)  // Next 4 bytes for d
	binary.LittleEndian.PutUint32(b[8:], c.m2) // Last 4 bytes for m2

	// Return the serialized byte slice.
	return b
}
