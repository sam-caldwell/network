//go:build !linux

package namespace

import (
	"testing"
)

func TestUnspecified_Get(t *testing.T) {
	err := Get()
	if err == nil {
		t.Fatal("expected error")
	}
	if err.Error() != ErrNotImplemented {
		t.Fatal("expect NotImplemented")
	}
}
