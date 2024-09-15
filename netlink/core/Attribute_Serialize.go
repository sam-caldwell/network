package core

import (
	"encoding/binary"
	"fmt"
)

// Serialize encodes the Attribute into a byte slice following the Netlink TLV (Type-Length-Value) format.
// It includes the attribute header (Type and Length), the Value, and any necessary padding to align to 4 bytes.
//
// Output Format:
//
//		0     2     4          n
//	 |--|--|--|--|----------|
//	 | Len | Type| Value... |
func (attr *Attribute) Serialize() ([]byte, error) {
	valueLen := len(attr.Value)
	nlaLen := 4 + valueLen // 4 bytes for the header (Type and Length)
	if nlaLen > 0xFFFF {
		return nil, fmt.Errorf("attribute too large: length %d exceeds maximum of %d", nlaLen, 0xFFFF)
	}

	// Align the total length to a 4-byte boundary as required by Netlink attributes
	alignedLen := rtaAlignOf(nlaLen)
	// paddingLen := alignedLen - nlaLen  // This is the number of padding bits (already accounted for)

	// Initialize a byte slice with the total aligned length
	b := make([]byte, alignedLen)

	// Write the attribute length and type in little-endian format
	binary.LittleEndian.PutUint16(b[0:2], uint16(nlaLen))
	binary.LittleEndian.PutUint16(b[2:4], uint16(attr.Type))

	// Copy the Value into the byte slice
	copy(b[4:], attr.Value)

	// The padding bytes are already zeroed by make(), no need to set them explicitly

	return b, nil
}
