package core

import (
	"testing"
	"unsafe"
)

// TestIfLaExtMaskEnum tests the size of the IfLaExtMask type and the values of its enumerated constants.
func TestIfLaExtMaskEnum(t *testing.T) {
	// Verify the size of the IfLaExtMask type is 1 byte (uint8)
	if size := unsafe.Sizeof(RtextFilterVf); size != 1 {
		t.Errorf("Expected size of IfLaExtMask to be 1 byte, but got %d", size)
	}

	// Test the values of the constants
	tests := []struct {
		name  string
		value IfLaExtMask
		want  IfLaExtMask
	}{
		{"RtextFilterVf", RtextFilterVf, 1},
		{"RtextFilterBrVlan", RtextFilterBrVlan, 2},
		{"RtextFilterBrVlanCompressed", RtextFilterBrVlanCompressed, 4},
		{"RtextFilterSkipStats", RtextFilterSkipStats, 8},
		{"RtextFilterMrp", RtextFilterMrp, 16},
		{"RtextFilterCfmConfig", RtextFilterCfmConfig, 32},
		{"RtextFilterCfmStatus", RtextFilterCfmStatus, 64},
		{"RtextFilterMst", RtextFilterMst, 128},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("Expected %s to have value %d, but got %d", tt.name, tt.want, tt.value)
			}
		})
	}
}
