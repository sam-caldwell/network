package core

import (
	"testing"
	"unsafe"
)

// TestIflaVfEnum tests the size of the IflaVfEnum type and the values of its enumerated constants.
func TestIflaVfEnum(t *testing.T) {
	// Verify that the size of the IflaVfEnum type is 1 byte (uint8)
	if size := unsafe.Sizeof(IflaVfUnspec); size != 1 {
		t.Errorf("Expected size of IflaVfEnum to be 1 byte, but got %d", size)
	}

	// Test the values of the constants
	tests := []struct {
		name  string
		value IflaVfEnum
		want  IflaVfEnum
	}{
		{"IflaVfUnspec", IflaVfUnspec, 0},
		{"IflaVfMac", IflaVfMac, 1},
		{"IflaVfVlan", IflaVfVlan, 2},
		{"IflaVfTxRate", IflaVfTxRate, 3},
		{"IflaVfSpoofchk", IflaVfSpoofchk, 4},
		{"IflaVfLinkState", IflaVfLinkState, 5},
		{"IflaVfRate", IflaVfRate, 6},
		{"IflaVfRssQueryEn", IflaVfRssQueryEn, 7},
		{"IflaVfStats", IflaVfStats, 8},
		{"IflaVfTrust", IflaVfTrust, 9},
		{"IflaVfIbNodeGuid", IflaVfIbNodeGuid, 10},
		{"IflaVfIbPortGuid", IflaVfIbPortGuid, 11},
		{"IflaVfVlanList", IflaVfVlanList, 12},
		{"IflaVfBroadcast", IflaVfBroadcast, 13},
		{"IflaVfMax", IflaVfMax, 13},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("Expected %s to have value %d, but got %d", tt.name, tt.want, tt.value)
			}
		})
	}
}
