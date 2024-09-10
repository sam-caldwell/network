package core

import "encoding/binary"

// DeserializeHfscCurve - Converts a byte slice into a Curve structure.
//
// This function reads the first 12 bytes of the input byte slice and deserializes
// them into the `m1`, `d`, and `m2` fields of the `Curve` struct. It assumes the
// data is encoded in little-endian byte order.
//
// Parameters:
// - b: The byte slice to deserialize.
//
// Returns:
// - A pointer to a `Curve` struct with the fields populated from the byte slice.
func DeserializeHfscCurve(b []byte) *Curve {
	return &Curve{
		m1: binary.LittleEndian.Uint32(b[0:4]),  // First 4 bytes represent `m1`
		d:  binary.LittleEndian.Uint32(b[4:8]),  // Next 4 bytes represent `d`
		m2: binary.LittleEndian.Uint32(b[8:12]), // Final 4 bytes represent `m2`
	}
}
