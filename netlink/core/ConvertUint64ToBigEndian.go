package core

import (
	"encoding/binary"
)

// ConvertUint64ToBigEndian - Flip bit order of a little-endian uint64 to big-endian.
func ConvertUint64ToBigEndian(i uint64) uint64 {
	const (
		byte0 uint64 = 0xff00000000000000
		byte1 uint64 = 0x00ff000000000000
		byte2 uint64 = 0x0000ff0000000000
		byte3 uint64 = 0x000000ff00000000
		byte4 uint64 = 0x00000000ff000000
		byte5 uint64 = 0x0000000000ff0000
		byte6 uint64 = 0x000000000000ff00
		byte7 uint64 = 0x00000000000000ff
	)

	if binary.BigEndian == NativeEndian {
		return i
	}

	return (i&byte0)>>56 |
		(i&byte1)>>40 |
		(i&byte2)>>24 |
		(i&byte3)>>8 |
		(i&byte4)<<8 |
		(i&byte5)<<24 |
		(i&byte6)<<40 |
		(i&byte7)<<56
}
