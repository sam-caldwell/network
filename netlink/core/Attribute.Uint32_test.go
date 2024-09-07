package core

import (
	"testing"
)

// Test cases for Uint32 method.
func TestAttribute_Uint32(t *testing.T) {
	tests := []struct {
		name        string
		attr        Attribute
		want        uint32
		expectPanic bool
	}{
		{
			name: "BigEndian with NLA_F_NET_BYTEORDER flag set",
			attr: Attribute{
				Type:  NlaFNetByteorder,               // Simulate the Netlink attribute type with the network byte order flag set.
				Value: []byte{0x00, 0x00, 0x01, 0x00}, // Big-endian encoded 256 (0x00000100).
			},
			want:        256, // Expected result
			expectPanic: false,
		},
		{
			name: "LittleEndian without NLA_F_NET_BYTEORDER flag",
			attr: Attribute{
				Type:  0,                              // No byte order flag, should use native byte order (Little Endian).
				Value: []byte{0x00, 0x01, 0x00, 0x00}, // Little-endian encoded 256.
			},
			want:        256, // Expected result
			expectPanic: false,
		},
		{
			name: "Invalid length - triggers panic",
			attr: Attribute{
				Type:  NlaFNetByteorder,
				Value: []byte{0x00, 0x01}, // Invalid byte length for uint32.
			},
			want:        0, // Doesn't matter, we expect a panic.
			expectPanic: true,
		},
		{
			name: "Nil value - triggers panic",
			attr: Attribute{
				Type:  NlaFNetByteorder,
				Value: nil, // Nil value should trigger panic.
			},
			want:        0, // Doesn't matter, we expect a panic.
			expectPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test for panic
			if tt.expectPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("expected panic but did not get one")
					}
				}()
				_ = tt.attr.Uint32() // Call the function
			} else {
				// Test for correct output
				if got := tt.attr.Uint32(); got != tt.want {
					t.Errorf("Uint32() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
