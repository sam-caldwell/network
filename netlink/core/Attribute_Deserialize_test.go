package core

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestAttributeDeserialize(t *testing.T) {
	tests := []struct {
		name        string
		data        []byte
		expected    *Attribute
		expectError bool
	}{
		{
			name: "Attribute without padding",
			data: []byte{
				// Length (nlaLen = 4 header + 4 value = 8), little-endian
				0x08, 0x00,
				// Type (NLA_F_NESTED = 0x8000), little-endian
				0x00, 0x80,
				// Value
				0x01, 0x02, 0x03, 0x04,
			},
			expected: &Attribute{
				Type:  NlaFNested,
				Value: []byte{0x01, 0x02, 0x03, 0x04},
			},
			expectError: false,
		},
		{
			name: "Attribute with padding",
			data: []byte{
				// Length (nlaLen = 4 header + 3 value = 7), little-endian
				0x07, 0x00,
				// Type (NLA_F_NET_BYTEORDER = 0x4000), little-endian
				0x00, 0x40,
				// Value (3 bytes)
				0xAA, 0xBB, 0xCC,
				// Padding (1 byte, should be ignored)
				0x00,
			},
			expected: &Attribute{
				Type:  NlaFNetByteorder,
				Value: []byte{0xAA, 0xBB, 0xCC},
			},
			expectError: false,
		},
		{
			name: "Attribute with insufficient data",
			data: []byte{
				0x04, 0x00, // Length (nlaLen = 4), little-endian
				0x00, 0x80, // Type (NLA_F_NESTED), little-endian
				// No Value
			},
			expected:    &Attribute{Type: NlaFNested, Value: []byte{}},
			expectError: false,
		},
		{
			name: "Data too short",
			data: []byte{
				0x03, 0x00, // Length (nlaLen = 3), little-endian (invalid, too short)
				0x00, 0x80, // Type
			},
			expected:    nil,
			expectError: true,
		},
		{
			name: "Declared length exceeds data length",
			data: []byte{
				0x0A, 0x00, // Length (nlaLen = 10), little-endian
				0x00, 0x80, // Type
				0xAA, 0xBB, // Only 2 bytes of Value, but length says there should be 6 bytes
			},
			expected:    nil,
			expectError: true,
		},
		{
			name: "Attribute with 1KB Value",
			data: func() []byte {
				nlaLen := 4 + 1024
				alignedLen := (nlaLen + 3) & ^3
				b := make([]byte, alignedLen)
				binary.LittleEndian.PutUint16(b[0:2], uint16(nlaLen))
				binary.LittleEndian.PutUint16(b[2:4], uint16(NlaFNested))
				// Value is already zeroed
				return b
			}(),
			expected: &Attribute{
				Type:  NlaFNested,
				Value: make([]byte, 1024),
			},
			expectError: false,
		},
		{
			name: "Attribute with invalid length (nlaLen less than header size)",
			data: []byte{
				0x02, 0x00, // Length (nlaLen = 2), invalid as it's less than header size
				0x00, 0x80, // Type
			},
			expected:    nil,
			expectError: true,
		},
	}

	t.Run("Test deserialize method", func(t *testing.T) {
		for _, tt := range tests {
			var attr Attribute
			t.Run(tt.name, func(t *testing.T) {
				err := attr.Deserialize(tt.data)
				if tt.expectError {
					if err == nil {
						t.Errorf("Expected error but got none")
					}
					return
				} else if err != nil {
					t.Errorf("Deserialize() returned unexpected error: %v", err)
					return
				}

				// Compare the Type
				if attr.Type != tt.expected.Type {
					t.Errorf("Type mismatch. Expected: %v, Got: %v", tt.expected.Type, attr.Type)
				}

				// Compare the Value
				if !bytes.Equal(attr.Value, tt.expected.Value) {
					t.Errorf("Value mismatch.\nExpected: %v\nGot:      %v", tt.expected.Value, attr.Value)
				}
			})
		}
	})
	t.Run("test deserialize function", func(t *testing.T) {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				attr, err := DeserializeAttribute(tt.data)
				if tt.expectError {
					if err == nil {
						t.Errorf("Expected error but got none")
					}
					return
				} else if err != nil {
					t.Errorf("Deserialize() returned unexpected error: %v", err)
					return
				}

				// Compare the Type
				if attr.Type != tt.expected.Type {
					t.Errorf("Type mismatch. Expected: %v, Got: %v", tt.expected.Type, attr.Type)
				}

				// Compare the Value
				if !bytes.Equal(attr.Value, tt.expected.Value) {
					t.Errorf("Value mismatch.\nExpected: %v\nGot:      %v", tt.expected.Value, attr.Value)
				}
			})
		}
	})
	t.Run("serialize-deserialize agreement", func(t *testing.T) {
		test := []struct {
			name string
			attr *Attribute
		}{
			{
				name: "",
				attr: &Attribute{
					Type:  NlaFNested,
					Value: []byte{0x01, 0x02, 0x03, 0x04},
				},
			},
			{
				name: "",
				attr: &Attribute{
					Type: NlaFNested,
					Value: []byte{
						0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
						0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F,
					},
				},
			},
		}
		var (
			err        error
			serialized []byte
		)
		for _, tt := range test {
			if serialized, err = tt.attr.Serialize(); err != nil {
				t.Fatal(err)
			}
			var attr Attribute
			if err = attr.Deserialize(serialized); err != nil {
				t.Fatal(err)
			}
			if attr.Type != tt.attr.Type {
				t.Fatalf("Type mismatch. Expected: %v, Got: %v", tt.attr.Type, attr.Type)
			}
			if !bytes.Equal(attr.Value, tt.attr.Value) {
				t.Fatalf("Value mismatch. Expected: %v, Got: %v", tt.attr.Value, attr.Value)
			}
		}
	})
}
