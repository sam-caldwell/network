package core

import "encoding/binary"

// Serialize - Converts the TcCsum struct into a byte slice.
// This method ensures safe serialization by manually converting each field to a byte slice using the binary encoding package.
//
// Returns the serialized byte slice.
func (msg *TcCsum) Serialize() []byte {
	buf := make([]byte, SizeofTcCsum)

	// Serialize fields
	binary.LittleEndian.PutUint32(buf[0:], msg.TcGen.Index)
	binary.LittleEndian.PutUint32(buf[4:], msg.TcGen.Capab)
	binary.LittleEndian.PutUint32(buf[8:], uint32(msg.TcGen.Action))
	binary.LittleEndian.PutUint32(buf[12:], uint32(msg.TcGen.Refcnt))
	binary.LittleEndian.PutUint32(buf[16:], uint32(msg.TcGen.Bindcnt))
	binary.LittleEndian.PutUint32(buf[20:], msg.UpdateFlags)

	return buf
}
