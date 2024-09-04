package core

import (
	"encoding/binary"
)

// ConvertUint32ToBigEndian -  Flip bit order of little endian uint32 to big endian
func ConvertUint32ToBigEndian(i uint32) uint32 {
	const (
		byte0 = 0xff000000
		byte1 = 0x00ff0000
		byte2 = 0x0000ff00
		byte3 = 0x000000ff
	)
	if NativeEndian == binary.BigEndian {
		return i
	}
	return (i&byte0)>>24 |
		(i&byte1)>>8 |
		(i&byte2)<<8 |
		(i&byte3)<<24
}
