package core

import (
	"testing"
)

// TestIflaVfLinkStateEnum tests the IflaVfLinkStateEnum values.
func TestIflaVfLinkStateEnum(t *testing.T) {
	// Test that IflaVfLinkStateAuto is set to 0
	if IflaVfLinkStateAuto != 0 {
		t.Errorf("Expected IflaVfLinkStateAuto to be 0, got %d", IflaVfLinkStateAuto)
	}

	// Test that IflaVfLinkStateEnable is set to 1
	if IflaVfLinkStateEnable != 1 {
		t.Errorf("Expected IflaVfLinkStateEnable to be 1, got %d", IflaVfLinkStateEnable)
	}

	// Test that IflaVfLinkStateDisable is set to 2
	if IflaVfLinkStateDisable != 2 {
		t.Errorf("Expected IflaVfLinkStateDisable to be 2, got %d", IflaVfLinkStateDisable)
	}

	// Test that IflaVfLinkStateMax is equal to IflaVfLinkStateDisable (2)
	if IflaVfLinkStateMax != IflaVfLinkStateDisable {
		t.Errorf("Expected IflaVfLinkStateMax to be %d, got %d", IflaVfLinkStateDisable, IflaVfLinkStateMax)
	}
}
