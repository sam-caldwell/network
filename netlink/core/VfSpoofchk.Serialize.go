package core

import (
	"bytes"
	"encoding/binary"
)

// Serialize - convert VfSpoofchk structure into a []byte slice.
func (msg *VfSpoofchk) Serialize() ([]byte, error) {

	buf := new(bytes.Buffer)

	if err := binary.Write(buf, binary.LittleEndian, msg); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil

}
