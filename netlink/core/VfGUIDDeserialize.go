//go:build linux

package core

import (
	"encoding/binary"
	"errors"
)

// Deserialize - deserialize []byte to VfGUID struct
func (msg *VfGUID) Deserialize(b []byte) error {
	if len(b) < SizeofVfGUID {
		return errors.New("byte slice too short to deserialize VfGUID")
	}

	// Deserialize the Vf field
	msg.Vf = binary.LittleEndian.Uint32(b[0:4])

	// Deserialize the Rsvd field
	msg.Rsvd = binary.LittleEndian.Uint32(b[4:8])

	// Deserialize the GUID field
	msg.GUID = binary.LittleEndian.Uint64(b[8:16])

	return nil
}
