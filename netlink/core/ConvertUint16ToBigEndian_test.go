package core

import (
	"encoding/binary"
	"testing"
)

func TestConvertUint16ToBigEndian(t *testing.T) {
	tests := []struct {
		input    uint16
		expected uint16
	}{
		{0x1234, 0x3412}, // 0x1234 in little-endian should become 0x3412 in big-endian
		{0xABCD, 0xCDAB}, // 0xABCD in little-endian should become 0xCDAB in big-endian
		{0x0001, 0x0100}, // 0x0001 in little-endian should become 0x0100 in big-endian
		{0xFF00, 0x00FF}, // 0xFF00 in little-endian should become 0x00FF in big-endian
	}

	for _, tt := range tests {
		result := ConvertUint16ToBigEndian(tt.input)
		if NativeEndian == binary.LittleEndian && result != tt.expected {
			t.Errorf("ConvertUint16ToBigEndian(%#04x) = %#04x; want %#04x", tt.input, result, tt.expected)
		} else if NativeEndian == binary.BigEndian && result != tt.input {
			t.Errorf("ConvertUint16ToBigEndian(%#04x) = %#04x; want %#04x (no change expected for BigEndian)", tt.input, result, tt.input)
		}
	}
}
