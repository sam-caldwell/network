package core

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// Serialize - Serializes the TcMsg struct into a byte slice using binary encoding.
func (msg *TcMsg) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, NativeEndian, msg)
	if err != nil {
		return nil, errors.New("failed to serialize TcMsg")
	}
	return buf.Bytes(), nil
}
