package core

import (
	"encoding/binary"
)

// ConvertUint32ToBigEndian -  Flip bit order of little endian uint32 to big endian
func ConvertUint32ToBigEndian(i uint32) uint32 {
	if nativeEndian == binary.BigEndian {
		return i
	}
	return (i&0xff000000)>>24 |
		(i&0xff0000)>>8 |
		(i&0xff00)<<8 |
		(i&0xff)<<24
}
