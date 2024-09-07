package core

import "encoding/binary"

// DeserializeTcRateSpec - Safely deserialize a byte slice into a TcRateSpec struct.
// This function reads the byte slice and converts it into a TcRateSpec struct using the binary package.
//
// Parameters:
// - b: The byte slice to deserialize.
//
// Returns:
// - A pointer to the deserialized TcRateSpec struct.
func DeserializeTcRateSpec(b []byte) *TcRateSpec {
	if len(b) < SizeOfTcRateSpec {
		return nil
	}

	msg := &TcRateSpec{}
	msg.CellLog = b[0]
	msg.Linklayer = b[1]
	msg.Overhead = binary.LittleEndian.Uint16(b[2:])
	msg.CellAlign = int16(binary.LittleEndian.Uint16(b[4:]))
	msg.Mpu = binary.LittleEndian.Uint16(b[6:])
	msg.Rate = binary.LittleEndian.Uint32(b[8:])
	return msg
}
