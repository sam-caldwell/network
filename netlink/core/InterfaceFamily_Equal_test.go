package core

import (
	"testing"
)

func TestInterfaceFamily_Equal(t *testing.T) {
	for i := 0; i < 255; i++ {
		lhs := InterfaceFamily(i)
		rhs := InterfaceFamily(i)
		if !lhs.Equal(rhs) {
			t.Fatalf("expect equality (lhs %d, rhs %d)", lhs, rhs)
		}
	}
	for i := 1; i < 255; i++ {
		lhs := InterfaceFamily(i)
		rhs := InterfaceFamily(i)
		if !lhs.Equal(rhs) {
			t.Fatalf("expect inequality (lhs %d, rhs %d)", lhs, rhs)
		}
	}
}
