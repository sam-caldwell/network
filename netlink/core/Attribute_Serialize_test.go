package core

import (
	"bytes"
	"testing"
)

func TestAttributeSerialize(t *testing.T) {
	tests := []struct {
		name        string
		attr        *Attribute
		expected    []byte
		expectError bool
	}{
		{
			name: "Attribute without padding",
			attr: &Attribute{
				Type:  NlaFNested,
				Value: []byte{0x01, 0x02, 0x03, 0x04}, // 4 bytes
			},
			expected: []byte{
				// Length (nlaLen = 4 header + 4 value = 8), little-endian
				0x08, 0x00,
				// Type (NLA_F_NESTED = 0x8000), little-endian
				0x00, 0x80,
				// Value
				0x01, 0x02, 0x03, 0x04,
			},
			expectError: false,
		},
		{
			name: "Attribute with padding",
			attr: &Attribute{
				Type:  NlaFNetByteorder,
				Value: []byte{0xAA, 0xBB, 0xCC}, // 3 bytes, requires 1 byte padding
			},
			expected: []byte{
				// Length (nlaLen = 4 header + 3 value = 7), little-endian
				0x07, 0x00,
				// Type (NLA_F_NET_BYTEORDER = 0x4000), little-endian
				0x00, 0x40,
				// Value
				0xAA, 0xBB, 0xCC,
				// Padding (1 byte, should be zero)
				0x00,
			},
			expectError: false,
		},
		{
			name: "Attribute exceeding maximum length",
			attr: &Attribute{
				Type:  NlaFNested,
				Value: make([]byte, 0xFFFF), // Exceeds maximum nlaLen when header is added
			},
			expected:    nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serialized, err := tt.attr.Serialize()
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
				return
			} else if err != nil {
				t.Errorf("Serialize() returned unexpected error: %v", err)
				return
			}

			if !bytes.Equal(serialized, tt.expected) {
				t.Errorf("Serialized output does not match expected value.\n"+
					"Expected: %v\n"+
					"Got:      %v", tt.expected, serialized)
			}
		})
	}
}
