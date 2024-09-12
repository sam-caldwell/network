package core

import (
	"testing"
)

// TestIflaVfStatsEnum tests the IflaVfStatsEnum values.
func TestIflaVfStatsEnum(t *testing.T) {
	tests := []struct {
		name     string
		enumVal  IflaVfStatsEnum
		expected IflaVfStatsEnum
	}{
		{"IflaVfStatsRxPackets", IflaVfStatsRxPackets, 0},
		{"IflaVfStatsTxPackets", IflaVfStatsTxPackets, 1},
		{"IflaVfStatsRxBytes", IflaVfStatsRxBytes, 2},
		{"IflaVfStatsTxBytes", IflaVfStatsTxBytes, 3},
		{"IflaVfStatsBroadcast", IflaVfStatsBroadcast, 4},
		{"IflaVfStatsMulticast", IflaVfStatsMulticast, 5},
		{"IflaVfStatsRxDropped", IflaVfStatsRxDropped, 6},
		{"IflaVfStatsTxDropped", IflaVfStatsTxDropped, 7},
		{"IflaVfStatsMax", IflaVfStatsMax, IflaVfStatsTxDropped},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.enumVal != tt.expected {
				t.Errorf("Expected %s to be %d, got %d", tt.name, tt.expected, tt.enumVal)
			}
		})
	}
}
