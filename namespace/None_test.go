package namespace

import (
	"testing"
)

func TestNone(t *testing.T) {
	if closedHandle != -1 {
		t.Fatalf("closedHandle should be -1")
	}
	if ns := None(); ns != Handle(closedHandle) {
		t.Fatalf("None() should return closedHandle")
	}
}
