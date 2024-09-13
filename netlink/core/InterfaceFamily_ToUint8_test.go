package core

import (
	"testing"
)

func TestInterfaceFamily_ToUint8(t *testing.T) {
	for i := 0; i < 255; i++ {
		lhs := InterfaceFamily(i)
		rhs := uint8(i)
		if lhs.ToUint8() != rhs {
			t.Fatalf("expect equality (lhs %d, rhs %d)", lhs, rhs)
		}
	}
}
