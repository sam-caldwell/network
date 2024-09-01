//go:build linux

package core

import (
	"encoding/binary"
	"errors"
)

// DeserializeVfTxRate - Deserialize a byte slice into a VfTxRate struct and return by reference
func DeserializeVfTxRate(b []byte) (*VfTxRate, error) {
	if len(b) < SizeofVfTxRate {
		return nil, errors.New("byte slice too short to deserialize VfTxRate")
	}

	vfTxRate := &VfTxRate{}

	// Deserialize the Vf field
	vfTxRate.Vf = binary.LittleEndian.Uint32(b[0:4])

	// Deserialize the Rate field
	vfTxRate.Rate = binary.LittleEndian.Uint32(b[4:8])

	return vfTxRate, nil
}
