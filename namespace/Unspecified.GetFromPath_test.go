//go:build !linux

package namespace

import (
	"testing"
)

func TestUnspecified_GetFromPath(t *testing.T) {
	err := GetFromPath("path")
	if err == nil {
		t.Fatal("expected error")
	}
	if err.Error() != ErrNotImplemented {
		t.Fatal("expect NotImplemented")
	}
}
