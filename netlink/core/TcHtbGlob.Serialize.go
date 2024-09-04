package core

import (
	"bytes"
	"encoding/binary"
)

// Serialize - Safely serializes the TcHtbGlob struct into a byte slice.
func (msg *TcHtbGlob) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	// Using LittleEndian encoding, as Linux generally uses this on most platforms.
	err := binary.Write(buf, NativeEndian, msg)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
