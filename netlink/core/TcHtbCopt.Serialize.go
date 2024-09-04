package core

import (
	"encoding/binary"
)

// Serialize - Serializes the TcHtbCopt structure into a byte slice using binary encoding.
// This method uses the binary package to safely convert each field into its byte representation.
func (msg *TcHtbCopt) Serialize() []byte {
	buf := make([]byte, SizeofTcHtbCopt)

	// Serialize the Rate field
	copy(buf[0:], msg.Rate.Serialize())

	// Serialize the Ceil field, which starts after Rate
	copy(buf[SizeofTcRateSpec:], msg.Ceil.Serialize())

	// Serialize the remaining fields (Buffer, Cbuffer, Quantum, Level, Prio)
	binary.LittleEndian.PutUint32(buf[2*SizeofTcRateSpec:], msg.Buffer)
	binary.LittleEndian.PutUint32(buf[2*SizeofTcRateSpec+4:], msg.Cbuffer)
	binary.LittleEndian.PutUint32(buf[2*SizeofTcRateSpec+8:], msg.Quantum)
	binary.LittleEndian.PutUint32(buf[2*SizeofTcRateSpec+12:], msg.Level)
	binary.LittleEndian.PutUint32(buf[2*SizeofTcRateSpec+16:], msg.Prio)

	return buf
}
