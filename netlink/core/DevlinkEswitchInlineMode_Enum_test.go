package core

import (
	"testing"
	"unsafe"
)

func TestDevlinkEswitchInlineModeEnum(t *testing.T) {
	t.Run("size check", func(t *testing.T) {
		const expectedSizeInBytes = 1
		if sz := int(unsafe.Sizeof(DevlinkEswitchInlineMode(0))); sz != expectedSizeInBytes {
			t.Fatal("size check failed")
		}
	})
	t.Run("value check", func(t *testing.T) {
		type TestData struct {
			actual DevlinkEswitchInlineMode
			expect DevlinkEswitchInlineMode
		}
		testData := []TestData{
			{actual: DevlinkEswitchInlineModeNone, expect: 0},
			{actual: DevlinkEswitchInlineModeLink, expect: 1},
			{actual: DevlinkEswitchInlineModeNetwork, expect: 2},
			{actual: DevlinkEswitchInlineModeTransport, expect: 3},
		}
		for _, v := range testData {
			if v.actual != v.expect {
				t.Fatalf("value check failed. actual: %v, expect: %v", v.actual, v.expect)
			}
		}
	})
}
