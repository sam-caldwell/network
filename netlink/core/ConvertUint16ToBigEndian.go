package core

import "encoding/binary"

// ConvertUint16ToBigEndian - Flip bit order of little endian uint16 to big endian
func ConvertUint16ToBigEndian(i uint16) uint16 {
	if NativeEndian() == binary.BigEndian {
		return i
	}
	return (i&0xff00)>>8 |
		(i&0xff)<<8
}
