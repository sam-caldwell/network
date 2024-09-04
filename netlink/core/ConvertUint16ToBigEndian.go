package core

import "encoding/binary"

// ConvertUint16ToBigEndian - Flip bit order of little endian uint16 to big endian
func ConvertUint16ToBigEndian(i uint16) uint16 {
	const (
		MsbMask = 0xFF00
		LsbMask = 0x00FF
	)

	if NativeEndian == binary.BigEndian {
		return i
	}
	return (i&MsbMask)>>8 | (i&LsbMask)<<8
}
