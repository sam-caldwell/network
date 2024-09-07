package core

import (
	"errors"
)

// Deserialize - Deserialize a byte slice into a VfMac struct
func (msg *VfMac) Deserialize(b []byte) error {

	if len(b) < SizeofVfMac {
		return errors.New("byte slice too short to deserialize VfMac")
	}

	// Deserialize the Vf field
	msg.Vf = NativeEndian.Uint32(b[0:4])

	// Deserialize the Mac field
	copy(msg.Mac[:], b[4:36])

	return nil
}
