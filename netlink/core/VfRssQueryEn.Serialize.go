package core

import (
	"bytes"
	"encoding/binary"
)

// Serialize - convert the VfRssQueryEn structure into a byte slice.
func (msg *VfRssQueryEn) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, NativeEndian, msg); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
