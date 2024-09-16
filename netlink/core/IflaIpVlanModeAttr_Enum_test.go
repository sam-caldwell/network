package core

import (
	"testing"
	"unsafe"
)

// TestIflaIpVlanModeEnum tests the size of the IflaIpVlanModeAttr type and the values of its enumerated constants.
func TestIflaIpVlanModeEnum(t *testing.T) {
	// Verify that the size of the IflaIpVlanModeAttr type is 1 byte (uint8)
	if size := unsafe.Sizeof(IflaIpVlanUnspec); size != 1 {
		t.Errorf("Expected size of IflaIpVlanModeAttr to be 1 byte, but got %d", size)
	}

	// Test the values of the constants
	tests := []struct {
		name  string
		value IflaIpVlanModeAttr
		want  IflaIpVlanModeAttr
	}{
		{"IflaIpVlanUnspec", IflaIpVlanUnspec, 0},
		{"IflaIpVlanMode", IflaIpVlanMode, 1},
		{"IflaIpVlanFlags", IflaIpVlanFlags, 2},
		{"IflaIpVlanMax", IflaIpVlanMax, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("Expected %s to have value %d, but got %d", tt.name, tt.want, tt.value)
			}
		})
	}
}
