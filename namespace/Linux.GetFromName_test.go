//go:build linux

package namespace

import (
	"golang.org/x/sys/unix"
	"testing"
)

func TestGetFromName(t *testing.T) {
	t.Run("retrieve a handle for a named network namespace", func(t *testing.T) {
		name := "my-namespace" // Adjust to a valid namespace name for testing
		handle, err := GetFromName(name)
		if err != nil {
			t.Fatalf("expected no error. got: %v", err)
		}
		if handle < 0 {
			t.Fatalf("expected handle to be a valid file descriptor, got: %d", handle)
		}

		// Check if the handle is valid by attempting to get file status
		var stat unix.Stat_t
		err = unix.Fstat(int(handle), &stat)
		if err != nil {
			t.Fatalf("expected no error. got: %v", err)
		}

		// Clean up by closing the handle
		if err := handle.Close(); err != nil {
			t.Fatalf("expected no error when closing the file descriptor. got: %v", err)
		}
	})

	t.Run("handle non-existent namespace", func(t *testing.T) {
		name := "non-existent-namespace"
		handle, err := GetFromName(name)
		if err == nil {
			t.Fatalf("expected an error for non-existent namespace")
		}
		if handle != closedHandle {
			t.Fatalf("expected handle to be %v, got %v", closedHandle, handle)
		}
	})
}
