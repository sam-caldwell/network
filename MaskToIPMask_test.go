package network

import (
	"net"
	"testing"
)

// TestMaskToIPMask tests the maskToIPMask function.
func TestMaskToIPMask(t *testing.T) {
	tests := []struct {
		maskStr  string
		expected net.IPMask
	}{
		{"255.255.255.0", net.CIDRMask(24, 32)},   // 255.255.255.0 should return a /24 mask
		{"255.0.0.0", net.CIDRMask(8, 32)},        // 255.0.0.0 should return a /8 mask
		{"0.0.0.0", net.CIDRMask(0, 32)},          // 0.0.0.0 should return a /0 mask
		{"255.255.255.255", net.CIDRMask(32, 32)}, // 255.255.255.255 should return a /32 mask
	}

	for _, test := range tests {
		t.Run(test.maskStr, func(t *testing.T) {
			ip := net.ParseIP(test.maskStr)
			if ip == nil {
				t.Fatalf("failed to parse IP address: %s", test.maskStr)
			}
			result := MaskToIPMask(ip)
			if !CompareIPMask(result, test.expected) {
				t.Errorf("maskToIPMask(%q) = %v; want %v", test.maskStr, result, test.expected)
			}
		})
	}
}
