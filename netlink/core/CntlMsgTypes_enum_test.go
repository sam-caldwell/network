package core

import (
	"testing"
	"unsafe"
)

func TestCntlMsgTypes_enum(t *testing.T) {
	t.Run("test size of CntlMsgTypes", func(t *testing.T) {
		const expectedSizeInBytes = 1
		if sz := int(unsafe.Sizeof(CntlMsgTypes(0))); sz != expectedSizeInBytes {
			t.Fatal("size mismatch")
		}
	})
	t.Run("value check", func(t *testing.T) {
		testData := map[CntlMsgTypes]CntlMsgTypes{
			0: IpCtnlMsgCtNew,
			1: IpCtnlMsgCtGet,
			2: IpCtnlMsgCtDelete,
			3: IpCtnlMsgCtGetCtrZero,
			4: IpCtnlMsgCtGetStatsCpu,
			5: IpCtnlMsgCtGetStats,
			6: IpCtnlMsgCtGetDying,
			7: IpCtnlMsgCtGetUnconfirmed,
		}
		for actual, expected := range testData {
			if actual != expected {
				t.Fatalf("value mismatch.  Actual:%d Expected %d", actual, expected)
			}
		}
		if IpCtnlMsgMax != IpCtnlMsgCtGetUnconfirmed {
			t.Fatal("max value mismatch")
		}
	})
}
