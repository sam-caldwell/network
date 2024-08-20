package namespace

import (
	"testing"
)

func TestConstants(t *testing.T) {

	if bindMountPath != "/run/netns" {
		t.Fatal("mismatch")
	}

	if closedHandle != -1 {
		t.Fatal("mismatch")
	}

	if ErrNotImplemented != "not implemented" {
		t.Fatal("mismatch")
	}

}
