//go:build linux

package core

import (
	"errors"
)

// DeserializeVfTrust - Deserialize a byte slice into a VfTrust struct
func DeserializeVfTrust(b []byte) (*VfTrust, error) {
	if len(b) < SizeofVfTrust {
		return nil, errors.New("input too short to deserialize VfTrust")
	}

	vfTrust := &VfTrust{}

	// Deserialize the Vf field
	vfTrust.Vf = NativeEndian.Uint32(b[0:4])

	// Deserialize the Setting field
	vfTrust.Setting = NativeEndian.Uint32(b[4:8])

	return vfTrust, nil
}
