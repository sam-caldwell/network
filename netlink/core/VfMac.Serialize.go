//go:build linux

package core

import (
	"encoding/binary"
)

// Serialize - Serialize VfMac to []byte
func (msg *VfMac) Serialize() []byte {
	buf := make([]byte, SizeofVfMac)

	// Serialize the Vf field
	binary.LittleEndian.PutUint32(buf[0:4], msg.Vf)

	// Serialize the Mac field
	copy(buf[4:], msg.Mac[:])

	return buf
}
