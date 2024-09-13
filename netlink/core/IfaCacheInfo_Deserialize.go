package core

import (
	"errors"
)

// Deserialize method converts a byte slice to an IfaCacheInfo struct
//
//	type IfaCacheinfo struct {
//	   Prefered uint32
//	   Valid    uint32
//	   Cstamp   uint32
//	   Tstamp   uint32
//	}
func (msg *IfaCacheInfo) Deserialize(data []byte) error {
	if len(data) < IfaCacheInfoSize {
		return errors.New(ErrInputTooShort)
	}
	msg.Prefered = NativeEndian.Uint32(data[0:4])
	msg.Valid = NativeEndian.Uint32(data[4:8])
	msg.Cstamp = NativeEndian.Uint32(data[8:12])
	msg.Tstamp = NativeEndian.Uint32(data[12:16])
	return nil
}

// DeserializeIfaCacheInfo converts a byte slice to an IfaCacheInfo struct
func DeserializeIfaCacheInfo(data []byte) (*IfaCacheInfo, error) {
	if err := checkInputSize(data, IfaCacheInfoSize, IfaCacheInfoSize); err != nil {
		return nil, err
	}
	var info IfaCacheInfo
	if err := info.Deserialize(data); err != nil {
		return nil, err
	}
	return &info, nil
}
