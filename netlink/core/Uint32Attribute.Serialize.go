package core

import "encoding/binary"

// Serialize converts the Uint32Attribute into a byte slice, ready for transmission or storage.
func (a *Uint32Attribute) Serialize() ([]byte, error) {

	// Create a byte slice buffer, aligned to the required size, to hold the serialized data.
	// The buffer size is calculated using rtaAlignOf(8), assuming 8 bytes for alignment.
	buf := make([]byte, rtaAlignOf(8))

	// Store the length of the attribute (8 bytes) in the first 2 bytes of the buffer.
	NativeEndian.PutUint16(buf[0:2], 8)

	// Store the attribute type identifier in the next 2 bytes of the buffer.
	NativeEndian.PutUint16(buf[2:4], a.Type)

	// Check if the NlaFNetByteorder flag is set in the type identifier.
	// If set, store the 32-bit value in big-endian order (network byte order).
	if a.Type&uint16(NlaFNetByteorder) != 0 {
		binary.BigEndian.PutUint32(buf[4:], a.Value)
	} else {
		// Otherwise, store the 32-bit value in the native endianness of the system.
		NativeEndian.PutUint32(buf[4:], a.Value)
	}

	// Return the serialized byte slice.
	return buf, nil
}
