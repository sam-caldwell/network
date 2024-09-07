//go:build linux

package core

import (
	"testing"
)

func TestBytesToString(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected string
	}{
		{
			name:     "NormalStringWithNullTerminator",
			input:    []byte{'h', 'e', 'l', 'l', 'o', 0, 'w', 'o', 'r', 'l', 'd'},
			expected: "hello",
		},
		{
			name:     "StringWithoutNullTerminator",
			input:    []byte{'h', 'e', 'l', 'l', 'o'},
			expected: "hello",
		},
		{
			name:     "EmptyByteSlice",
			input:    []byte{},
			expected: "",
		},
		{
			name:     "OnlyNullByte",
			input:    []byte{0},
			expected: "",
		},
		{
			name:     "NullByteAtBeginning",
			input:    []byte{0, 'h', 'e', 'l', 'l', 'o'},
			expected: "",
		},
		{
			name:     "NullByteInMiddle",
			input:    []byte{'h', 'i', 0, 't', 'h', 'e', 'r', 'e'},
			expected: "hi",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BytesToString(tt.input)
			if result != tt.expected {
				t.Errorf("BytesToString() = %v, want %v", result, tt.expected)
			}
		})
	}
}
