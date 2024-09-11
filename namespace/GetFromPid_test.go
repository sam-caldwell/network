package namespace

import (
	"golang.org/x/sys/unix"
	"os"
	"testing"
)

func TestGetFromPid(t *testing.T) {
	// checks if GetFromPid returns a valid handle for the current process.
	t.Run("checks if GetFromPid returns a valid handle for the current process", func(t *testing.T) {
		// Get the PID of the current process
		pid := os.Getpid()

		// Call GetFromPid
		handle, err := GetFromPid(pid)
		if err != nil {
			t.Fatalf("expected no error. got: %v", err)
		}
		if handle == -1 {
			t.Fatalf("expected handle to not be -1")
		}

		// Check if the file descriptor is valid by getting file status
		var stat unix.Stat_t
		err = unix.Fstat(int(handle), &stat)
		if err != nil {
			t.Fatalf("expected no error. got: %v", err)
		}
		// Clean up by closing the handle
		if err = unix.Close(int(handle)); err != nil {
			t.Fatalf("Expected no error when closing the file descriptor. Got: %v", err)
		}
	})

	t.Run("checks if GetFromPid returns an error for an invalid PID", func(t *testing.T) {
		invalidPid := -1

		// Call GetFromPid and expect an error
		handle, err := GetFromPid(invalidPid)
		if err == nil {
			t.Fatalf("expected an error")
		}
		if handle != -1 {
			t.Fatalf("expected handle to be -1")
		}
	})
}
