package core

import (
	"bytes"
	"encoding/binary"
	"testing"
	"unsafe"
)

func TestBridgeVlanInfo(t *testing.T) {
	t.Run("verify SizeOfBridgeVlanInfo", func(t *testing.T) {
		const expectedSizeInBytes = 4
		if SizeOfBridgeVlanInfo != expectedSizeInBytes {
			t.Fatalf("SizeOfBridgeVlanInfo mismatch")
		}
	})
	t.Run("BridgeVlanInfo struct test", func(t *testing.T) {
		t.Run("check structure", func(t *testing.T) {
			_ = BridgeVlanInfo{
				Flags: BridgeVlanInfoFlags(0),
				Vid:   VlanIdType(0),
			}
		})
		t.Run("test struct size", func(t *testing.T) {
			const expectedSizeInBytes = 4
			o := BridgeVlanInfo{}
			if sz := unsafe.Sizeof(o); sz != expectedSizeInBytes {
				t.Fatalf("struct size mismatch: got %d, want %d", sz, expectedSizeInBytes)
			}
		})
		t.Run("test Flags size", func(t *testing.T) {
			const expectedSizeInBytes = 2
			o := BridgeVlanInfo{}
			if sz := unsafe.Sizeof(o.Flags); sz != expectedSizeInBytes {
				t.Fatalf("size (Flags) mismatch: got %d, want %d", sz, expectedSizeInBytes)
			}
		})
		t.Run("test Vid size", func(t *testing.T) {
			const expectedSizeInBytes = 2
			o := BridgeVlanInfo{}
			if sz := unsafe.Sizeof(o.Vid); sz != expectedSizeInBytes {
				t.Fatalf("size (Vid) mismatch: got %d, want %d", sz, expectedSizeInBytes)
			}
		})
	})

	t.Run("Test BridgeVlanInfoFlags enum", func(t *testing.T) {
		t.Run("size test", func(t *testing.T) {
			const expectedSizeInBytes = 2
			o := BridgeVlanInfoFlags(0)
			if unsafe.Sizeof(o) != expectedSizeInBytes {
				t.Fatalf("size mismatch")
			}
		})
		t.Run("value test", func(t *testing.T) {
			testData := map[BridgeVlanInfoFlags]BridgeVlanInfoFlags{
				BridgeVlanInfoFlagsUnset: 0,
				BridgeVlanInfoMaster:     1,
				BridgeVlanInfoPvid:       2,
				BridgeVlanInfoUntagged:   4,
				BridgeVlanInfoRangeBegin: 8,
				BridgeVlanInfoRangeEnd:   16,
				BridgeVlanInfoBrentry:    32,
				BridgeVlanInfoOnlyOpts:   64,
			}
			for actual, expected := range testData {
				if actual != expected {
					t.Fatalf("value mismatch: expected %d, actual %d", expected, actual)
				}
			}
		})
	})

	t.Run("IngressUntag() method", func(t *testing.T) {
		tests := []struct {
			name   string
			bridge BridgeVlanInfo
			want   bool
		}{
			{
				name:   "Untagged flag is set",
				bridge: BridgeVlanInfo{Flags: BridgeVlanInfoUntagged},
				want:   true,
			},
			{
				name:   "Untagged flag is not set",
				bridge: BridgeVlanInfo{Flags: 0},
				want:   false,
			},
			{
				name:   "Multiple flags, including untagged",
				bridge: BridgeVlanInfo{Flags: BridgeVlanInfoUntagged | 2},
				want:   true,
			},
			{
				name:   "Multiple flags, untagged not set",
				bridge: BridgeVlanInfo{Flags: 2},
				want:   false,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := tt.bridge.IngressUntag()
				if got != tt.want {
					t.Errorf("IngressUntag() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("Test PortVid() method", func(t *testing.T) {
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
	})

	t.Run("Test String() method", func(t *testing.T) {
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
	})
	t.Run("test Serialize() method", func(t *testing.T) {
		tests := []struct {
			name    string
			input   BridgeVlanInfo
			want    []byte
			wantErr bool
		}{
			{
				name: "Valid BridgeVlanInfo",
				input: BridgeVlanInfo{
					Flags: 0x0100, // Example flag value
					Vid:   100,    // Example VLAN ID (100)
				},
				// Expected result is little-endian representation of Flags and Vid.
				want: func() []byte {
					buf := new(bytes.Buffer)
					_ = binary.Write(buf, NativeEndian, uint16(0x0100)) // Serialize Flags
					_ = binary.Write(buf, NativeEndian, uint16(100))    // Serialize Vid
					return buf.Bytes()
				}(),
				wantErr: false,
			},
			{
				name: "Zero BridgeVlanInfo",
				input: BridgeVlanInfo{
					Flags: 0x0000, // No flags
					Vid:   0,      // VLAN ID = 0 (undefined VLAN ID)
				},
				// Expected result for zero fields.
				want: func() []byte {
					buf := new(bytes.Buffer)
					_ = binary.Write(buf, NativeEndian, uint16(0x0000)) // Serialize Flags
					_ = binary.Write(buf, NativeEndian, uint16(0))      // Serialize Vid
					return buf.Bytes()
				}(),
				wantErr: false,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := tt.input.Serialize()

				if (err != nil) != tt.wantErr {
					t.Errorf("Serialize() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				if !bytes.Equal(got, tt.want) {
					t.Errorf("Serialize() got = %v, want %v", got, tt.want)
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
}
