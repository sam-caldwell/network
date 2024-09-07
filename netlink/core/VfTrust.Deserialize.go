//go:build linux

package core

import (
	"errors"
)

// Deserialize - Serialize VfTrust from []byte
func (msg *VfTrust) Deserialize(b []byte) error {
	if len(b) < SizeOfVfTrust {
		return errors.New("byte slice too short to deserialize VfTrust")
	}
	// Deserialize the Vf field
	msg.Vf = NativeEndian.Uint32(b[0:4])

	// Deserialize the Setting field
	msg.Setting = NativeEndian.Uint32(b[4:8])

	return nil
}
