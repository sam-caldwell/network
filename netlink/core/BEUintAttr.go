package core

import "encoding/binary"

// BEUintAttr converts uint16, uint32, or uint64 values into a big-endian byte slice.
func BEUintAttr[T uint16 | uint32 | uint64](v T) (bytes []byte) {
	switch any(v).(type) {
	case uint16:
		bytes = make([]byte, 2)
		binary.BigEndian.PutUint16(bytes, uint16(v))
	case uint32:
		bytes = make([]byte, 4)
		binary.BigEndian.PutUint32(bytes, uint32(v))
	case uint64:
		bytes = make([]byte, 8)
		binary.BigEndian.PutUint64(bytes, uint64(v))
	}
	return bytes
}
