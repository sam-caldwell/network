package core

import (
	"errors"
)

// Deserialize - Deserialize a given byte slice from position 0 for the first attribute and return the end position.
func (attr *Attribute) Deserialize(data []byte) (endPos int, err error) {
	const (
		attributeHeaderSize = 4 // Assuming the header size is 4 bytes (2 bytes for length and 2 for type)
		invalidPosition     = -1
	)

	if len(data) < attributeHeaderSize {
		return invalidPosition, errors.New(ErrInputTooShort)
	}

	length := int(NativeEndian.Uint16(data[0:2]))
	if length < attributeHeaderSize || length > len(data) {
		return invalidPosition, errors.New(ErrInvalidAttributeLength)
	}

	attr.Type = NlaFlags(NativeEndian.Uint16(data[2:4]))
	attr.Value = data[attributeHeaderSize:length]

	// Calculate the ending position (aligned to the next 4-byte boundary)
	endPos = rtaAlignOf(length)

	return endPos, nil
}
