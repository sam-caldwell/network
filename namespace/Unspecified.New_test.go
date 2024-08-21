//go:build !linux

package namespace

import (
	"testing"
)

func TestUnspecified_New(t *testing.T) {
	err := New(Handle{}, 0)
	if err == nil {
		t.Fatal("expected error")
	}
	if err.Error() != ErrNotImplemented {
		t.Fatal("expect NotImplemented")
	}
}
