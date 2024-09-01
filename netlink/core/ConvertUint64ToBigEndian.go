package core

import (
	"encoding/binary"
)

// ConvertUin64ToBigEndian - Flip bit order of little endian uint64 to big endian
func ConvertUin64ToBigEndian(i uint64) uint64 {
	if nativeEndian == binary.BigEndian {
		return i
	}
	return (i&0xff00000000000000)>>56 |
		(i&0x00ff000000000000)>>40 |
		(i&0x0000ff0000000000)>>24 |
		(i&0x000000ff00000000)>>8 |
		(i&0x00000000ff000000)<<8 |
		(i&0x0000000000ff0000)<<24 |
		(i&0x000000000000ff00)<<40 |
		(i&0x00000000000000ff)<<56
}
