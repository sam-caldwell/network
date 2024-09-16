package core

import (
	"testing"
	"unsafe"
)

func TestDevlinkParamTypeEnum(t *testing.T) {
	t.Run("size check", func(t *testing.T) {
		const expectedSizeInBytes = 1
		if sz := int(unsafe.Sizeof(DevlinkParamType(0))); sz != expectedSizeInBytes {
			t.Fatal("size check failed")
		}
	})
	t.Run("value check", func(t *testing.T) {
		type TestData struct {
			actual DevlinkParamType
			expect DevlinkParamType
		}
		testData := []TestData{
			{actual: DevlinkParamTypeU8, expect: 1},
			{actual: DevlinkParamTypeU16, expect: 2},
			{actual: DevlinkParamTypeU32, expect: 3},
			{actual: DevlinkParamTypeString, expect: 5},
			{actual: DevlinkParamTypeBool, expect: 6},
		}
		for _, v := range testData {
			if v.actual != v.expect {
				t.Fatalf("value check failed. actual: %v, expect: %v", v.actual, v.expect)
			}
		}
	})
}
