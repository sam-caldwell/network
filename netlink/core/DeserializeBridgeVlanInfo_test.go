package core

import (
	"testing"
)

func TestDeserializeBridgeVlanInfo(t *testing.T) {
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
}
