package namespace

import (
	"github.com/sam-caldwell/network/test"
	"golang.org/x/sys/unix"
	"os"
	"path/filepath"
	"testing"
)

func TestOpsysInterfaceFunctions(t *testing.T) {
	t.Run("OpsysGet()", func(t *testing.T) {
		handle, err := OpsysGet()
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
	t.Run("OpsysGetFromName()", func(t *testing.T) {
		const testDockerImage = "network-test:latest"

		t.Run("build container", func(t *testing.T) {
			if err := test.BuildTestContainer(testDockerImage); err != nil {
				t.Fatal(err)
			}
		})

		t.Run("run test container", func(t *testing.T) {
			if err := test.RunContainer(testDockerImage, "TestCreateGetFromName"); err != nil {
				t.Fatal(err)
			}
		})
	})

	t.Run("OpsysGetFromPath()", func(t *testing.T) {
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
			handle, err = OpsysGetFromPath(tmpFile)
			if err != nil {
				t.Fatalf("Expected no error when opening a valid path. Got: %v", err)
			}
			if handle == -1 {
				t.Fatalf("Expected handle to not be -1. Got: %v", handle)
			}
		})
		t.Run("test with invalid path", func(t *testing.T) {
			invalidPath := "/non/existent/path"
			handle, err := OpsysGetFromPath(invalidPath)
			if err == nil {
				t.Fatalf("Expected an error when opening a invalid path. Got: %v", handle)
			}
			if handle != -1 {
				t.Fatalf("Expected handle to be -1. Got: %v", handle)
			}
		})
	})
	t.Run("OpsysGetFromPid()", func(t *testing.T) {
		// checks if GetFromPid returns a valid handle for the current process.
		t.Run("checks if GetFromPid returns a valid handle for the current process", func(t *testing.T) {
			// Get the PID of the current process
			pid := os.Getpid()

			// Call GetFromPid
			handle, err := OpsysGetFromPid(pid)
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
			handle, err := OpsysGetFromPid(invalidPid)
			if err == nil {
				t.Fatalf("expected an error")
			}
			if handle != -1 {
				t.Fatalf("expected handle to be -1")
			}
		})
	})
	t.Run("GetFromThread()", func(t *testing.T) {
		t.Run("checks if GetFromThread returns a valid handle for the current thread", func(t *testing.T) {
			pid := os.Getpid()
			tid := unix.Gettid()

			handle, err := OpsysGetFromThread(pid, tid)
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

			handle, err := OpsysGetFromThread(pid, invalidTid)
			if err == nil {
				t.Fatalf("expected an error")
			}
			if handle != -1 {
				t.Fatalf("expected handle to be -1")
			}
		})
	})
	t.Run("GetNewSetDelete test", func(t *testing.T) {
		const testDockerImage = "network-test:latest"

		t.Run("build container", func(t *testing.T) {
			if err := test.BuildTestContainer(testDockerImage); err != nil {
				t.Fatal(err)
			}
		})
		t.Run("run test container", func(t *testing.T) {
			if err := test.RunContainer(testDockerImage, "TestCreateAndDeleteNamespace"); err != nil {
				t.Fatal(err)
			}
		})
	})
}
