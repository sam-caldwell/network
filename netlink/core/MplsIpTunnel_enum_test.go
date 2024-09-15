package core

import (
	"testing"
	"unsafe"
)

func TestMplsIptunnel(t *testing.T) {
	// Subtest for checking the size of the MplsIptunnel type
	t.Run("TestSize", func(t *testing.T) {
		const expectedSizeInBytes = 1
		var tunnel MplsIptunnel

		if size := unsafe.Sizeof(tunnel); size != uintptr(expectedSizeInBytes) {
			t.Errorf("Expected size %d but got %d", expectedSizeInBytes, size)
		}
	})

	// Subtest for checking the values of the constants
	t.Run("TestValues", func(t *testing.T) {
		tests := []struct {
			name     string
			value    MplsIptunnel
			expected MplsIptunnel
		}{
			{"MplsIptunnelUnspec", MplsIptunnelUnspec, 0},
			{"MplsIptunnelDst", MplsIptunnelDst, 1},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				if test.value != test.expected {
					t.Errorf("Expected %d but got %d for %s", test.expected, test.value, test.name)
				}
			})
		}
	})
}
