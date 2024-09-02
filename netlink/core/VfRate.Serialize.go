package core

import (
	"bytes"
	"encoding/binary"
)

// Serialize - convert the VfRate struct into a byte slice, which can be used in netlink messages.
func (v *VfRate) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
