package namespace

import (
	"github.com/sam-caldwell/network/test"
	"golang.org/x/sys/unix"
	"math"
	"os"
	"path/filepath"
	"testing"
	"unsafe"
)

func TestHandle_Type(t *testing.T) {

	t.Run("test Handle type size", func(t *testing.T) {
		const expectedSizeInBytes = int(unsafe.Sizeof(int(0)))
		if sz := int(unsafe.Sizeof(Handle(0))); sz != expectedSizeInBytes {
			t.Errorf("Expected size of %d, got %d", expectedSizeInBytes, sz)
		}
		_ = Handle(math.MinInt)
		_ = Handle(math.MaxInt)
	})

	t.Run("Test Handle.NewHandle() method", func(t *testing.T) {
		const testDockerImage = "network-test:latest"

		t.Run("build container", func(t *testing.T) {
			if err := test.BuildTestContainer(testDockerImage); err != nil {
				t.Fatal(err)
			}
		})

		t.Run("run test container", func(t *testing.T) {
			if err := test.RunContainer(testDockerImage, "TestCreateNewNamespace"); err != nil {
				t.Fatal(err)
			}
		})
	})

	t.Run("Test Handle.String() method", func(t *testing.T) {
		startsWith := func(s, prefix string) bool {
			return len(s) >= len(prefix) && s[:len(prefix)] == prefix
		}
		t.Run("handle with valid device and inode", func(t *testing.T) {
			handle, err := Opsys.Get() // Assumes Get() provides a valid handle for the current thread's namespace
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
			var closedHandle Handle = -1
			expected := "NS(none)"
			if got := closedHandle.String(); got != expected {
				t.Fatalf("expected %s, got %s", expected, got)
			}
		})
	})

	t.Run("Test Handle.UniqueID() method", func(t *testing.T) {
		startsWith := func(s, prefix string) bool {
			return len(s) >= len(prefix) && s[:len(prefix)] == prefix
		}
		t.Run("valid namespace handle", func(t *testing.T) {
			handle, err := Opsys.Get() // Assumes Get() provides a valid handle for the current thread's namespace
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
	})

	t.Run("Test None() method", func(t *testing.T) {
		if None() != closedHandle {
			t.Fatal("expected -1")
		}
	})

	t.Run("Test IsOpen() method", func(t *testing.T) {
		t.Run("IsOpen(1) should be true", func(t *testing.T) {
			if namespace := Handle(1); !namespace.IsOpen() {
				t.Fatal("expected true")
			}
		})
		t.Run("IsOpen(-1) should be false", func(t *testing.T) {
			if namespace := Handle(-1); namespace.IsOpen() {
				t.Fatal("expected false")
			}
		})
		t.Run("IsOpen(closedHandle) should be false", func(t *testing.T) {
			if namespace := Handle(closedHandle); namespace.IsOpen() {
				t.Fatal("expected false")
			}
		})
	})

	t.Run("Test Close() method", func(t *testing.T) {
		t.Run("close a handle successfully", func(t *testing.T) {
			handle, err := Opsys.GetFromPid(os.Getpid())
			if err != nil {
				t.Fatalf("expected no error. got: %v", err)
			}
			if err := handle.Close(); err != nil {
				t.Fatalf("expected no error. got: %v", err)
			}
			if handle != closedHandle {
				t.Fatalf("expected handle to be %v, got %v", closedHandle, handle)
			}
		})

		t.Run("close an already closed handle", func(t *testing.T) {
			handle, err := Opsys.GetFromPid(os.Getpid())
			if err != nil {
				t.Fatalf("expected no error. got: %v", err)
			}
			if err := handle.Close(); err != nil {
				t.Fatalf("expected no error. got: %v", err)
			}
			if err := handle.Close(); err == nil {
				t.Fatal("expected an error when closing an already closed handle")
			}
		})
	})

	t.Run("Test Equal() method", func(t *testing.T) {
		t.Run("handles referring to the same namespace", func(t *testing.T) {
			var (
				err              error
				handle1, handle2 Handle
			)
			if handle1, err = Opsys.GetFromPid(os.Getpid()); err != nil {
				t.Fatalf("expected no error. got: %v", err)
			}
			if handle2, err = Opsys.GetFromPid(os.Getpid()); err != nil {
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

			if handle1, err = Opsys.GetFromPid(os.Getpid()); err != nil {
				t.Fatalf("expected no error. got: %v", err)
			}
			// Create a temporary file to simulate the network namespace file
			tmpFile := filepath.Join(tmpDir, "test_ns")
			if f, err = os.Create(tmpFile); err != nil {
				t.Fatalf("Failed to create temp file: %v", err)
			}
			// Now call GetFromPath
			handle2, err = Opsys.GetFromPath(tmpFile)
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
	})
}
