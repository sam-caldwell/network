package core

import (
	"testing"
)

func TestBridgeVlanInfo_Deserialize(t *testing.T) {
	// Test case: Valid byte slice input
	t.Run("ValidInput", func(t *testing.T) {
		// Create a byte slice representing a valid BridgeVlanInfo struct (4 bytes)
		// Example: Flags = 0x0001 (little-endian), Vid = 0x0002 (little-endian)
		input := []byte{0x01, 0x00, 0x02, 0x00}

		// Create an empty BridgeVlanInfo struct
		var bridge BridgeVlanInfo

		// Call Deserialize and check for no error
		err := bridge.Deserialize(input)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		// Check the deserialized values
		if bridge.Flags != 0x0001 {
			t.Errorf("expected Flags to be 0x0001, got 0x%x", bridge.Flags)
		}
		if bridge.Vid != 0x0002 {
			t.Errorf("expected Vid to be 0x0002, got 0x%x", bridge.Vid)
		}
	})

	// Test case: Byte slice too short
	t.Run("InvalidInputTooShort", func(t *testing.T) {
		// Create a byte slice that is too short (less than 4 bytes)
		input := []byte{0x01, 0x00}

		// Create an empty BridgeVlanInfo struct
		var bridge BridgeVlanInfo

		// Call Deserialize and expect an error
		err := bridge.Deserialize(input)
		if err == nil {
			t.Fatal("expected an error due to short input, but got none")
		}
	})

	// Test case: Empty byte slice
	t.Run("EmptyInput", func(t *testing.T) {
		// Create an empty byte slice
		input := []byte{}

		// Create an empty BridgeVlanInfo struct
		var bridge BridgeVlanInfo

		// Call Deserialize and expect an error
		err := bridge.Deserialize(input)
		if err == nil {
			t.Fatal("expected an error with empty input, but got none")
		}
	})
}
