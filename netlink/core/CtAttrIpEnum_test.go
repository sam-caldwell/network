package core

import (
	"testing"
	"unsafe"
)

func TestCtAttrIpEnum(t *testing.T) {
	t.Run("test type", func(t *testing.T) {
		t.Run("size check", func(t *testing.T) {
			const expectedSizeInBytes = 1
			var o CtAttrIpEnum
			if sz := unsafe.Sizeof(o); sz != expectedSizeInBytes {
				t.Fatalf("size mismatch")
			}
		})
		t.Run("value check", func(t *testing.T) {
			testData := map[int]CtAttrIpEnum{
				0: CtaIpUnspec,
				1: CtaIpV4Src,
				2: CtaIpV4Dst,
				3: CtaIpV6Src,
				4: CtaIpV6Dst,
			}
			for k, v := range testData {
				if k != int(v) {
					t.Fatalf("value mismatch.  Actual: %d, Expected: %d", k, v)
				}
			}
		})
	})
	t.Run("test .String() method", func(t *testing.T) {
		t.Run("happy test", func(t *testing.T) {
			tests := []struct {
				input    CtAttrIpEnum
				expected string
			}{
				{CtaIpUnspec, "CTA_IP_UNSPEC"},
				{CtaIpV4Src, "CTA_IP_V4_SRC"},
				{CtaIpV4Dst, "CTA_IP_V4_DST"},
				{CtaIpV6Src, "CTA_IP_V6_SRC"},
				{CtaIpV6Dst, "CTA_IP_V6_DST"},
			}

			for _, tt := range tests {
				t.Run(t.Name(), func(t *testing.T) {
					result := tt.input.String()
					if result != tt.expected {
						t.Errorf("CtAttrIpEnum.String(): for input %v, expected %q, got %q", tt.input, tt.expected, result)
					}
				})
			}
		})
		t.Run("sad path 255", func(t *testing.T) {
			// Test default case panic
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("Expected panic for unhandled default case, but no panic occurred")
				}
			}()
			var invalid CtAttrIpEnum = 255
			_ = invalid.String() // should panic
		})
	})
}
