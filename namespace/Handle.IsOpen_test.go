package namespace

import (
	"testing"
)

func TestHandle_IsOpen(t *testing.T) {
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
}
