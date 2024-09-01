//go:build linux

package core

import (
	"encoding/binary"
	"errors"
)

// Deserialize - Deserialize a byte slice into a VfTxRate struct
func (msg *VfTxRate) Deserialize(b []byte) error {
	if len(b) < SizeofVfTxRate {
		return errors.New("byte slice too short to deserialize VfTxRate")
	}

	// Deserialize the Vf field
	msg.Vf = binary.LittleEndian.Uint32(b[0:4])

	// Deserialize the Rate field
	msg.Rate = binary.LittleEndian.Uint32(b[4:8])

	return nil
}
