package core

import (
	"testing"
	"unsafe"
)

func TestDevlinkPortTypeEnum(t *testing.T) {
	t.Run("size check", func(t *testing.T) {
		const expectedSizeInBytes = 1
		if sz := int(unsafe.Sizeof(DevlinkPortType(0))); sz != expectedSizeInBytes {
			t.Fatal("size check failed")
		}
	})
	t.Run("value check", func(t *testing.T) {
		type TestData struct {
			actual DevlinkPortType
			expect DevlinkPortType
		}
		testData := []TestData{
			{actual: DevlinkPortTypeNotset, expect: 0},
			{actual: DevlinkPortTypeAuto, expect: 1},
			{actual: DevlinkPortTypeEth, expect: 2},
			{actual: DevlinkPortTypeIb, expect: 3},
		}
		for _, v := range testData {
			if v.actual != v.expect {
				t.Fatalf("value check failed. actual: %v, expect: %v", v.actual, v.expect)
			}
		}
	})
}
