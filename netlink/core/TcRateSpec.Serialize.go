package core

import "encoding/binary"

// Serialize - Serialize the TcRateSpec struct into a byte slice.
// This function ensures safe serialization by manually converting each field
// to a byte slice using the binary encoding package.
func (msg *TcRateSpec) Serialize() []byte {
	buf := make([]byte, SizeofTcRateSpec)
	buf[0] = msg.CellLog   // Directly assign for uint8
	buf[1] = msg.Linklayer // Directly assign for uint8
	binary.LittleEndian.PutUint16(buf[2:], msg.Overhead)
	binary.LittleEndian.PutUint16(buf[4:], uint16(msg.CellAlign))
	binary.LittleEndian.PutUint16(buf[6:], msg.Mpu)
	binary.LittleEndian.PutUint32(buf[8:], msg.Rate)
	return buf
}
