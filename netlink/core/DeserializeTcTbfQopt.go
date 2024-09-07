package core

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// DeserializeTcTbfQopt safely deserializes a byte slice into a TcTbfQopt object.
func DeserializeTcTbfQopt(b []byte) (*TcTbfQopt, error) {
	if len(b) < SizeOfTcTbfQopt {
		return nil, errors.New("DeserializeTcTbfQopt: insufficient byte slice size")
	}

	buf := bytes.NewReader(b)
	tbf := &TcTbfQopt{}

	// Use LittleEndian or BigEndian depending on your architecture
	err := binary.Read(buf, binary.LittleEndian, &tbf.Rate)
	if err != nil {
		return nil, err
	}

	err = binary.Read(buf, binary.LittleEndian, &tbf.Peakrate)
	if err != nil {
		return nil, err
	}

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
