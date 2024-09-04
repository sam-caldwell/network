package core

import "encoding/binary"

// DeserializeTcNetemReorder - Safely deserializes a byte slice into TcNetemReorder.
func DeserializeTcNetemReorder(b []byte) *TcNetemReorder {

	return &TcNetemReorder{
		Probability: binary.LittleEndian.Uint32(b[0:4]),
		Correlation: binary.LittleEndian.Uint32(b[4:8]),
	}

}
