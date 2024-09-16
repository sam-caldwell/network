package core

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestIfAddressMessage_Serialize(t *testing.T) {
	var (
		err          error
		testData     IfAddressMessage
		expectedData []byte
		actualData   []byte
	)
	t.Run("setup testData", func(t *testing.T) {
		testData.Family = AfInet
		testData.Prefixlen = uint8(0x10)
		testData.Flags = uint8(0x11)
		testData.Scope = uint8(0x12)
		testData.Index = binary.NativeEndian.Uint32([]byte{0x23, 0x22, 0x21, 0x20})
	})
	t.Run("setup expectedData", func(t *testing.T) {
		expectedData = []byte{
			0x02,
			0x10,
			0x11,
			0x12,
			0x23, 0x22, 0x21, 0x20,
		}
	})
	t.Run("run serialize method", func(t *testing.T) {
		actualData, err = testData.Serialize()
		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("compare outcome", func(t *testing.T) {
		if !bytes.Equal(actualData, expectedData) {
			t.Fatalf("serialize method failed.\n"+
				"  actual:%v\n"+
				"expected:%v",
				actualData, expectedData)
		}
	})
}
