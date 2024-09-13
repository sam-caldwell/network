package core

import "encoding/binary"

// Len returns the length of the GenericNetlinkMessage structure in bytes.
// It calculates the length using the binary.Size function, which ensures safe and portable size calculation.
func (msg *GenericNetlinkMessage) Len() int {
	return binary.Size(*msg)
}
