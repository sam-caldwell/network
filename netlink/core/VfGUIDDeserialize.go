//go:build linux

package core

import (
	"errors"
)

// Deserialize - deserialize []byte to VfGUID struct
func (msg *VfGUID) Deserialize(b []byte) error {
	if len(b) < SizeofVfGUID {
		return errors.New("byte slice too short to deserialize VfGUID")
	}

	// Deserialize the Vf field
	msg.Vf = NativeEndian.Uint32(b[0:4])

	// Deserialize the Rsvd field
	msg.Rsvd = NativeEndian.Uint32(b[4:8])

	// Deserialize the GUID field
	msg.GUID = NativeEndian.Uint64(b[8:16])

	return nil
}
