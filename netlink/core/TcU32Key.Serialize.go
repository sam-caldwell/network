package core

import (
	"bytes"
	"encoding/binary"
)

// Serialize - converts the TcU32Key struct into a byte slice.
func (msg *TcU32Key) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	// Serialize each field with proper endianness
	if err := binary.Write(buf, binary.BigEndian, msg.Mask); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.BigEndian, msg.Val); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, NativeEndian, msg.Off); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, NativeEndian, msg.OffMask); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
