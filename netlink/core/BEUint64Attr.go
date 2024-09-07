package core

import "encoding/binary"

// BEUint64Attr - convert uint64 to big endian []byte slice
func BEUint64Attr(v uint64) []byte {

	bytes := make([]byte, 8)

	binary.BigEndian.PutUint64(bytes, v)

	return bytes

}
