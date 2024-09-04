package core

import "encoding/binary"

// Serialize - Converts the TcNetemCorrupt object into a byte slice using binary encoding.
func (msg *TcNetemCorrupt) Serialize() []byte {
	buf := make([]byte, SizeofTcNetemCorrupt)
	binary.LittleEndian.PutUint32(buf[0:], msg.Probability)
	binary.LittleEndian.PutUint32(buf[4:], msg.Correlation)
	return buf
}
