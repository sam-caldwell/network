package namespace

import (
	"testing"
)

func TestNone(t *testing.T) {
	if None() != closedHandle {
		t.Fatal("expected -1")
	}
}
