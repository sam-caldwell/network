package core

import "encoding/binary"

// DeserializeTcNetemCorr deserializes a byte slice into TcNetemCorr.
func DeserializeTcNetemCorr(b []byte) *TcNetemCorr {
	return &TcNetemCorr{
		DelayCorr: binary.LittleEndian.Uint32(b[0:4]),
		LossCorr:  binary.LittleEndian.Uint32(b[4:8]),
		DupCorr:   binary.LittleEndian.Uint32(b[8:12]),
	}
}
