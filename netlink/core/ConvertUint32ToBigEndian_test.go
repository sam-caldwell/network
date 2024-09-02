package core

import (
	"encoding/binary"
	"testing"
)

func TestConvertUin32ToBigEndian(t *testing.T) {
	tests := []struct {
		input    uint32
		expected uint32
	}{
		{0x12345678, 0x78563412}, // 0x12345678 in little-endian should become 0x78563412 in big-endian
		{0xAABBCCDD, 0xDDCCBBAA}, // 0xAABBCCDD in little-endian should become 0xDDCCBBAA in big-endian
		{0x00000001, 0x01000000}, // 0x00000001 in little-endian should become 0x01000000 in big-endian
		{0xFF000000, 0x000000FF}, // 0xFF000000 in little-endian should become 0x000000FF in big-endian
	}

	for _, tt := range tests {
		result := ConvertUint32ToBigEndian(tt.input)
		if nativeEndian == binary.LittleEndian && result != tt.expected {
			t.Errorf("ConvertUin32ToBigEndian(%#08x) = %#08x; want %#08x", tt.input, result, tt.expected)
		} else if nativeEndian == binary.BigEndian && result != tt.input {
			t.Errorf("ConvertUin32ToBigEndian(%#08x) = %#08x; want %#08x (no change expected for BigEndian)", tt.input, result, tt.input)
		}
	}
}
