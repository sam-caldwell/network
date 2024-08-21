//go:build !linux

package namespace

import (
	"testing"
)

func TestUnspecified_SetNamespace(t *testing.T) {
	err := SetNamespace(Handle{}, 0)
	if err == nil {
		t.Fatal("expected error")
	}
	if err.Error() != ErrNotImplemented {
		t.Fatal("expect NotImplemented")
	}
}
