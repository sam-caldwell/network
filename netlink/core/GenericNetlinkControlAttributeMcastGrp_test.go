package core

import (
	"testing"
	"unsafe"
)

// TestGenericNetlinkControlAttributeMcastGrp tests the size of the GenericNetlinkControlAttributeMcastGrp type
// and the values of its enumerated constants.
func TestGenericNetlinkControlAttributeMcastGrp(t *testing.T) {
	// Verify the size of the GenericNetlinkControlAttributeMcastGrp type is 1 byte (uint8)
	if size := unsafe.Sizeof(GenericNetlinkControlAttributeMcastGrpUnspec); size != 1 {
		t.Errorf("Expected size of GenericNetlinkControlAttributeMcastGrp to be 1 byte, got %d", size)
	}

	// Test the values of the constants
	tests := []struct {
		name  string
		value GenericNetlinkControlAttributeMcastGrp
		want  GenericNetlinkControlAttributeMcastGrp
	}{
		{"GenericNetlinkControlAttributeMcastGrpUnspec", GenericNetlinkControlAttributeMcastGrpUnspec, 0},
		{"GenericNetlinkControlAttributeMcastGrpName", GenericNetlinkControlAttributeMcastGrpName, 1},
		{"GenericNetlinkControlAttributeMcastGrpId", GenericNetlinkControlAttributeMcastGrpId, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("Expected %s to have value %d, but got %d", tt.name, tt.want, tt.value)
			}
		})
	}
}
