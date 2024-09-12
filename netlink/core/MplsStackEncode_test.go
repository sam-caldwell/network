package core

import (
	"fmt"
	"testing"
)

func TestMplsStackEncode(t *testing.T) {
	tests := []struct {
		name     string
		input    []uint32
		expected []byte
	}{
		{
			name:  "Valid MPLS Stack Encoding",
			input: []uint32{0x000, 0x123, 0x234, 0x345, 0x355},
			expected: []byte{
				//|<-----label (20 bit)----->|TCS|<---TTL--->|
				//|             20           |3+1|           |
				0b00000000, 0b00000000, 0b00000000, 0b00000000, // Label: 0x000, TC: 000, S: 0, TTL: 0x00
				0b00000000, 0b00010010, 0b00110000, 0b00000000, // Label: 0x123, TC: 000, S: 0, TTL: 0x00
				0b00000000, 0b00100011, 0b01000000, 0b00000000, // Label: 0x234, TC: 000, S: 0, TTL: 0x00
				0b00000000, 0b00110100, 0b01010000, 0b00000000, // Label: 0x345, TC: 000, S: 0, TTL: 0x00
				0b00000000, 0b00110101, 0b01010001, 0b00000000, // Label: 0x355, TC: 000, S: 1, TTL: 0x00
			},
		},
		{
			name:     "Empty MPLS Stack",
			input:    []uint32{},
			expected: []byte{},
		},
	}

	equalByteSlices := func(a, b []byte) bool {
		if len(a) != len(b) {
			return false
		}
		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}

	binaryPrint := func(data []byte) string {
		s := ""
		for i, b := range data {
			s += fmt.Sprintf("%08b ", b)
			if (i+1)%4 == 0 {
				s += "\n"
			}
		}
		return s
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MplsStackEncode(tt.input...)
			if !equalByteSlices(result, tt.expected) {
				t.Fatalf("MplsStackEncode() failure\n"+
					"expected: %s\n"+
					"  result: %s", binaryPrint(tt.expected), binaryPrint(result))
			}
		})
	}
}
