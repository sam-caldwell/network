package core

import (
	"fmt"
	"testing"
)

func TestCtAttrCounters_ToInt(t *testing.T) {
	for v := range []int{0, 1, 128, 255} {
		t.Run(fmt.Sprintf("%s(%d)", t.Name(), v), func(t *testing.T) {
			o := CtAttrCounters(v)
			if int(o) != v {
				t.Fatalf("value mismatch on %d", v)
			}
		})
	}
}
