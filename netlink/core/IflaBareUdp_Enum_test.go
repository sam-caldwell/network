package core

import (
	"testing"
	"unsafe"
)

// TestIflaBareUdpEnum tests the size of the IflaBareUdp type and the values of its enumerated constants.
func TestIflaBareUdpEnum(t *testing.T) {
	// Verify the size of the IflaBareUdp type is 1 byte (uint8)
	if size := unsafe.Sizeof(IflaBareUdpUnspec); size != 1 {
		t.Errorf("Expected size of IflaBareUdp to be 1 byte, but got %d", size)
	}

	// Test the values of the constants
	tests := []struct {
		name  string
		value IflaBareUdp
		want  IflaBareUdp
	}{
		{"IflaBareUdpUnspec", IflaBareUdpUnspec, 0},
		{"IflaBareUdpPort", IflaBareUdpPort, 1},
		{"IflaBareUdpEthertype", IflaBareUdpEthertype, 2},
		{"IflaBareUdpSrcportMin", IflaBareUdpSrcportMin, 3},
		{"IflaBareUdpMultiprotoMode", IflaBareUdpMultiprotoMode, 4},
		{"IflaBareUdpMax", IflaBareUdpMax, 4}, // Based on iota, the max value is expected to be 4
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("Expected %s to have value %d, but got %d", tt.name, tt.want, tt.value)
			}
		})
	}
}
