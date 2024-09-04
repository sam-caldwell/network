package core

import (
	"errors"
	"golang.org/x/sys/unix"
)

// DeserializeIfaCacheInfo converts a byte slice to an IfaCacheInfo struct
//
//	type IfaCacheinfo struct {
//	   Prefered uint32
//	   Valid    uint32
//	   Cstamp   uint32
//	   Tstamp   uint32
//	}
func DeserializeIfaCacheInfo(data []byte) (*IfaCacheInfo, error) {
	if len(data) < int(SizeOfIfaCacheinfo) {
		return nil, errors.New("data slice too short")
	}
	info := &IfaCacheInfo{
		unix.IfaCacheinfo{
			Prefered: NativeEndian.Uint32(data[0:4]),
			Valid:    NativeEndian.Uint32(data[4:8]),
			Cstamp:   NativeEndian.Uint32(data[8:12]),
			Tstamp:   NativeEndian.Uint32(data[12:16]),
		},
	}
	return info, nil
}
