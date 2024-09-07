package core

import "encoding/binary"

// DeserializeTcNetemRate - Converts a byte slice into a TcNetemRate object using binary encoding.
func DeserializeTcNetemRate(b []byte) *TcNetemRate {
	if len(b) < SizeOfTcRateSpec {
		return nil
	}
	msg := &TcNetemRate{}
	msg.Rate = binary.LittleEndian.Uint32(b[0:])
	msg.PacketOverhead = int32(binary.LittleEndian.Uint32(b[4:]))
	msg.CellSize = binary.LittleEndian.Uint32(b[8:])
	msg.CellOverhead = int32(binary.LittleEndian.Uint32(b[12:]))
	return msg
}
