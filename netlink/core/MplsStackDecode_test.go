package core

import (
	"testing"
)

func TestMplsStackDecode(t *testing.T) {
	tests := []struct {
		name      string
		input     []byte
		expected  []uint32
		expectNil bool
	}{
		{
			//row:0
			name: "Valid MPLS Stack",
			// Each label is encoded into 4 bytes as:
			// {0x001230FF, 0x002340FF, 0x003450FF, 0x004560FF}
			// 20-bit labels: 0x123, 0x234, 0x345, 0x456 (last label has the Bottom of Stack bit set)
			input: []byte{
				//|<-----label (20 bit)----->|TCS|<---TTL--->|
				//|             20           |3+1|           |
				0b00000000, 0b00000000, 0b00000000, 0b11111111, // Label: 0x000, TC: 000, S: 0, TTL: 0xFF
				0b00000000, 0b00010010, 0b00110000, 0b11111111, // Label: 0x123, TC: 000, S: 0, TTL: 0xFF
				0b00000000, 0b00100011, 0b01000000, 0b11111111, // Label: 0x234, TC: 000, S: 0, TTL: 0xFF
				0b00000000, 0b00110100, 0b01010000, 0b11111111, // Label: 0x345, TC: 000, S: 0, TTL: 0xFF
				0b00000000, 0b00110101, 0b01011111, 0b11111111, // Label: 0x355, TC: FFF, S: 1, TTL: 0xFF
			},
			expected:  []uint32{0x000, 0x123, 0x234, 0x345, 0x355},
			expectNil: false,
		},
		{
			//row:1
			name:      "Non-multiple of 4 bytes",
			input:     []byte{0x00, 0x01, 0x23, 0xFF, 0x00},
			expected:  nil,
			expectNil: true,
		},
	}

	equalMplsLabelStacks := func(a, b []uint32) bool {
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

	for row, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MplsStackDecode(tt.input)
			if tt.expectNil && result != nil {
				t.Errorf("(row:%d) Expected nil, got %v", row, result)
			} else if !tt.expectNil && !equalMplsLabelStacks(result, tt.expected) {
				t.Errorf("(row:%d) Expected %v, got %v", row, tt.expected, result)
			}
		})
	}
}
