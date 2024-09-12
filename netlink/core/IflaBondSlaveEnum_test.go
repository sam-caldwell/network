package core

import (
	"testing"
	"unsafe"
)

// TestIflaBondSlaveEnum tests the size of the IflaBondSlaveEnum type and the values of its enumerated constants.
func TestIflaBondSlaveEnum(t *testing.T) {
	// Verify the size of the IflaBondSlaveEnum type is 1 byte (uint8)
	if size := unsafe.Sizeof(IflaBondSlaveUnspec); size != 1 {
		t.Errorf("Expected size of IflaBondSlaveEnum to be 1 byte, but got %d", size)
	}

	// Test the values of the constants
	tests := []struct {
		name  string
		value IflaBondSlaveEnum
		want  IflaBondSlaveEnum
	}{
		{"IflaBondSlaveUnspec", IflaBondSlaveUnspec, 0},
		{"IflaBondSlaveState", IflaBondSlaveState, 1},
		{"IflaBondSlaveMiiStatus", IflaBondSlaveMiiStatus, 2},
		{"IflaBondSlaveLinkFailureCount", IflaBondSlaveLinkFailureCount, 3},
		{"IflaBondSlavePermHwAddr", IflaBondSlavePermHwAddr, 4},
		{"IflaBondSlaveQueueId", IflaBondSlaveQueueId, 5},
		{"IflaBondSlaveAdAggregatorId", IflaBondSlaveAdAggregatorId, 6},
		{"IflaBondSlaveAdActorOperPortState", IflaBondSlaveAdActorOperPortState, 7},
		{"IflaBondSlaveAdPartnerOperPortState", IflaBondSlaveAdPartnerOperPortState, 8},
		{"IflaBondSlavePrio", IflaBondSlavePrio, 9},
		{"IflaBondSlaveMax", IflaBondSlaveMax, 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("Expected %s to have value %d, but got %d", tt.name, tt.want, tt.value)
			}
		})
	}
}
