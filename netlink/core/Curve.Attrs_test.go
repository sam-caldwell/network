package core

import (
	"testing"
)

func TestCurve_Attrs(t *testing.T) {
	c := Curve{
		m1: 1,
		d:  2,
		m2: 3,
	}
	if m1, d, m2 := c.Attrs(); m1 != 1 || d != 2 || m2 != 3 {
		t.Fatalf("attrs failed")
	}
}
