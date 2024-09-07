package core

import "encoding/binary"

// DeserializeTcNetemCorrupt - Safely converts a byte slice into a TcNetemCorrupt object.
func DeserializeTcNetemCorrupt(b []byte) *TcNetemCorrupt {
	if len(b) < SizeOfTcNetemCorrupt {
		return nil // Return nil if the byte slice is too short
	}
	msg := &TcNetemCorrupt{}
	msg.Probability = binary.LittleEndian.Uint32(b[0:])
	msg.Correlation = binary.LittleEndian.Uint32(b[4:])
	return msg
}
