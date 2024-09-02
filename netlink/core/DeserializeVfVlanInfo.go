package core

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// DeserializeVfVlanInfo deserializes a byte slice into an IfLaVfVlanInfoStruct.
func DeserializeVfVlanInfo(b []byte) (*IfLaVfVlanInfoStruct, error) {
	if len(b) < SizeofVfVlanInfo {
		return nil, errors.New("byte slice is too short to contain IfLaVfVlanInfoStruct")
	}

	// Deserialize VfVlan
	vfVlan := VfVlan{}
	vfVlanLen := SizeofVfVlan // Assuming SizeofVfVlan is the correct size for VfVlan struct
	if err := binary.Read(bytes.NewReader(b[:vfVlanLen]), binary.BigEndian, &vfVlan); err != nil {
		return nil, err
	}

	// Deserialize VlanProto
	vlanProto := binary.BigEndian.Uint16(b[vfVlanLen:SizeofVfVlanInfo])

	return &IfLaVfVlanInfoStruct{
		VfVlan:    vfVlan,
		VlanProto: vlanProto,
	}, nil
}
