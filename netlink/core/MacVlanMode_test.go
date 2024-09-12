package core

import (
	"testing"
	"unsafe"
)

func TestMacVlanMode(t *testing.T) {
	t.Run("TestSize", func(t *testing.T) {
		const expectedSizeInBytes = 1
		var mode MacVlanMode

		if size := unsafe.Sizeof(mode); size != uintptr(expectedSizeInBytes) {
			t.Errorf("Expected size %d but got %d", expectedSizeInBytes, size)
		}
	})

	t.Run("TestValues", func(t *testing.T) {
		tests := []struct {
			name     string
			value    MacVlanMode
			expected MacVlanMode
		}{
			{"MacvlanModePrivate", MacvlanModePrivate, 1},
			{"MacvlanModeVepa", MacvlanModeVepa, 2},
			{"MacvlanModeBridge", MacvlanModeBridge, 4},
			{"MacvlanModePassthru", MacvlanModePassthru, 8},
			{"MacvlanModeSource", MacvlanModeSource, 16},
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
