package namespace

import (
	"math"
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

}
