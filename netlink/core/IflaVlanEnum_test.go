package core

import (
	"testing"
)

// TestIflaVlanEnum tests the IflaVlanEnum values.
func TestIflaVlanEnum(t *testing.T) {
	tests := []struct {
		name     string
		enumVal  IflaVlanEnum
		expected IflaVlanEnum
	}{
		{"IflaVlanUnspec", IflaVlanUnspec, 0},
		{"IflaVlanId", IflaVlanId, 1},
		{"IflaVlanFlags", IflaVlanFlags, 2},
		{"IflaVlanEgressQos", IflaVlanEgressQos, 3},
		{"IflaVlanIngressQos", IflaVlanIngressQos, 4},
		{"IflaVlanProtocol", IflaVlanProtocol, 5},
		{"IflaVlanMax", IflaVlanMax, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.enumVal != tt.expected {
				t.Errorf("Expected %s to be %d, got %d", tt.name, tt.expected, tt.enumVal)
			}
		})
	}
}
