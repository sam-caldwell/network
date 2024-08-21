package core

import (
	"encoding/binary"
	"testing"
)

func TestConvertUin64ToBigEndian(t *testing.T) {
	tests := []struct {
		input    uint64
		expected uint64
	}{
		{0x0123456789ABCDEF, 0xEFCDAB8967452301}, // Little-endian to big-endian
		{0xAABBCCDDEEFF0011, 0x1100FFEEDDCCBBAA}, // Little-endian to big-endian
		{0x0000000000000001, 0x0100000000000000}, // Little-endian to big-endian
		{0xFF00000000000000, 0x00000000000000FF}, // Little-endian to big-endian
	}

	for _, tt := range tests {
		result := ConvertUin64ToBigEndian(tt.input)
		if NativeEndian() == binary.LittleEndian && result != tt.expected {
			t.Errorf("ConvertUin64ToBigEndian(%#016x) = %#016x; want %#016x", tt.input, result, tt.expected)
		} else if NativeEndian() == binary.BigEndian && result != tt.input {
			t.Errorf("ConvertUin64ToBigEndian(%#016x) = %#016x; want %#016x (no change expected for BigEndian)", tt.input, result, tt.input)
		}
	}
}
