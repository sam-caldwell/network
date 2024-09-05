package core

import (
	"bytes"
	"encoding/binary"
)

// Serialize - convert the VfLinkState structure into a byte slice.
func (msg *VfLinkState) Serialize() ([]byte, error) {

	buf := new(bytes.Buffer)

	if err := binary.Write(buf, NativeEndian, msg); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil

}
