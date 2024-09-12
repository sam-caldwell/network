package core

import (
	"testing"
	"unsafe"
)

// TestGenericControlAttributesOpEnum tests the size of the GenericControlAttributesOpEnum type
// and the values of its enumerated constants.
func TestGenericControlAttributesOpEnum(t *testing.T) {
	// Verify the size of the GenericControlAttributesOpEnum type is 1 byte (uint8)
	if size := unsafe.Sizeof(GenericNetlinkControlAttributeOpUnspec); size != 1 {
		t.Errorf("Expected size of GenericControlAttributesOpEnum to be 1 byte, got %d", size)
	}

	// Test the values of the constants
	tests := []struct {
		name  string
		value GenericControlAttributesOpEnum
		want  GenericControlAttributesOpEnum
	}{
		{"GenericNetlinkControlAttributeOpUnspec", GenericNetlinkControlAttributeOpUnspec, 0},
		{"GenericNetlinkControlAttributeOpId", GenericNetlinkControlAttributeOpId, 1},
		{"GenericNetlinkControlAttributeOpFlags", GenericNetlinkControlAttributeOpFlags, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("Expected %s to have value %d, but got %d", tt.name, tt.want, tt.value)
			}
		})
	}
}
