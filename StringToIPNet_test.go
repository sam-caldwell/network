package network

import (
	"net"
	"testing"
)

// TestStringToIPNet tests the StringToIPNet function.
func TestStringToIPNet(t *testing.T) {
	tests := []struct {
		cidr      string
		expected  net.IPNet
		expectErr bool
	}{
		{"192.168.1.0/24", net.IPNet{IP: net.ParseIP("192.168.1.0"), Mask: net.CIDRMask(24, 32)}, false},
		{"10.0.0.0/8", net.IPNet{IP: net.ParseIP("10.0.0.0"), Mask: net.CIDRMask(8, 32)}, false},
		{"2001:db8::/32", net.IPNet{IP: net.ParseIP("2001:db8::"), Mask: net.CIDRMask(32, 128)}, false},
		{"::1", net.IPNet{}, true},                // Invalid CIDR notation
		{"256.256.256.256/32", net.IPNet{}, true}, // Invalid IP address
		{"", net.IPNet{}, true},                   // Empty string
	}

	for _, test := range tests {
		t.Run(test.cidr, func(t *testing.T) {
			result, err := StringToIPNet(test.cidr)
			if (err != nil) != test.expectErr {
				t.Errorf("StringToIPNet(%q) returned error %v; expectErr=%v", test.cidr, err, test.expectErr)
				return
			}
			if result.IP.String() != test.expected.IP.String() || !CompareIPMask(result.Mask, test.expected.Mask) {
				t.Errorf("StringToIPNet(%q) = %v; want %v", test.cidr, result, test.expected)
			}
		})
	}
}
