package network

import (
	"net"
	"testing"
)

func TestStringToIPMask(t *testing.T) {
	tests := []struct {
		maskStr   string
		expected  net.IPMask
		expectErr bool
	}{
		{"192.168.1.0/24", net.CIDRMask(24, 32), false}, // CIDR notation
		{"255.255.255.0", net.CIDRMask(24, 32), false},  // Dotted-decimal notation
		{"2001:db8::/32", net.CIDRMask(32, 128), false}, // CIDR notation for IPv6
		{"::1", net.CIDRMask(128, 128), false},          // IPv6 address with full mask
		{"255.0.0.0", net.CIDRMask(8, 32), false},       // Dotted-decimal notation
		{"256.256.256.256", nil, true},                  // Invalid IP address
		{"", nil, true},                                 // Empty string
	}

	for _, test := range tests {
		t.Run(test.maskStr, func(t *testing.T) {
			result, err := StringToIPMask(test.maskStr)
			if (err != nil) != test.expectErr {
				t.Errorf("StringToIPMask(%q) returned error %v; expectErr=%v", test.maskStr, err, test.expectErr)
				return
			}
			if !CompareIPMask(result, test.expected) {
				t.Errorf("StringToIPMask(%q) = %v; want %v", test.maskStr, result, test.expected)
			}
		})
	}
}
