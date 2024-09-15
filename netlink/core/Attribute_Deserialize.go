package core

import (
	"encoding/binary"
	"fmt"
)

// Deserialize parses a byte slice representing a Netlink attribute and returns an Attribute instance.
// It expects the byte slice to be in Netlink TLV (Type-Length-Value) format, including any padding.
func Deserialize(data []byte) (*Attribute, error) {
	if len(data) < 4 {
		return nil, fmt.Errorf("data too short to contain an attribute header")
	}

	// Read the attribute length and type from the first 4 bytes
	nlaLen := binary.LittleEndian.Uint16(data[0:2])
	nlaType := binary.LittleEndian.Uint16(data[2:4])

	// Validate the length
	if nlaLen < 4 || int(nlaLen) > len(data) {
		return nil, fmt.Errorf("invalid attribute length: %d", nlaLen)
	}

	// Extract the Value, excluding the padding
	valueLen := int(nlaLen) - 4
	value := make([]byte, valueLen)
	copy(value, data[4:4+valueLen])

	attr := &Attribute{
		Type:  NlaFlags(nlaType),
		Value: value,
	}

	return attr, nil
}
