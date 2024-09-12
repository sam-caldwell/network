package core

import (
	"testing"
	"unsafe"
)

// TestFraEnum tests the size of the FraEnum type and the values of the enumerated constants.
func TestFraEnum(t *testing.T) {
	// Verify the size of the FraEnum type is 1 byte (uint8)
	if size := unsafe.Sizeof(FraUnspec); size != 1 {
		t.Errorf("Expected size of FraEnum to be 1 byte, got %d", size)
	}

	// Test the values of the constants
	tests := []struct {
		name  string
		value FraEnum
		want  FraEnum
	}{
		{"FraUnspec", FraUnspec, 0},
		{"FraDst", FraDst, 1},
		{"FraSrc", FraSrc, 2},
		{"FraIifName", FraIifName, 3},
		{"FraGoto", FraGoto, 4},
		{"FraUnused2", FraUnused2, 5},
		{"FraPriority", FraPriority, 6},
		{"FraUnused3", FraUnused3, 7},
		{"FraUnused4", FraUnused4, 8},
		{"FraUnused5", FraUnused5, 9},
		{"FraFwMark", FraFwMark, 10},
		{"FraFlow", FraFlow, 11},
		{"FraTunId", FraTunId, 12},
		{"FraSuppressIfGroup", FraSuppressIfGroup, 13},
		{"FraSuppressPrefixLen", FraSuppressPrefixLen, 14},
		{"FraTable", FraTable, 15},
		{"FraFwMask", FraFwMask, 16},
		{"FraOifName", FraOifName, 17},
		{"FraPad", FraPad, 18},
		{"FraL3Mdev", FraL3Mdev, 19},
		{"FraUidRange", FraUidRange, 20},
		{"FraProtocol", FraProtocol, 21},
		{"FraIpProto", FraIpProto, 22},
		{"FraSportRange", FraSportRange, 23},
		{"FraDportRange", FraDportRange, 24},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("Expected %s to have value %d, but got %d", tt.name, tt.want, tt.value)
			}
		})
	}
}
