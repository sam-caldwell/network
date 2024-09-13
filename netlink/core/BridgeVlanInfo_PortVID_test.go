package core

import (
	"testing"
)

func TestBridgeVlanInfo_PortVID(t *testing.T) {
	// Define a test case with the PVID flag set
	t.Run("PVID set", func(t *testing.T) {
		bridge := BridgeVlanInfo{
			Flags: BridgeVlanInfoPvid, // PVID flag set
			Vid:   100,
		}

		if !bridge.PortVID() {
			t.Errorf("Expected PortVID() to return true, but got false")
		}
	})

	// Define a test case with the PVID flag not set
	t.Run("PVID not set", func(t *testing.T) {
		bridge := BridgeVlanInfo{
			Flags: 0, // No flags set
			Vid:   100,
		}

		if bridge.PortVID() {
			t.Errorf("Expected PortVID() to return false, but got true")
		}
	})
}
