package namespace

import (
	"golang.org/x/sys/unix"
	"os"
	"path/filepath"
	"testing"
)

func TestHandle_Equal(t *testing.T) {
	t.Run("handles referring to the same namespace", func(t *testing.T) {
		var (
			err              error
			handle1, handle2 Handle
		)
		if handle1, err = GetFromPid(os.Getpid()); err != nil {
			t.Fatalf("expected no error. got: %v", err)
		}
		if handle2, err = GetFromPid(os.Getpid()); err != nil {
			t.Fatalf("expected no error. got: %v", err)
		}
		if !handle1.Equal(handle2) {
			t.Fatalf("expected handles to be equal")
		}
	})

	t.Run("handles referring to different namespaces", func(t *testing.T) {
		var (
			f                *os.File
			err              error
			handle1, handle2 Handle
		)
		defer func() {
			if f != nil {
				if err = f.Close(); err != nil {
					t.Fatalf("error closing temp file: %v", err)
				}
			}
			if handle1 == -1 {
				return
			}
			if err = unix.Close(int(handle1)); err != nil {
				t.Fatalf("error closing unix socket: %v", err)
			}
			if handle2 == -1 {
				return
			}
			if err = unix.Close(int(handle2)); err != nil {
				t.Fatalf("error closing unix socket: %v", err)
			}
		}()
		// Create a temporary directory
		tmpDir := t.TempDir()

		if handle1, err = GetFromPid(os.Getpid()); err != nil {
			t.Fatalf("expected no error. got: %v", err)
		}

		// Create a temporary file to simulate the network namespace file
		tmpFile := filepath.Join(tmpDir, "test_ns")
		if f, err = os.Create(tmpFile); err != nil {
			t.Fatalf("Failed to create temp file: %v", err)
		}

		// Now call GetFromPath
		handle2, err = GetFromPath(tmpFile)
		if err != nil {
			t.Fatalf("Expected no error when opening a valid path. Got: %v", err)
		}
		if handle2 == -1 {
			t.Fatalf("Expected handle to not be -1. Got: %v", handle2)
		}

		// Assuming you have a way to get a handle for a different namespace or process
		if handle1.Equal(handle2) {
			t.Fatalf("expected handles to be different")
		}
	})
}
