//go:build !linux

package namespace

import (
	"testing"
)

func TestUnspecified_GetFromName(t *testing.T) {
	err := GetFromName("foo")
	if err == nil {
		t.Fatal("expected error")
	}
	if err.Error() != ErrNotImplemented {
		t.Fatal("expect NotImplemented")
	}
}
