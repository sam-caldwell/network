package core

import (
	"bytes"
	"encoding/binary"
)

// Serialize - serialize an interface address message as a byte slice.
func (msg *IfAddressMessage) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	// Write each field to the buffer in the correct order and format
	if err := binary.Write(buf, binary.LittleEndian, msg.Family); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.LittleEndian, msg.Prefixlen); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.LittleEndian, msg.Flags); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.LittleEndian, msg.Scope); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.LittleEndian, msg.Index); err != nil {
		return nil, err
	}

	// Return the byte slice
	return buf.Bytes(), nil
}
