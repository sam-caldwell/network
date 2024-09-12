package core

import (
	"testing"
	"unsafe"
)

// TestIflaMacVlanEnum tests the size of the IflaMacVlanEnum type and the values of its enumerated constants.
func TestIflaMacVlanEnum(t *testing.T) {
	// Verify that the size of the IflaMacVlanEnum type is 1 byte (uint8)
	if size := unsafe.Sizeof(IflaMacvlanUnspec); size != 1 {
		t.Errorf("Expected size of IflaMacVlanEnum to be 1 byte, but got %d", size)
	}

	// Test the values of the constants
	tests := []struct {
		name  string
		value IflaMacVlanEnum
		want  IflaMacVlanEnum
	}{
		{"IflaMacvlanUnspec", IflaMacvlanUnspec, 0},
		{"IflaMacvlanMode", IflaMacvlanMode, 1},
		{"IflaMacvlanFlags", IflaMacvlanFlags, 2},
		{"IflaMacvlanMacaddrMode", IflaMacvlanMacaddrMode, 3},
		{"IflaMacvlanMacaddr", IflaMacvlanMacaddr, 4},
		{"IflaMacvlanMacaddrData", IflaMacvlanMacaddrData, 5},
		{"IflaMacvlanMacaddrCount", IflaMacvlanMacaddrCount, 6},
		{"IflaMacvlanBcQueueLen", IflaMacvlanBcQueueLen, 7},
		{"IflaMacvlanBcQueueLenUsed", IflaMacvlanBcQueueLenUsed, 8},
		{"IflaMacvlanBcCutoff", IflaMacvlanBcCutoff, 9},
		{"IflaMacvlanMax", IflaMacvlanMax, 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("Expected %s to have value %d, but got %d", tt.name, tt.want, tt.value)
			}
		})
	}
}
