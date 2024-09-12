package core

import (
	"testing"
	"unsafe"
)

// TestIfLaVfVlanInfoEnum tests the size of the IfLaVfVlanInfoEnum type and the values of its enumerated constants.
func TestIfLaVfVlanInfoEnum(t *testing.T) {
	// Verify that the size of the IfLaVfVlanInfoEnum type is 2 bytes (uint16)
	if size := unsafe.Sizeof(IfLaVfVlanInfoUnspec); size != 2 {
		t.Errorf("Expected size of IfLaVfVlanInfoEnum to be 2 bytes, but got %d", size)
	}

	// Test the values of the constants
	tests := []struct {
		name  string
		value IfLaVfVlanInfoEnum
		want  IfLaVfVlanInfoEnum
	}{
		{"IfLaVfVlanInfoUnspec", IfLaVfVlanInfoUnspec, 0},
		{"IfLaVfVlanInfo", IfLaVfVlanInfo, 1},
		{"IfLaVfVlanInfoMax", IfLaVfVlanInfoMax, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("Expected %s to have value %d, but got %d", tt.name, tt.want, tt.value)
			}
		})
	}
}
