package core

import (
	"testing"
	"unsafe"
)

func TestNlaFlagsConstants(t *testing.T) {
	t.Run("Test NlaFlags constants", func(t *testing.T) {
		// Validate NlaFNested
		if NlaFNested != NlaFlags(32768) {
			t.Errorf("Expected NlaFNested to be 32768, got %d", NlaFNested)
		}

		// Validate NlaFNetByteorder
		if NlaFNetByteorder != NlaFlags(16384) {
			t.Errorf("Expected NlaFNetByteorder to be 16384, got %d", NlaFNetByteorder)
		}

		// Validate NlaTypeMask
		expectedTypeMask := NlaFlags(16383)
		if NlaTypeMask != expectedTypeMask {
			t.Errorf("Expected NlaTypeMask to be %d, got %d", expectedTypeMask, NlaTypeMask)
		}

		// Validate NlaAlignto
		expectedAlignTo := NlaFlags(unsafe.Sizeof(NlaTypeMask) * 2)
		if NlaAlignto != expectedAlignTo {
			t.Errorf("Expected NlaAlignto to be %d, got %d", expectedAlignTo, NlaAlignto)
		}
	})
}
