package namespace

import (
	"golang.org/x/sys/unix"
	"os"
	"path/filepath"
	"testing"
)

func TestGetFromPath(t *testing.T) {
	// TestGetFromPath verifies that GetFromPath correctly opens a file and returns a handle.
	t.Run("Happy path", func(t *testing.T) {
		var (
			f      *os.File
			err    error
			handle Handle
		)
		defer func() {
			if f != nil {
				if err = f.Close(); err != nil {
					t.Fatalf("error closing temp file: %v", err)
				}
			}
			if handle == -1 {
				return
			}
			if err = unix.Close(int(handle)); err != nil {
				t.Fatalf("error closing unix socket: %v", err)
			}
		}()
		// Create a temporary directory
		tmpDir := t.TempDir()
		// Create a temporary file to simulate the network namespace file
		tmpFile := filepath.Join(tmpDir, "test_ns")
		if f, err = os.Create(tmpFile); err != nil {
			t.Fatalf("Failed to create temp file: %v", err)
		}
		// Now call GetFromPath
		handle, err = GetFromPath(tmpFile)
		if err != nil {
			t.Fatalf("Expected no error when opening a valid path. Got: %v", err)
		}
		if handle == -1 {
			t.Fatalf("Expected handle to not be -1. Got: %v", handle)
		}
	})
	t.Run("test with invalid path", func(t *testing.T) {
		invalidPath := "/non/existent/path"
		handle, err := GetFromPath(invalidPath)
		if err == nil {
			t.Fatalf("Expected an error when opening a invalid path. Got: %v", handle)
		}
		if handle != -1 {
			t.Fatalf("Expected handle to be -1. Got: %v", handle)
		}
	})
}
