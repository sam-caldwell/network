package core

import (
	"testing"
)

// TestIflaXdpEnum tests the values of IflaXdpEnum constants.
func TestIflaXdpEnum(t *testing.T) {
	tests := []struct {
		name     string
		value    IflaXdpEnum
		expected IflaXdpEnum
	}{
		{"IflaXdpUnspec", IflaXdpUnspec, 0},
		{"IflaXdpFd", IflaXdpFd, 1},
		{"IflaXdpAttached", IflaXdpAttached, 2},
		{"IflaXdpFlags", IflaXdpFlags, 3},
		{"IflaXdpProgId", IflaXdpProgId, 4},
		{"IflaXdpDrvProgId", IflaXdpDrvProgId, 5},
		{"IflaXdpSkbProgId", IflaXdpSkbProgId, 6},
		{"IflaXdpHwProgId", IflaXdpHwProgId, 7},
		{"IflaXdpExpectedFd", IflaXdpExpectedFd, 8},
		{"IflaXdpMax", IflaXdpMax, 8}, // IflaXdpMax is expected to be the highest value, same as IflaXdpExpectedFd.
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.value != test.expected {
				t.Errorf("Expected %d but got %d for %s", test.expected, test.value, test.name)
			}
		})
	}
}
