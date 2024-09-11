package namespace

import (
	"math"
	"testing"
)

func TestVariables(t *testing.T) {
	t.Run("verify locking", func(t *testing.T) {
		lock.Lock()
		lock.Unlock()
	})
	t.Run("verify Handle", func(t *testing.T) {
		_ = Handle(math.MinInt)
		_ = Handle(0)
		_ = Handle(math.MaxInt)
	})
	t.Run("namespace map", func(t *testing.T) {
		namespaces["test1"] = Handle(1)
		namespaces["test2"] = Handle(2)
		delete(namespaces, "test1")
	})
}
