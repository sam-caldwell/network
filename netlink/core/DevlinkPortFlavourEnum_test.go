package core

import (
	"testing"
	"unsafe"
)

func TestDevlinkPortFlavourEnum(t *testing.T) {
	t.Run("size check", func(t *testing.T) {
		const expectedSizeInBytes = 1
		if sz := int(unsafe.Sizeof(DevlinkParamTypeEnum(0))); sz != expectedSizeInBytes {
			t.Fatal("size check failed")
		}
	})
	t.Run("value check", func(t *testing.T) {
		type TestData struct {
			actual DevlinkPortFlavourEnum
			expect DevlinkPortFlavourEnum
		}
		testData := []TestData{
			{actual: DevlinkPortFlavourPhysical, expect: 0},
			{actual: DevlinkPortFlavourCpu, expect: 1},
			{actual: DevlinkPortFlavourDsa, expect: 2},
			{actual: DevlinkPortFlavourPciPf, expect: 3},
			{actual: DevlinkPortFlavourPciVf, expect: 4},
			{actual: DevlinkPortFlavourVirtual, expect: 5},
			{actual: DevlinkPortFlavourUnused, expect: 6},
			{actual: DevlinkPortFlavourPciSf, expect: 7},
		}
		for _, v := range testData {
			if v.actual != v.expect {
				t.Fatalf("value check failed. actual: %v, expect: %v", v.actual, v.expect)
			}
		}
	})
}
