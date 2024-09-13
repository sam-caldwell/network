package core

import (
	"testing"
)

func TestCtAttrCounters_FromInt(t *testing.T) {
	var c CtAttrCounters
	for i := range []int{0, 1, 2, 4, 128, 255} {

		if c.FromInt(i); int(c) != i {
			t.Errorf("CtAttrCounters.FromInt(0,0)=%d, want %d", c, i)
		}

	}
}
