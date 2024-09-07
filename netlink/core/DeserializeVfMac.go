package core

import (
	"errors"
)

// DeserializeVfMac - Deserialize a byte slice into a VfMac struct
func DeserializeVfMac(b []byte) (*VfMac, error) {

	if len(b) < SizeofVfMac {
		return nil, errors.New("input too short to deserialize VfMac")
	}

	vfMac := &VfMac{}

	// Deserialize the Vf field
	vfMac.Vf = NativeEndian.Uint32(b[0:4])

	// Deserialize the Mac field
	copy(vfMac.Mac[:], b[4:36])

	return vfMac, nil
}
