//go:build !linux

package namespace

import (
	"testing"
)

func TestUnspecified_GetFromThread(t *testing.T) {
	err := GetFromThread(1, 1)
	if err == nil {
		t.Fatal("expected error")
	}
	if err.Error() != ErrNotImplemented {
		t.Fatal("expect NotImplemented")
	}
}
