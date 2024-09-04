package core

import (
	"encoding/binary"
)

// Serialize - Converts the TcNetemRate structure into a byte slice using binary encoding.
func (msg *TcNetemRate) Serialize() []byte {
	buf := make([]byte, SizeofTcRateSpec)
	binary.LittleEndian.PutUint32(buf[0:], msg.Rate)
	binary.LittleEndian.PutUint32(buf[4:], uint32(msg.PacketOverhead))
	binary.LittleEndian.PutUint32(buf[8:], msg.CellSize)
	binary.LittleEndian.PutUint32(buf[12:], uint32(msg.CellOverhead))
	return buf
}
