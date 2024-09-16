package core

import (
	"testing"
	"unsafe"
)

func TestCtAttrTupleEnum(t *testing.T) {
	t.Run("size check", func(t *testing.T) {
		const expectedSizeInBytes = 1
		if sz := unsafe.Sizeof(CtAttrTuple(0)); sz != expectedSizeInBytes {
			t.Fatal("size check failed")
		}
	})
	t.Run("value check", func(t *testing.T) {
		type TestData struct {
			actual   CtAttrTuple
			expected CtAttrTuple
		}
		testData := []TestData{
			{actual: 0, expected: CtaTupleUnspec},
			{actual: 1, expected: CtaTupleIp},
			{actual: 2, expected: CtaTupleProto},
			{actual: 3, expected: CtaTupleZone},
			{actual: 3, expected: CtaTupleMax},
		}
		for _, v := range testData {
			if v.actual != v.expected {
				t.Fatalf("value check failed.  actual: %d, expected: %d", v.actual, v.expected)
			}
		}
	})
}
