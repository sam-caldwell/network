//go:build linux

package core

import "encoding/binary"

// BEUint32Attr - convert uint32 to big endian []byte slice
func BEUint32Attr(v uint32) []byte {

	bytes := make([]byte, 4)

	binary.BigEndian.PutUint32(bytes, v)

	return bytes

}
