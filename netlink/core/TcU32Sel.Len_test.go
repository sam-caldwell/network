package core

import (
	"testing"
)

func TestTcU32SelLen(t *testing.T) {
	// Test case for default (empty) TcU32Sel
	t.Run("test with defaults", func(t *testing.T) {
		msg := TcU32Sel{}

		actualSize := msg.Len()
		expectedSize := SizeOfTcU32Sel // No keys, so just the base size of the struct

		if actualSize != expectedSize {
			t.Fatalf("size mismatch.\n"+
				"    actualSize: %d\n"+
				"  expectedSize: %d\n"+
				"SizeOfTcU32Key: %d\n"+
				"SizeOfTcU32Sel: %d",
				actualSize, expectedSize,
				SizeOfTcU32Key, SizeOfTcU32Sel)
		}
	})

	// Test case for varying numbers of keys (Nkeys)
	t.Run("test with various Nkeys values", func(t *testing.T) {
		for i := 1; i < 255; i += 2 {
			msg := TcU32Sel{
				Nkeys: uint8(i),
				Keys:  make([]TcU32Key, i),
			}

			actualSize := msg.Len()
			expectedSize := SizeOfTcU32Sel + i*SizeOfTcU32Key // Base size + Nkeys * size of each key

			if actualSize != expectedSize {
				t.Fatalf("size mismatch.\n"+
					"             i: %d,\n"+
					"    actualSize: %d\n"+
					"  expectedSize: %d\n"+
					"SizeOfTcU32Key: %d\n"+
					"SizeOfTcU32Sel: %d\n",
					i, actualSize, expectedSize,
					SizeOfTcU32Key, SizeOfTcU32Sel)
			}
		}
	})
}
