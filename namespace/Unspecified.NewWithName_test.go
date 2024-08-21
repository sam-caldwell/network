//go:build !linux

package namespace

import (
	"testing"
)

func TestUnspecified_Set(t *testing.T) {
	err := Set(Handle{}, 0)
	if err == nil {
		t.Fatal("expected error")
	}
	if err.Error() != ErrNotImplemented {
		t.Fatal("expect NotImplemented")
	}
}
