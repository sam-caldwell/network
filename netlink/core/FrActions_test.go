package core

import (
	"testing"
	"unsafe"
)

// TestFrActionsEnum tests the size of the FrActions type and the values of the enumerated constants.
func TestFrActionsEnum(t *testing.T) {
	const expectedSizeInBytes = 1
	// Verify the size of the FrActions type is 1 byte (uint8)
	if size := unsafe.Sizeof(FrActUnspec); size != expectedSizeInBytes {
		t.Errorf("Expected size of FrActions to be 1 byte, got %d", size)
	}

	// Test the values of the constants
	tests := []struct {
		name  string
		value FrActions
		want  FrActions
	}{
		{"FrActUnspec", FrActUnspec, 0},
		{"FrActToTbl", FrActToTbl, 1},
		{"FrActGoto", FrActGoto, 2},
		{"FrActNop", FrActNop, 3},
		{"FrActRes3", FrActRes3, 4},
		{"FrActRes4", FrActRes4, 5},
		{"FrActBlackhole", FrActBlackhole, 6},
		{"FrActUnreachable", FrActUnreachable, 7},
		{"FrActProhibit", FrActProhibit, 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("Expected %s to have value %d, but got %d", tt.name, tt.want, tt.value)
			}
		})
	}
}
