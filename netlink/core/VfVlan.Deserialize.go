package core

import (
	"encoding/binary"
	"errors"
)

// DeserializeVfVlan - deserialize []byte to VfVlan struct
func DeserializeVfVlan(b []byte) (*VfVlan, error) {
	if len(b) < SizeofVfVlan {
		return nil, errors.New("byte slice too short to deserialize VfVlan")
	}

	vfVlan := &VfVlan{
		Vf:   binary.LittleEndian.Uint32(b[0:4]),
		Vlan: binary.LittleEndian.Uint32(b[4:8]),
		Qos:  binary.LittleEndian.Uint32(b[8:12]),
	}

	return vfVlan, nil
}
