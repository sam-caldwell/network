//go:build !linux

package namespace

import (
	"testing"
)

func TestUnspecified_GetFromPid(t *testing.T) {
	err := GetFromPid(1)
	if err == nil {
		t.Fatal("expected error")
	}
	if err.Error() != ErrNotImplemented {
		t.Fatal("expect NotImplemented")
	}
}
