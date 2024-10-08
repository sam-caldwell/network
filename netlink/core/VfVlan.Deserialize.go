package core

import (
	"errors"
)

// DeserializeVfVlan - deserialize []byte to VfVlan struct
func DeserializeVfVlan(b []byte) (*VfVlan, error) {
	if len(b) < SizeOfVfVlan {
		return nil, errors.New("byte slice too short to deserialize VfVlan")
	}

	vfVlan := &VfVlan{
		Vf:   NativeEndian.Uint32(b[0:4]),
		Vlan: NativeEndian.Uint32(b[4:8]),
		Qos:  NativeEndian.Uint32(b[8:12]),
	}

	return vfVlan, nil
}
