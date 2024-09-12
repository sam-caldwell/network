package namespace

import (
	"testing"
)

func TestHandle_Equal(t *testing.T) {
	// Mock valid file descriptors for testing purposes
	// These would normally be real file descriptors, but for unit testing, we mock them.
	handle1 := Handle(10)
	handle2 := Handle(10) // Same as handle1
	handle3 := Handle(20) // Different handle
	closedHandle := Handle(-1)

	t.Run("Handles are equal (same handle)", func(t *testing.T) {
		if !handle1.Equal(handle2) {
			t.Errorf("Expected handles to be equal, got false")
		}
	})

	t.Run("Handles are not equal (different handles)", func(t *testing.T) {
		if handle1.Equal(handle3) {
			t.Errorf("Expected handles to be different, got true")
		}
	})

	t.Run("Handle is closed", func(t *testing.T) {
		if closedHandle.Equal(handle1) {
			t.Errorf("Expected closed handle to be unequal, got true")
		}
	})

	t.Run("Handle is closed (both closed)", func(t *testing.T) {
		if !closedHandle.Equal(Handle(-1)) {
			t.Errorf("Expected two closed handles to be equal, got false")
		}
	})

	t.Run("Fstat fails", func(t *testing.T) {
		// Assuming unix.Fstat would fail for an invalid file descriptor (-1)
		if closedHandle.Equal(handle3) {
			t.Errorf("Expected unequal for closed handle and valid handle, got true")
		}
	})
}
