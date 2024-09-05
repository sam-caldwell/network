package core

import (
	"bytes"
	"encoding/binary"
)

// Serialize - Converts the TcConnmark structure into a byte slice.
// This method uses binary encoding to safely serialize the structure fields.
// It avoids using unsafe.Pointer by manually converting each field to bytes.
//
// Returns a serialized byte slice.
func (msg *TcConnmark) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, msg.Index); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.LittleEndian, msg.Capab); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.LittleEndian, msg.Action); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.LittleEndian, msg.Refcnt); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.LittleEndian, msg.Bindcnt); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.LittleEndian, msg.Zone); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
