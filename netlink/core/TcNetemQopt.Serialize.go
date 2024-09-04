package core

import (
	"encoding/binary"
)

// Serialize - Serialize the TcNetemQopt struct into a byte slice.
// This function ensures safe serialization by manually converting each field
// to a byte slice using the binary encoding package.
//
// Returns the serialized byte slice.
func (msg *TcNetemQopt) Serialize() []byte {
	buf := make([]byte, SizeofTcNetemQopt)
	binary.LittleEndian.PutUint32(buf[0:], msg.Latency)
	binary.LittleEndian.PutUint32(buf[4:], msg.Limit)
	binary.LittleEndian.PutUint32(buf[8:], msg.Loss)
	binary.LittleEndian.PutUint32(buf[12:], msg.Gap)
	binary.LittleEndian.PutUint32(buf[16:], msg.Duplicate)
	binary.LittleEndian.PutUint32(buf[20:], msg.Jitter)
	return buf
}
