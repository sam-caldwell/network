package core

import (
	"encoding/binary"
)

// Serialize - serialize VfVlan to []byte
func (msg *VfVlan) Serialize() []byte {
	buf := make([]byte, SizeofVfVlan)

	// Serialize each field into the byte slice
	binary.LittleEndian.PutUint32(buf[0:4], msg.Vf)
	binary.LittleEndian.PutUint32(buf[4:8], msg.Vlan)
	binary.LittleEndian.PutUint32(buf[8:12], msg.Qos)

	return buf
}
