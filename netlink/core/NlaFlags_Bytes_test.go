package core

import (
	"testing"
)

func TestNlaFlagsBytes(t *testing.T) {
	tests := []struct {
		name     string
		input    NlaFlags
		expected []byte
	}{
		{
			name:     "Zero value",
			input:    NlaFlags(0),
			expected: []byte{0x00, 0x00},
		},
		{
			name:     "Small value",
			input:    NlaFlags(1),
			expected: []byte{0x01, 0x00},
		},
		{
			name:     "Mid value",
			input:    NlaFlags(256),
			expected: []byte{0x00, 0x01},
		},
		{
			name:     "Max value",
			input:    NlaFlags(65535),
			expected: []byte{0xFF, 0xFF},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.Bytes()
			if len(result) != 2 || result[0] != tt.expected[0] || result[1] != tt.expected[1] {
				t.Errorf("NlaFlags(%d).Bytes() = %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}
