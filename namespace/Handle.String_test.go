package namespace

import (
	"testing"
)

func TestHandle_String(t *testing.T) {

	startsWith := func(s, prefix string) bool {
		return len(s) >= len(prefix) && s[:len(prefix)] == prefix
	}

	t.Run("handle with valid device and inode", func(t *testing.T) {
		handle, err := Get() // Assumes Get() provides a valid handle for the current thread's namespace

		if err != nil {
			t.Fatalf("expected no error. got: %v", err)
		}

		expectedPrefix := "NS("

		if got := handle.String(); !startsWith(got, expectedPrefix) {
			t.Fatalf("unexpected string representation: %s", got)
		}

		// Clean up by closing the handle
		if err := handle.Close(); err != nil {
			t.Fatalf("expected no error when closing the file descriptor. got: %v", err)
		}
	})

	t.Run("handle with closed status", func(t *testing.T) {
		var testData Handle = -1
		expected := "NS(none)"
		if got := testData.String(); got != expected {
			t.Fatalf("expected %s, got %s", expected, got)
		}
	})
}
