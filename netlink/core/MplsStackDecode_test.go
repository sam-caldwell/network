package core

import (
	"testing"
)

func TestMplsStackDecode(t *testing.T) {
	tests := []struct {
		name      string
		input     []byte
		expected  []MplsLabels
		expectNil bool
	}{
		{
			name: "Valid MPLS Stack",
			input: []byte{
				0x00, 0x01, 0x23, 0x00, 0x00, 0x02, 0x34, 0x00,
				0x00, 0x03, 0x45, 0x00, 0x00, 0x04, 0x56, 0x00,
			},
			expected:  []MplsLabels{0x123, 0x234, 0x345, 0x456},
			expectNil: false,
		},
		{
			name:      "Non-multiple of 4 bytes",
			input:     []byte{0x00, 0x01, 0x23, 0x00, 0x00},
			expected:  nil,
			expectNil: true,
		},
		{
			name: "MPLS stack with bottom-of-stack (S bit set)",
			input: []byte{
				0x00, 0x01, 0x23, 0x00, 0x00, 0x02, 0x34, 0x00,
				0x00, 0x03, 0x45, 0x00, 0x00, 0x04, 0x56, 0x80,
			},
			expected:  []MplsLabels{0x123, 0x234, 0x345, 0x456},
			expectNil: false,
		},
	}

	equalMplsLabelStacks := func(a, b []MplsLabels) bool {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MplsStackDecode(tt.input)
			if tt.expectNil && result != nil {
				t.Errorf("Expected nil, got %v", result)
			} else if !tt.expectNil && !equalMplsLabelStacks(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
