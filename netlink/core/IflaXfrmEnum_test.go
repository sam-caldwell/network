package core

import (
	"testing"
)

// TestIflaXfrmEnum tests the values of IflaXfrmEnum constants.
func TestIflaXfrmEnum(t *testing.T) {
	tests := []struct {
		name     string
		value    IflaXfrmEnum
		expected IflaXfrmEnum
	}{
		{"IflaXfrmUnspec", IflaXfrmUnspec, 0},
		{"IflaXfrmLink", IflaXfrmLink, 1},
		{"IflaXfrmIfId", IflaXfrmIfId, 2},
		{"IflaXfrmMax", IflaXfrmMax, 2}, // IflaXfrmMax should be the same as IflaXfrmIfId.
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.value != test.expected {
				t.Errorf("Expected %d but got %d for %s", test.expected, test.value, test.name)
			}
		})
	}
}
