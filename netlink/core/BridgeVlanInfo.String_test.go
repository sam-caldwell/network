//go:build linux

package core

import (
	"testing"
)

func TestBridgeVlanInfo_String(t *testing.T) {
	tests := []struct {
		name     string
		bridge   BridgeVlanInfo
		expected string
	}{
		{
			name:     "Default",
			bridge:   BridgeVlanInfo{Flags: 0x0100, Vid: 100},
			expected: "BridgeVlanInfo{Flags: 0x0100, Vid: 100}",
		},
		{
			name:     "ZeroValues",
			bridge:   BridgeVlanInfo{Flags: 0x0000, Vid: 0},
			expected: "BridgeVlanInfo{Flags: 0x0000, Vid: 0}",
		},
		{
			name:     "MaxValues",
			bridge:   BridgeVlanInfo{Flags: 0xffff, Vid: 4095},
			expected: "BridgeVlanInfo{Flags: 0xffff, Vid: 4095}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.bridge.String()

			if result != tt.expected {
				t.Errorf("BridgeVlanInfo.String() = %v, want %v", result, tt.expected)
			}
		})
	}
}
