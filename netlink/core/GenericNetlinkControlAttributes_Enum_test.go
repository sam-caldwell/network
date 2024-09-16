package core

import (
	"testing"
	"unsafe"
)

// TestGenericNetlinkControlAttributesEnum tests the size of the GenericNetlinkControlAttributes type
// and the values of its enumerated constants.
func TestGenericNetlinkControlAttributesEnum(t *testing.T) {
	// Verify the size of the GenericNetlinkControlAttributes type is 1 byte (uint8)
	if size := unsafe.Sizeof(GenericNetlinkControlAttributesUnspec); size != 1 {
		t.Errorf("Expected size of GenericNetlinkControlAttributes to be 1 byte, got %d", size)
	}

	// Test the values of the constants
	tests := []struct {
		name  string
		value GenericNetlinkControlAttributes
		want  GenericNetlinkControlAttributes
	}{
		{"GenericNetlinkControlAttributesUnspec", GenericNetlinkControlAttributesUnspec, 0},
		{"GenericNetlinkControlAttributeFamilyId", GenericNetlinkControlAttributeFamilyId, 1},
		{"GenericNetlinkControlAttributeFamilyName", GenericNetlinkControlAttributeFamilyName, 2},
		{"GenericNetlinkControlAttributeVersion", GenericNetlinkControlAttributeVersion, 3},
		{"GenericNetlinkControlAttributeHdrsize", GenericNetlinkControlAttributeHdrsize, 4},
		{"GenericNetlinkControlAttributeMaxattr", GenericNetlinkControlAttributeMaxattr, 5},
		{"GenericNetlinkControlAttributeOps", GenericNetlinkControlAttributeOps, 6},
		{"GenericNetlinkControlAttributeMcastGroups", GenericNetlinkControlAttributeMcastGroups, 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("Expected %s to have value %d, but got %d", tt.name, tt.want, tt.value)
			}
		})
	}
}
