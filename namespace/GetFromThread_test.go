package namespace

import (
	"golang.org/x/sys/unix"
	"os"
	"testing"
)

func TestGetFromThread(t *testing.T) {
	t.Run("checks if GetFromThread returns a valid handle for the current thread", func(t *testing.T) {
		pid := os.Getpid()
		tid := unix.Gettid()

		handle, err := GetFromThread(pid, tid)
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
		if err = unix.Close(int(handle)); err != nil {
			t.Fatalf("expected no error when closing the file descriptor. got: %v", err)
		}
	})

	t.Run("checks if GetFromThread returns an error for an invalid TID", func(t *testing.T) {
		invalidTid := -1
		pid := os.Getpid()

		handle, err := GetFromThread(pid, invalidTid)
		if err == nil {
			t.Fatalf("expected an error")
		}
		if handle != -1 {
			t.Fatalf("expected handle to be -1")
		}
	})
}
