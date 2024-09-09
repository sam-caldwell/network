package core

import (
	"testing"
)

func TestRtaAlignOf(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{name: "alignment for 1", input: 1, expected: 4},
		{name: "alignment for 2", input: 2, expected: 4},
		{name: "alignment for 3", input: 3, expected: 4},
		{name: "alignment for 4", input: 4, expected: 4},
		{name: "alignment for 5", input: 5, expected: 8},
		{name: "alignment for 7", input: 7, expected: 8},
		{name: "alignment for 8", input: 8, expected: 8},
		{name: "alignment for 9", input: 9, expected: 12},
		{name: "alignment for 15", input: 15, expected: 16},
		{name: "alignment for 16", input: 16, expected: 16},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := rtaAlignOf(tt.input)
			if result != tt.expected {
				t.Errorf("rtaAlignOf(%d): expected %d, got %d", tt.input, tt.expected, result)
			}
		})
	}
}
