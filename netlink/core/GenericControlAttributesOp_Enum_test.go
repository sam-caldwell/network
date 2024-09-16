package core

import (
	"testing"
	"unsafe"
)

// TestGenericControlAttributesOpEnum tests the size of the GenericControlAttributesOp type
// and the values of its enumerated constants.
func TestGenericControlAttributesOpEnum(t *testing.T) {
	// Verify the size of the GenericControlAttributesOp type is 1 byte (uint8)
	if size := unsafe.Sizeof(GenericNetlinkControlAttributeOpUnspec); size != 1 {
		t.Errorf("Expected size of GenericControlAttributesOp to be 1 byte, got %d", size)
	}

	// Test the values of the constants
	tests := []struct {
		name  string
		value GenericControlAttributesOp
		want  GenericControlAttributesOp
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
