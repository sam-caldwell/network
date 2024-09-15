package core

import (
	"bytes"
	"os"
	"testing"
)

func TestPrintRawAttributes(t *testing.T) {
	tests := []struct {
		name        string
		attributes  []byte
		expectedOut string
		expectError bool
	}{
		{
			name: "Simple attribute without padding",
			attributes: func() []byte {
				attr := &Attribute{
					Type:  NlaFlags(1),
					Value: []byte{0xAA, 0xBB, 0xCC, 0xDD}, // 4 bytes, no padding needed
				}
				data, _ := attr.Serialize()
				return data
			}(),
			expectedOut: "type=1 nested=false len=4 [170 187 204 221]\n",
			expectError: false,
		},
		{
			name: "Attribute with padding",
			attributes: func() []byte {
				attr := &Attribute{
					Type:  NlaFlags(2),
					Value: []byte{0xAA, 0xBB, 0xCC}, // 3 bytes, requires 1 byte padding
				}
				data, _ := attr.Serialize()
				return data
			}(),
			expectedOut: "type=2 nested=false len=3 [170 187 204]\n",
			expectError: false,
		},
		{
			name: "Nested attributes",
			attributes: func() []byte {
				// Create a nested attribute
				nestedAttr := &Attribute{
					Type:  NlaFlags(3),
					Value: []byte{0x11, 0x22}, // Inner attribute value
				}
				nestedData, _ := nestedAttr.Serialize()

				// Wrap it in an outer attribute marked as nested
				outerAttr := &Attribute{
					Type:  NlaFlags(NlaFNested | 4), // Set nested flag
					Value: nestedData,
				}
				data, _ := outerAttr.Serialize()
				return data
			}(),
			expectedOut: "" +
				"type=4 nested=true len=8 [6 0 3 0 17 34 0 0]\n" +
				"> type=3 nested=false len=2 [17 34]\n",
			expectError: false,
		},
		{
			name: "Invalid attribute length",
			attributes: []byte{
				0x02, 0x00, // Length (nlaLen = 2), invalid as it's less than header size
				0x01, 0x00, // Type
			},
			expectedOut: "",
			expectError: true,
		},
		{
			name: "Attribute exceeding data size",
			attributes: []byte{
				0x08, 0x00, // Length (nlaLen = 8)
				0x01, 0x00, // Type
				0xAA, 0xBB, 0xCC, // Only 3 bytes of value, but length says 4 bytes
			},
			expectedOut: "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Redirect stdout
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Run the function
			err := PrintRawAttributes(tt.attributes)

			// Close and restore stdout
			w.Close()
			os.Stdout = oldStdout

			// Read the captured output
			var buf bytes.Buffer
			_, _ = buf.ReadFrom(r)
			actualOutput := buf.String()

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			// Compare the output
			if actualOutput != tt.expectedOut {
				t.Errorf("Output mismatch.\n"+
					"Expected:\n%s\n"+
					"Got:\n%s", tt.expectedOut, actualOutput)
			}
		})
	}
}
