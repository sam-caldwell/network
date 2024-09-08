package core

import (
	"testing"
)

func TestCurve_Set(t *testing.T) {
	var c Curve
	t.Run("Pre-test", func(t *testing.T) {
		if c.m1 != 0 || c.m2 != 0 || c.d != 0 {
			t.Fatal("initial state is wrong")
		}
	})
	t.Run("Use .Set() method", func(t *testing.T) {
		c.Set(1, 2, 3)
	})
	t.Run("verify outcome", func(t *testing.T) {
		if c.m1 != 1 {
			t.Fatalf("outcome (%d) verification failed", c.m1)
		}
		if c.d != 2 {
			t.Fatalf("outcome (%d) verification failed", c.d)
		}
		if c.m2 != 3 {
			t.Fatalf("outcome (%d) verification failed", c.m2)
		}
	})
}
