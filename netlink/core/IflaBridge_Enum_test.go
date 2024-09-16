package core

import (
	"testing"
	"unsafe"
)

// TestIflaBridgeEnum tests the size of the IflaBridge type and the values of its enumerated constants.
func TestIflaBridgeEnum(t *testing.T) {
	// Verify the size of the IflaBridge type is 8 bytes (int)
	if size := unsafe.Sizeof(IflaBridgeFlags); size != unsafe.Sizeof(int(0)) {
		t.Errorf("Expected size of IflaBridge to be 4 bytes, but got %d", size)
	}

	// Test the values of the constants
	tests := []struct {
		name  string
		value IflaBridge
		want  IflaBridge
	}{
		{"IflaBridgeFlags", IflaBridgeFlags, 0},
		{"IflaBridgeMode", IflaBridgeMode, 1},
		{"IflaBridgeVlanInfo", IflaBridgeVlanInfo, 2},
		{"IflaBridgeVlanTunnelInfo", IflaBridgeVlanTunnelInfo, 3},
		{"IflaBridgeMrp", IflaBridgeMrp, 4},
		{"IflaBridgeCfm", IflaBridgeCfm, 5},
		{"IflaBridgeMst", IflaBridgeMst, 6},
		{"IflaBridgeMax", IflaBridgeMax, 6}, // Maximum value equals IflaBridgeMst
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("Expected %s to have value %d, but got %d", tt.name, tt.want, tt.value)
			}
		})
	}
}
