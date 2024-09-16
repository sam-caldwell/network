package core

import (
	"testing"
	"unsafe"
)

// TestIflaIpVlanEnum tests the size of the IflaIpVlan type and the values of its enumerated constants.
func TestIflaIpVlanEnum(t *testing.T) {
	// Verify that the size of the IflaIpVlan type is 1 byte (uint8)
	if size := unsafe.Sizeof(IflaVfInfoUnspec); size != 1 {
		t.Errorf("Expected size of IflaIpVlan to be 1 byte, but got %d", size)
	}

	// Test the values of the constants
	tests := []struct {
		name  string
		value IflaIpVlan
		want  IflaIpVlan
	}{
		{"IflaVfInfoUnspec", IflaVfInfoUnspec, 0},
		{"IflaVfInfo", IflaVfInfo, 1},
		{"IflaVfInfoMax", IflaVfInfoMax, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("Expected %s to have value %d, but got %d", tt.name, tt.want, tt.value)
			}
		})
	}
}
