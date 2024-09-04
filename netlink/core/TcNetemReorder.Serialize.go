package core

import "encoding/binary"

// Serialize - Safely serializes TcNetemReorder into a byte slice using binary encoding.
func (msg *TcNetemReorder) Serialize() []byte {

	buf := make([]byte, SizeofTcNetemReorder)

	// Serialize each field
	binary.LittleEndian.PutUint32(buf[0:], msg.Probability)
	binary.LittleEndian.PutUint32(buf[4:], msg.Correlation)

	return buf

}
