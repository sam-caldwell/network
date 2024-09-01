//go:build linux

package core

import (
	"encoding/binary"
)

// Serialize - Serialize VfTxRate to []byte
func (msg *VfTxRate) Serialize() []byte {
	buf := make([]byte, SizeofVfTxRate)

	// Serialize the Vf field
	binary.LittleEndian.PutUint32(buf[0:4], msg.Vf)

	// Serialize the Rate field
	binary.LittleEndian.PutUint32(buf[4:8], msg.Rate)

	return buf
}
