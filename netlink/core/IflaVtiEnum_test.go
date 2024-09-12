package core

import (
	"testing"
)

// TestIflaVtiEnum tests the values of IflaVtiEnum constants.
func TestIflaVtiEnum(t *testing.T) {
	tests := []struct {
		name     string
		value    IflaVtiEnum
		expected IflaVtiEnum
	}{
		{"IflaVtiUnspec", IflaVtiUnspec, 0},
		{"IflaVtiLink", IflaVtiLink, 1},
		{"IflaVtiIkey", IflaVtiIkey, 2},
		{"IflaVtiOkey", IflaVtiOkey, 3},
		{"IflaVtiLocal", IflaVtiLocal, 4},
		{"IflaVtiRemote", IflaVtiRemote, 5},
		{"IflaVtiMax", IflaVtiMax, 5}, // Max value should be the same as IflaVtiRemote, as per the iota - 1 logic
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.value != test.expected {
				t.Errorf("Expected %d but got %d for %s", test.expected, test.value, test.name)
			}
		})
	}
}
