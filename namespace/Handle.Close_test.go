package namespace

import (
	"os"
	"testing"
)

func TestHandle_Close(t *testing.T) {
	t.Run("close a handle successfully", func(t *testing.T) {
		handle, err := GetFromPid(os.Getpid())
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
		handle := Handle(closedHandle)
		err := handle.Close()
		if err == nil {
			t.Fatal("expected error when closing closed handle")
		}
		if err != nil && err.Error() != "handle is closed" {
			t.Fatalf("expected no error. got: %v", err)
		}
	})
}
