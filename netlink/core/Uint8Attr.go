//go:build linux

package core

// Uint8Attr - convert uint8 to []byte using system-native endianness
func Uint8Attr(b uint8) []byte {

	return []byte{byte(b)}

}
