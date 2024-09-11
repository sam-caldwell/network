package namespace

import (
	"golang.org/x/sys/unix"
	"testing"
)

func TestGet(t *testing.T) {
	t.Run("retrieve the current thread's network namespace handle", func(t *testing.T) {
		handle, err := Get()
		if err != nil {
			t.Fatalf("expected no error. got: %v", err)
		}
		if handle == closedHandle {
			t.Fatalf("expected handle to be valid, got closedHandle")
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
}
