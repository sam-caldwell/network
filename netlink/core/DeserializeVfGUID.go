//go:build linux

package core

import (
	"errors"
)

// DeserializeVfGUID - Deserialize a byte slice into a VfGUID struct and return by reference
func DeserializeVfGUID(b []byte) (*VfGUID, error) {
	if len(b) < SizeofVfGUID {
		return nil, errors.New("byte slice too short to deserialize VfGUID")
	}

	vfGUID := &VfGUID{}

	// Deserialize the Vf field
	vfGUID.Vf = NativeEndian.Uint32(b[0:4])

	// Deserialize the Rsvd field
	vfGUID.Rsvd = NativeEndian.Uint32(b[4:8])

	// Deserialize the GUID field
	vfGUID.GUID = NativeEndian.Uint64(b[8:16])

	return vfGUID, nil
}
