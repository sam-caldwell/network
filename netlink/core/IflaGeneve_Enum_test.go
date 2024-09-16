package core

import (
	"testing"
	"unsafe"
)

// TestIflaGeneveEnum tests the size of the IflaGeneve type and the values of its enumerated constants.
func TestIflaGeneveEnum(t *testing.T) {
	// Verify the size of the IflaGeneve type is 1 byte (uint8)
	if size := unsafe.Sizeof(IflaGeneveUnspec); size != 1 {
		t.Errorf("Expected size of IflaGeneve to be 1 byte, but got %d", size)
	}

	// Test the values of the constants
	tests := []struct {
		name  string
		value IflaGeneve
		want  IflaGeneve
	}{
		{"IflaGeneveUnspec", IflaGeneveUnspec, 0},
		{"IflaGeneveId", IflaGeneveId, 1},
		{"IflaGeneveRemote", IflaGeneveRemote, 2},
		{"IflaGeneveRemote6", IflaGeneveRemote6, 3},
		{"IflaGeneveTtl", IflaGeneveTtl, 4},
		{"IflaGeneveTos", IflaGeneveTos, 5},
		{"IflaGeneveLabel", IflaGeneveLabel, 6},
		{"IflaGenevePort", IflaGenevePort, 7},
		{"IflaGeneveCollectMetadata", IflaGeneveCollectMetadata, 8},
		{"IflaGeneveUdpCsum", IflaGeneveUdpCsum, 9},
		{"IflaGeneveUdpZeroCsum6Tx", IflaGeneveUdpZeroCsum6Tx, 10},
		{"IflaGeneveUdpZeroCsum6Rx", IflaGeneveUdpZeroCsum6Rx, 11},
		{"IflaGeneveTtlInherit", IflaGeneveTtlInherit, 12},
		{"IflaGeneveDf", IflaGeneveDf, 13},
		{"IflaGeneveInnerProtoInherit", IflaGeneveInnerProtoInherit, 14},
		{"IflaGeneveMax", IflaGeneveMax, 14}, // Since it equals the last constant, IflaGeneveMax should match IflaGeneveInnerProtoInherit
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("Expected %s to have value %d, but got %d", tt.name, tt.want, tt.value)
			}
		})
	}
}
