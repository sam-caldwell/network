package core

import (
	"encoding/binary"
	"errors"
	"golang.org/x/sys/unix"
)

// Deserialize parses a byte slice into an RtAttr structure, including its data and any child attributes.
func (attr *RtAttr) Deserialize(b []byte) error {
	if len(b) < unix.SizeofRtAttr {
		return errors.New("byte slice too short to contain RtAttr")
	}

	// Parse the RtAttr header
	attr.RtAttr.Len = binary.LittleEndian.Uint16(b[0:2])
	attr.RtAttr.Type = binary.LittleEndian.Uint16(b[2:4])

	// Validate length
	totalLen := int(attr.RtAttr.Len)
	if totalLen < unix.SizeofRtAttr || totalLen > len(b) {
		return errors.New("invalid RtAttr length")
	}

	// Extract the Data part
	attr.Data = b[unix.SizeofRtAttr:totalLen]

	// Initialize children slice
	attr.children = []NetlinkRequestData{}

	// Parse potential child attributes
	childOffset := rtaAlignOf(unix.SizeofRtAttr + len(attr.Data))
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
