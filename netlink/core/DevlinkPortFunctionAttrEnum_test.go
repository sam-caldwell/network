package core

import (
	"testing"
	"unsafe"
)

func TestDevlinkPortFunctionAttrEnum(t *testing.T) {
	t.Run("size check", func(t *testing.T) {
		const expectedSizeInBytes = 1
		if sz := int(unsafe.Sizeof(DevlinkParamTypeEnum(0))); sz != expectedSizeInBytes {
			t.Fatal("size check failed")
		}
	})
	t.Run("value check", func(t *testing.T) {
		type TestData struct {
			actual DevlinkPortFunctionAttrEnum
			expect DevlinkPortFunctionAttrEnum
		}
		testData := []TestData{
			{actual: DevlinkPortFunctionAttrUnspec, expect: 0},
			{actual: DevlinkPortFunctionAttrHwAddr, expect: 1},
			{actual: DevlinkPortFnAttrState, expect: 2},
			{actual: DevlinkPortFnAttrOpstate, expect: 3},
			{actual: DevlinkPortFnAttrCaps, expect: 4},
			{actual: DevlinkPortFnAttrDevlink, expect: 5},
			{actual: DevlinkPortFnAttrMaxIoEqs, expect: 6},
			{actual: DevlinkPortFunctionAttrMax, expect: 6},
		}
		for _, v := range testData {
			if v.actual != v.expect {
				t.Fatalf("value check failed. actual: %v, expect: %v", v.actual, v.expect)
			}
		}
	})
}
