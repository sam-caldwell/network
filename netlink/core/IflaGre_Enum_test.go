package core

import (
	"testing"
	"unsafe"
)

// TestIflaGreEnum tests the size of the IflaGre type and the values of its enumerated constants.
func TestIflaGreEnum(t *testing.T) {
	// Verify the size of the IflaGre type is 1 byte (uint8)
	if size := unsafe.Sizeof(IflaGreUnspec); size != 1 {
		t.Errorf("Expected size of IflaGre to be 1 byte, but got %d", size)
	}

	// Test the values of the constants
	tests := []struct {
		name  string
		value IflaGre
		want  IflaGre
	}{
		{"IflaGreUnspec", IflaGreUnspec, 0},
		{"IflaGreLink", IflaGreLink, 1},
		{"IflaGreIflags", IflaGreIflags, 2},
		{"IflaGreOflags", IflaGreOflags, 3},
		{"IflaGreIkey", IflaGreIkey, 4},
		{"IflaGreOkey", IflaGreOkey, 5},
		{"IflaGreLocal", IflaGreLocal, 6},
		{"IflaGreRemote", IflaGreRemote, 7},
		{"IflaGreTtl", IflaGreTtl, 8},
		{"IflaGreTos", IflaGreTos, 9},
		{"IflaGrePmtudisc", IflaGrePmtudisc, 10},
		{"IflaGreEncapLimit", IflaGreEncapLimit, 11},
		{"IflaGreFlowinfo", IflaGreFlowinfo, 12},
		{"IflaGreFlags", IflaGreFlags, 13},
		{"IflaGreEncapType", IflaGreEncapType, 14},
		{"IflaGreEncapFlags", IflaGreEncapFlags, 15},
		{"IflaGreEncapSport", IflaGreEncapSport, 16},
		{"IflaGreEncapDport", IflaGreEncapDport, 17},
		{"IflaGreCollectMetadata", IflaGreCollectMetadata, 18},
		{"IflaGreMax", IflaGreMax, IflaGreCollectMetadata},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("Expected %s to have value %d, but got %d", tt.name, tt.want, tt.value)
			}
		})
	}
}
