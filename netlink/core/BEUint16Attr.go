package core

import "encoding/binary"

// BEUint16Attr - convert uint16 to big endian []byte slice
func BEUint16Attr(v uint16) []byte {

	bytes := make([]byte, 2)

	binary.BigEndian.PutUint16(bytes, v)

	return bytes

}
