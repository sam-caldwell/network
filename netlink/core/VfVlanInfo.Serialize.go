package core

import (
	"encoding/binary"
	"unsafe"
)

// Serialize - serialize VfVlanInfo
func (msg *VfVlanInfo) Serialize() []byte {
	return (*(*[SizeofVfVlanInfo]byte)(unsafe.Pointer(msg)))[:]
}

// DeserializeVfVlanInfo - deserialize VfVlanInfo
func DeserializeVfVlanInfo(b []byte) *VfVlanInfo {
	return &VfVlanInfo{
		*(*VfVlan)(unsafe.Pointer(&b[0:SizeofVfVlan][0])),
		binary.BigEndian.Uint16(b[SizeofVfVlan:SizeofVfVlanInfo]),
	}
}
