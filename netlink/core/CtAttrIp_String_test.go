package core

import (
	"testing"
)

func TestCtAttrIp_String(t *testing.T) {
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
}
