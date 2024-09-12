package core

import (
	"testing"
	"unsafe"
)

func TestMplsLabelStackShiftMask(t *testing.T) {
	// Subtest for checking the size of the MplsLabelStackShiftMask type
	t.Run("TestSize", func(t *testing.T) {
		const expectedSizeInBytes = 4
		var mask MplsLabelStackShiftMask

		if size := unsafe.Sizeof(mask); size != uintptr(expectedSizeInBytes) {
			t.Errorf("Expected size %d but got %d", expectedSizeInBytes, size)
		}
	})

	// Subtest for checking the values of the constants
	t.Run("TestValues", func(t *testing.T) {
		tests := []struct {
			name     string
			value    MplsLabelStackShiftMask
			expected MplsLabelStackShiftMask
		}{
			{"MplsLsLabelMask", MplsLsLabelMask, 0xFFFFF000},
			{"MplsLsTcMask", MplsLsTcMask, 0x00000E00},
			{"MplsLsSMask", MplsLsSMask, 0x00000100},
			{"MplsLsTtlMask", MplsLsTtlMask, 0x000000FF},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				if test.value != test.expected {
					t.Errorf("Expected 0x%X but got 0x%X for %s", test.expected, test.value, test.name)
				}
			})
		}
	})
}
