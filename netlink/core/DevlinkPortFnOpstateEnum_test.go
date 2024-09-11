package core

import (
	"testing"
	"unsafe"
)

func TestDevlinkPortFnOpstateEnum(t *testing.T) {
	t.Run("size check", func(t *testing.T) {
		const expectedSizeInBytes = 1
		if sz := int(unsafe.Sizeof(DevlinkPortFnOpstateEnum(0))); sz != expectedSizeInBytes {
			t.Fatal("size check failed")
		}
	})
	t.Run("value check", func(t *testing.T) {
		type TestData struct {
			actual DevlinkPortFnOpstateEnum
			expect DevlinkPortFnOpstateEnum
		}
		testData := []TestData{
			{actual: DevlinkPortFnOpstateDetached, expect: 0},
			{actual: DevlinkPortFnOpstateAttached, expect: 1},
		}
		for _, v := range testData {
			if v.actual != v.expect {
				t.Fatalf("value check failed. actual: %v, expect: %v", v.actual, v.expect)
			}
		}
	})
}
