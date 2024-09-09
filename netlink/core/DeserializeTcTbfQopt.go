package core

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// DeserializeTcTbfQopt safely deserializes a byte slice into a TcTbfQopt object.
//
//	type TcTbfQopt struct {
//	   Rate     TcRateSpec
//	   Peakrate TcRateSpec
//	   Limit    uint32
//	   Buffer   uint32
//	   Mtu      uint32
//	}
func DeserializeTcTbfQopt(b []byte) (*TcTbfQopt, error) {
	if len(b) < SizeOfTcTbfQopt {
		return nil, errors.New("DeserializeTcTbfQopt: insufficient byte slice size")
	}

	tbf := &TcTbfQopt{}

	temp, err := DeserializeTcRateSpec(b[0:SizeOfTcRateSpec])
	if err != nil {
		return nil, err
	} else {
		tbf.Rate = *temp
	}
	temp, err = DeserializeTcRateSpec(b[SizeOfTcRateSpec : 2*SizeOfTcRateSpec])
	if err != nil {
		return nil, err
	} else {
		tbf.Peakrate = *temp
	}

	buf := bytes.NewReader(b[2*SizeOfTcRateSpec:])
	err = binary.Read(buf, binary.LittleEndian, &tbf.Limit)
	if err != nil {
		return nil, err
	}

	err = binary.Read(buf, binary.LittleEndian, &tbf.Buffer)
	if err != nil {
		return nil, err
	}

	err = binary.Read(buf, binary.LittleEndian, &tbf.Mtu)
	if err != nil {
		return nil, err
	}

	return tbf, nil
}
