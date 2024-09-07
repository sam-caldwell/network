package core

import (
	"errors"
)

// Deserialize parses a byte slice into an RtAttr structure, including its data and any child attributes.
func (attr *RtAttr) Deserialize(b []byte) error {
	if len(b) < SizeOfUnixRtAttr {
		return errors.New("byte slice too short to contain RtAttr")
	}

	// Parse the RtAttr header
	attr.RtAttr.Len = NativeEndian.Uint16(b[0:2])
	attr.RtAttr.Type = NativeEndian.Uint16(b[2:4])

	// Validate length
	totalLen := int(attr.RtAttr.Len)
	if totalLen < SizeOfUnixRtAttr || totalLen > len(b) {
		return errors.New("invalid RtAttr length")
	}

	// Extract the Data part
	attr.Data = b[SizeOfUnixRtAttr:totalLen]

	// Initialize children slice
	attr.children = []NetlinkRequestData{}

	// Parse potential child attributes
	childOffset := rtaAlignOf(SizeOfUnixRtAttr + len(attr.Data))
	for childOffset < totalLen {
		child := RtAttr{}
		if err := child.Deserialize(b[childOffset:]); err != nil {
			return err
		}
		attr.children = append(attr.children, &child)
		childOffset += rtaAlignOf(int(child.RtAttr.Len))
	}

	return nil
}
