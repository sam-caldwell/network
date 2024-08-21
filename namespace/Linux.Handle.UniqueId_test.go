package namespace

import (
	"testing"
)

func TestUniqueId(t *testing.T) {
	startsWith := func(s, prefix string) bool {
		return len(s) >= len(prefix) && s[:len(prefix)] == prefix
	}
	t.Run("valid namespace handle", func(t *testing.T) {
		handle, err := Get() // Assumes Get() provides a valid handle for the current thread's namespace
		if err != nil {
			t.Fatalf("expected no error. got: %v", err)
		}
		expectedPrefix := "NS("
		if got := handle.UniqueId(); !startsWith(got, expectedPrefix) {
			t.Fatalf("unexpected unique ID representation: %s", got)
		}

		// Clean up by closing the handle
		if err := handle.Close(); err != nil {
			t.Fatalf("expected no error when closing the file descriptor. got: %v", err)
		}
	})

	t.Run("closed handle", func(t *testing.T) {
		var closedHandle Handle = -1
		expected := "NS(none)"
		if got := closedHandle.UniqueId(); got != expected {
			t.Fatalf("expected %s, got %s", expected, got)
		}
	})
}
