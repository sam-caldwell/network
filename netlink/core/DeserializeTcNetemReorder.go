package core

import (
	"encoding/binary"
	"errors"
)

// DeserializeTcNetemReorder - Safely deserializes a byte slice into TcNetemReorder.
//
//	type TcNetemReorder struct {
//	   Probability uint32
//	   Correlation uint32
//	}
func DeserializeTcNetemReorder(b []byte) (*TcNetemReorder, error) {
	if b == nil || len(b) < 8 {
		return nil, errors.New("input too small")
	}

	return &TcNetemReorder{

		Probability: binary.LittleEndian.Uint32(b[0:4]),

		Correlation: binary.LittleEndian.Uint32(b[4:8]),
	}, nil

}
