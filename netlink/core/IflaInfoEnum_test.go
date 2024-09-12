package core

import (
	"testing"
	"unsafe"
)

// TestIflaInfoEnum tests the size of the IflaInfoEnum type and the values of its enumerated constants.
func TestIflaInfoEnum(t *testing.T) {
	// Verify the size of the IflaInfoEnum type is 1 byte (uint8)
	if size := unsafe.Sizeof(IflaInfoUnspec); size != 1 {
		t.Errorf("Expected size of IflaInfoEnum to be 1 byte, but got %d", size)
	}

	// Test the values of the constants
	tests := []struct {
		name  string
		value IflaInfoEnum
		want  IflaInfoEnum
	}{
		{"IflaInfoUnspec", IflaInfoUnspec, 0},
		{"IflaInfoKind", IflaInfoKind, 1},
		{"IflaInfoData", IflaInfoData, 2},
		{"IflaInfoXstats", IflaInfoXstats, 3},
		{"IflaInfoSlaveKind", IflaInfoSlaveKind, 4},
		{"IflaInfoSlaveData", IflaInfoSlaveData, 5},
		{"IflaInfoMax", IflaInfoMax, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("Expected %s to have value %d, but got %d", tt.name, tt.want, tt.value)
			}
		})
	}
}
