package core

import (
	"testing"
	"unsafe"
)

// TestIflaTunEnum tests the size of the IflaTunEnum type and the values of its enumerated constants.
func TestIflaTunEnum(t *testing.T) {
	// Verify that the size of the IflaTunEnum type is 1 byte (uint8)
	if size := unsafe.Sizeof(IflaTunUnspec); size != 1 {
		t.Errorf("Expected size of IflaTunEnum to be 1 byte, but got %d", size)
	}

	// Test the values of the constants
	tests := []struct {
		name  string
		value IflaTunEnum
		want  IflaTunEnum
	}{
		{"IflaTunUnspec", IflaTunUnspec, 0},
		{"IflaTunOwner", IflaTunOwner, 1},
		{"IflaTunGroup", IflaTunGroup, 2},
		{"IflaTunType", IflaTunType, 3},
		{"IflaTunPi", IflaTunPi, 4},
		{"IflaTunVnetHdr", IflaTunVnetHdr, 5},
		{"IflaTunPersist", IflaTunPersist, 6},
		{"IflaTunMultiQueue", IflaTunMultiQueue, 7},
		{"IflaTunNumQueues", IflaTunNumQueues, 8},
		{"IflaTunNumDisabledQueues", IflaTunNumDisabledQueues, 9},
		{"IflaTunMax", IflaTunMax, 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("Expected %s to have value %d, but got %d", tt.name, tt.want, tt.value)
			}
		})
	}
}
