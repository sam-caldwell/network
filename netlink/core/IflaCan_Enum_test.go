package core

import (
	"testing"
	"unsafe"
)

// TestIflaCanEnum tests the size of the IflaCan_Enum type and the values of its enumerated constants.
func TestIflaCanEnum(t *testing.T) {
	// Verify the size of the IflaCan_Enum type is 1 byte (uint8)
	if size := unsafe.Sizeof(IflaCanUnspec); size != 1 {
		t.Errorf("Expected size of IflaCan_Enum to be 1 byte, but got %d", size)
	}

	// Test the values of the constants
	tests := []struct {
		name  string
		value IflaCan_Enum
		want  IflaCan_Enum
	}{
		{"IflaCanUnspec", IflaCanUnspec, 0},
		{"IflaCanBittiming", IflaCanBittiming, 1},
		{"IflaCanBittimingConst", IflaCanBittimingConst, 2},
		{"IflaCanClock", IflaCanClock, 3},
		{"IflaCanState", IflaCanState, 4},
		{"IflaCanCtrlmode", IflaCanCtrlmode, 5},
		{"IflaCanRestartMs", IflaCanRestartMs, 6},
		{"IflaCanRestart", IflaCanRestart, 7},
		{"IflaCanBerrCounter", IflaCanBerrCounter, 8},
		{"IflaCanDataBittiming", IflaCanDataBittiming, 9},
		{"IflaCanDataBittimingConst", IflaCanDataBittimingConst, 10},
		{"IflaCanTermination", IflaCanTermination, 11},
		{"IflaCanTerminationConst", IflaCanTerminationConst, 12},
		{"IflaCanBitrateConst", IflaCanBitrateConst, 13},
		{"IflaCanDataBitrateConst", IflaCanDataBitrateConst, 14},
		{"IflaCanBitrateMax", IflaCanBitrateMax, 15},
		{"IflaCanMax", IflaCanMax, 15}, // Max equals the last value
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("Expected %s to have value %d, but got %d", tt.name, tt.want, tt.value)
			}
		})
	}
}
