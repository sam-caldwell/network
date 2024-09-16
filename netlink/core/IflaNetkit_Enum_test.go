package core

import (
	"testing"
	"unsafe"
)

// TestIflaNetkitEnum tests the size of the IflaNetkit type and the values of its enumerated constants.
func TestIflaNetkitEnum(t *testing.T) {
	// Verify that the size of the IflaNetkit type is 1 byte (uint8)
	if size := unsafe.Sizeof(IflaNetkitUnspec); size != 1 {
		t.Errorf("Expected size of IflaNetkit to be 1 byte, but got %d", size)
	}

	// Test the values of the constants
	tests := []struct {
		name  string
		value IflaNetkit
		want  IflaNetkit
	}{
		{"IflaNetkitUnspec", IflaNetkitUnspec, 0},
		{"IflaNetkitPeerInfo", IflaNetkitPeerInfo, 1},
		{"IflaNetkitPrimary", IflaNetkitPrimary, 2},
		{"IflaNetkitPolicy", IflaNetkitPolicy, 3},
		{"IflaNetkitPeerPolicy", IflaNetkitPeerPolicy, 4},
		{"IflaNetkitMode", IflaNetkitMode, 5},
		{"IflaNetkitMax", IflaNetkitMax, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("Expected %s to have value %d, but got %d", tt.name, tt.want, tt.value)
			}
		})
	}
}
