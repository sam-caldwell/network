package core

import (
	"testing"
)

func TestBridgeVlanInfo_Deserialize_Methods_and_Function(t *testing.T) {

	t.Run("DeserializeBridgeVlanInfo() function", func(t *testing.T) {
		// Define some test cases
		tests := []struct {
			name          string
			input         []byte
			expectedVid   VlanIdType
			expectedFlags BridgeVlanInfoFlags
			expectError   bool
		}{
			{
				name: "Valid VLAN Info",
				// Example: Flags = 1 (0x0001), VID = 5 (0x0005)
				input:         []byte{0x01, 0x00, 0x05, 0x00},
				expectedVid:   5,
				expectedFlags: 1,
				expectError:   false,
			},
			{
				name:          "Too Short Input",
				input:         []byte{0x01, 0x00},
				expectedVid:   0,
				expectedFlags: 0,
				expectError:   true,
			},
			{
				name: "Another Valid VLAN Info",
				// Example: Flags = 2 (0x0002), VID = 10 (0x000A)
				input:         []byte{0x02, 0x00, 0x0A, 0x00},
				expectedVid:   10,
				expectedFlags: 2,
				expectError:   false,
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				vlanInfo, err := DeserializeBridgeVlanInfo(test.input)

				if (err != nil) != test.expectError {
					t.Errorf("expected error: %v, got error: %v", test.expectError, err)
				}

				if err == nil {
					if vlanInfo.Vid != test.expectedVid {
						t.Errorf("expected VID: %d, got: %d", test.expectedVid, vlanInfo.Vid)
					}

					if vlanInfo.Flags != test.expectedFlags {
						t.Errorf("expected Flags: %d, got: %d", test.expectedFlags, vlanInfo.Flags)
					}
				}
			})
		}
	})

	t.Run("test Deserialize() method", func(t *testing.T) {
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
	})
}
