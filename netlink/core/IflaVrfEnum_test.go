package core

import (
	"testing"
)

// TestIflaVrfEnum tests the values of IflaVrfEnum constants.
func TestIflaVrfEnum(t *testing.T) {
	tests := []struct {
		name     string
		value    IflaVrfEnum
		expected IflaVrfEnum
	}{
		{"IflaVrfUnspec", IflaVrfUnspec, 0},
		{"IflaVrfTable", IflaVrfTable, 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.value != test.expected {
				t.Errorf("Expected %d but got %d", test.expected, test.value)
			}
		})
	}
}
