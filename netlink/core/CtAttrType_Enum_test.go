package core

import (
	"testing"
	"unsafe"
)

func TestCtAttrTypeEnum(t *testing.T) {
	t.Run("size check", func(t *testing.T) {
		const expectedSizeInBytes = 1
		if sz := unsafe.Sizeof(CtAttrType(0)); sz != expectedSizeInBytes {
			t.Fatalf("size check failed.  actual: %d, expected: %d", sz, expectedSizeInBytes)
		}
	})
	t.Run("value check", func(t *testing.T) {
		type TestData struct {
			actual   CtAttrType
			expected CtAttrType
		}
		testData := []TestData{
			{actual: 0, expected: CtaUnspec},
			{actual: 1, expected: CtaTupleOrig},
			{actual: 2, expected: CtaTupleReply},
			{actual: 3, expected: CtaStatus},
			{actual: 4, expected: CtaProtoInfo},
			{actual: 5, expected: CtaHelp},
			{actual: 6, expected: CtaNatSrc},
			{actual: 6, expected: CtaNat},
			{actual: 7, expected: CtaTimeout},
			{actual: 8, expected: CtaMark},
			{actual: 9, expected: CtaCountersOrig},
			{actual: 10, expected: CtaCountersReply},
			{actual: 11, expected: CtaUse},
			{actual: 12, expected: CtaId},
			{actual: 13, expected: CtaNatDst},
			{actual: 14, expected: CtaTupleMaster},
			{actual: 15, expected: CtaSeqAdjOrig},
			{actual: 15, expected: CtaNatSeqAdjOrig},
			{actual: 16, expected: CtaSeqAdjReply},
			{actual: 16, expected: CtaNatSeqAdjReply},
			{actual: 17, expected: CtaSecmark},
			{actual: 18, expected: CtaZone},
			{actual: 19, expected: CtaSeccTx},
			{actual: 20, expected: CtaTimestamp},
			{actual: 21, expected: CtaMarkMask},
			{actual: 22, expected: CtaLabels},
			{actual: 23, expected: CtaLabelsMask},
			{actual: 24, expected: CtaSynProxy},
			{actual: 25, expected: CtaFilter},
			{actual: 26, expected: CtaStatusMask},
			{actual: 26, expected: CtaMax},
		}
		for _, v := range testData {
			if v.actual != v.expected {
				t.Fatalf("value check mismatch.  Actual: %d, Expected: %d", v.actual, v.expected)
			}
		}
	})
}
