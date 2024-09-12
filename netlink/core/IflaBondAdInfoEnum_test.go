package core

import (
	"testing"
	"unsafe"
)

// TestIflaBondAdInfoEnum tests the size of the IflaBondAdInfoEnum type and the values of its enumerated constants.
func TestIflaBondAdInfoEnum(t *testing.T) {
	// Verify the size of the IflaBondAdInfoEnum type is 1 byte (uint8)
	if size := unsafe.Sizeof(IflaBondAdInfoUnspec); size != 1 {
		t.Errorf("Expected size of IflaBondAdInfoEnum to be 1 byte, but got %d", size)
	}

	// Test the values of the constants
	tests := []struct {
		name  string
		value IflaBondAdInfoEnum
		want  IflaBondAdInfoEnum
	}{
		{"IflaBondAdInfoUnspec", IflaBondAdInfoUnspec, 0},
		{"IflaBondAdInfoAggregator", IflaBondAdInfoAggregator, 1},
		{"IflaBondAdInfoNumPorts", IflaBondAdInfoNumPorts, 2},
		{"IflaBondAdInfoActorKey", IflaBondAdInfoActorKey, 3},
		{"IflaBondAdInfoPartnerKey", IflaBondAdInfoPartnerKey, 4},
		{"IflaBondAdInfoPartnerMac", IflaBondAdInfoPartnerMac, 5},
		{"IflaBondAdInfoMax", IflaBondAdInfoMax, 5}, // Since it's iota - 1
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("Expected %s to have value %d, but got %d", tt.name, tt.want, tt.value)
			}
		})
	}
}
