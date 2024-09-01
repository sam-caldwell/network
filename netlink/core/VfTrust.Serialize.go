//go:build linux

package core

import (
	"encoding/binary"
)

// Serialize - Serialize VfTrust to []byte
func (msg *VfTrust) Serialize() []byte {
	buf := make([]byte, SizeofVfTrust)

	// Serialize the Vf field
	binary.LittleEndian.PutUint32(buf[0:4], msg.Vf)

	// Serialize the Setting field
	binary.LittleEndian.PutUint32(buf[4:8], msg.Setting)

	return buf
}
