//go:build linux

package core

import (
	"testing"
)

func TestConvertGoStringToNullTerminated(t *testing.T) {
	tests := []struct {
		input    string
		expected []byte
	}{
		{
			input:    "test",
			expected: []byte{'t', 'e', 's', 't', 0},
		},
		{
			input:    "",
			expected: []byte{0},
		},
		{
			input:    "hello",
			expected: []byte{'h', 'e', 'l', 'l', 'o', 0},
		},
	}

	for _, tt := range tests {
		result := ConvertGoStringToNullTerminated(tt.input)
		if len(result) != len(tt.expected) {
			t.Errorf("TestConvertGoStringToNullTerminated(%q): expected length %d, got %d", tt.input, len(tt.expected), len(result))
		}
		for i := range result {
			if result[i] != tt.expected[i] {
				t.Errorf("TestConvertGoStringToNullTerminated(%q): expected byte %v at index %d, got %v", tt.input, tt.expected[i], i, result[i])
			}
		}
	}
}
