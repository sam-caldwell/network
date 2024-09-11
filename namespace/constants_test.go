package namespace

import (
	"testing"
)

func TestConstants(t *testing.T) {
	t.Run("test with constant values", func(t *testing.T) {
		testData := map[any]any{
			bindMountPath:        "/run/netns",
			processNamespacePath: "/proc/%d/ns/net",
			threadNamespacePath:  "/proc/%d/task/%d/ns/net",
			closedHandle:         Handle(-1),
			ErrNotImplemented:    "not implemented",
		}
		for actual, expected := range testData {
			if actual != expected {
				t.Fatalf("value mismatch.  actual: %v, expected: %v", actual, expected)
			}
		}
	})
}
