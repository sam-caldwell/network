package core

import (
	"testing"
	"unsafe"
)

func TestDevlinkParamCmodeEnum(t *testing.T) {
	t.Run("size check", func(t *testing.T) {
		const expectedSizeInBytes = 1
		if sz := int(unsafe.Sizeof(DevlinkParamCmode(0))); sz != expectedSizeInBytes {
			t.Fatal("size check failed")
		}
	})
	t.Run("value check", func(t *testing.T) {
		type TestData struct {
			actual DevlinkParamCmode
			expect DevlinkParamCmode
		}
		testData := []TestData{
			{actual: DevlinkParamCmodeRuntime, expect: 0},
			{actual: DevlinkParamCmodeDriverinit, expect: 1},
			{actual: DevlinkParamCmodePermanent, expect: 2},
			{actual: DevlinkParamCmodeMax, expect: DevlinkParamCmodePermanent},
		}
		for _, v := range testData {
			if v.actual != v.expect {
				t.Fatalf("value check failed. actual: %v, expect: %v", v.actual, v.expect)
			}
		}
	})
}
