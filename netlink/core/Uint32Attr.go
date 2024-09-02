//go:build linux

package core

// Uint32Attr - convert uint32 to []byte using system-native endianness
func Uint32Attr(v uint32) []byte {
	bytes := make([]byte, 4)
	nativeEndian.PutUint32(bytes, v)
	return bytes
}
