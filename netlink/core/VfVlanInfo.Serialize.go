package core

import (
	"bytes"
	"encoding/binary"
)

// Serialize - output the state of IfLaVfVlanInfoStruct as []byte
func (msg *IfLaVfVlanInfoStruct) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	// Serialize VfVlan
	if err := binary.Write(buf, binary.LittleEndian, msg.VfVlan); err != nil {
		return nil, err
	}

	// Serialize VlanProto
	if err := binary.Write(buf, binary.LittleEndian, msg.VlanProto); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
