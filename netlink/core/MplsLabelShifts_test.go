package core

import (
	"testing"
	"unsafe"
)

func TestMplsLabels(t *testing.T) {
	// Subtest for checking the size of the MplsLabels type
	t.Run("TestSize", func(t *testing.T) {
		const expectedSizeInBytes = 4
		var label uint32

		if size := unsafe.Sizeof(label); size != uintptr(expectedSizeInBytes) {
			t.Errorf("Expected size %d but got %d", expectedSizeInBytes, size)
		}
	})

	// Subtest for checking the values of the constants
	t.Run("TestValues", func(t *testing.T) {
		tests := []struct {
			name     string
			value    uint32
			expected uint32
		}{
			{"MplsLsLabelShift", MplsLsLabelShift, 12},
			{"MplsLsTcShift", MplsLsTcShift, 9},
			{"MplsLsSShift", MplsLsSShift, 8},
			{"MplsLsTtlShift", MplsLsTtlShift, 0},
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
