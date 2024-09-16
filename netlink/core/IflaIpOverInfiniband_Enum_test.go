package core

import (
	"testing"
	"unsafe"
)

// TestIflaIpOverInfinibandEnum tests the size of the IflaIpOverInfiniband type and the values of its enumerated constants.
func TestIflaIpOverInfinibandEnum(t *testing.T) {
	// Verify the size of the IflaIpOverInfiniband type is 1 byte (uint8)
	if size := unsafe.Sizeof(IflaIpOverInfinibandUnspec); size != 1 {
		t.Errorf("Expected size of IflaIpOverInfiniband to be 1 byte, but got %d", size)
	}

	// Test the values of the constants
	tests := []struct {
		name  string
		value IflaIpOverInfiniband
		want  IflaIpOverInfiniband
	}{
		{"IflaIpOverInfinibandUnspec", IflaIpOverInfinibandUnspec, 0},
		{"IflaIpOverInfinibandPkey", IflaIpOverInfinibandPkey, 1},
		{"IflaIpOverInfinibandMode", IflaIpOverInfinibandMode, 2},
		{"IflaIpOverInfinibandUmcast", IflaIpOverInfinibandUmcast, 3},
		{"IflaIpOverInfinibandMax", IflaIpOverInfinibandMax, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("Expected %s to have value %d, but got %d", tt.name, tt.want, tt.value)
			}
		})
	}
}
