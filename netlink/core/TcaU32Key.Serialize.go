package core

import (
	"bytes"
	"encoding/binary"
)

// Serialize - Converts the TcU32Key struct into a byte slice in a safe way.
// This function ensures that each field is serialized with the correct endianness and alignment.
func (msg *TcU32Key) Serialize() []byte {
	buf := new(bytes.Buffer)

	// Serialize each field with proper endianness
	_ = binary.Write(buf, binary.BigEndian, msg.Mask)
	_ = binary.Write(buf, binary.BigEndian, msg.Val)
	_ = binary.Write(buf, binary.LittleEndian, msg.Off)
	_ = binary.Write(buf, binary.LittleEndian, msg.OffMask)

	return buf.Bytes()
}
