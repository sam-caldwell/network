package core

import (
	"testing"
	"unsafe"
)

func TestAtmCell_Enum(t *testing.T) {
	t.Run("Size validation", func(t *testing.T) {
		const expectedSizeInBytes = 1
		var o AtmCell
		if sz := int(unsafe.Sizeof(o)); sz != expectedSizeInBytes {
			t.Fatalf("Size mismatch.  Got %d", sz)
		}
	})
	t.Run("value validation", func(t *testing.T) {
		if int(ATMCellPayload) != 48 {
			t.Fatal("Payload mismatch")
		}
		if int(ATMCellSize) != 53 {
			t.Fatal("Size mismatch")
		}
	})
}
