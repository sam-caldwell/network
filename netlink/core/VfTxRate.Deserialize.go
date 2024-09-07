//go:build linux

package core

import (
	"errors"
)

// Deserialize - Deserialize a byte slice into a VfTxRate struct
func (msg *VfTxRate) Deserialize(b []byte) error {
	if len(b) < SizeOfVfTxRate {
		return errors.New("byte slice too short to deserialize VfTxRate")
	}

	// Deserialize the Vf field
	msg.Vf = NativeEndian.Uint32(b[0:4])

	// Deserialize the Rate field
	msg.Rate = NativeEndian.Uint32(b[4:8])

	return nil
}
