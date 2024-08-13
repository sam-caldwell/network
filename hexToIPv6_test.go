package network

import (
	"net"
	"testing"
)

func TestHexToIPv6(t *testing.T) {
	tests := []struct {
		hexStr    string
		expected  string
		expectErr bool
	}{
		{"00000000000000000000000000000001", "::1", false},
		{"20010db8000000000000000000000000", "2001:db8::", false},
		{"fe800000000000000200000000000001", "fe80::0200:0:0:1", false},
		{"fe800000000000000002000000000001", "fe80::2:0:0:1", false},

		{"ffffffffffffffffffffffffffffffff", "ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff", false},
		{"00000000000000000000000000000000", "::", false},
		{"fe800000000000001fa804cd82105e87", "fe80::1fa8:4cd:8210:5e87", false},
		{"invalid hex string", "", true},
	}
	for testNo, test := range tests {
		t.Run(test.hexStr, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil && !test.expectErr {
					t.Errorf("hexToIPv6(%q) panicked unexpectedly: %v", test.hexStr, r)
				}
			}()
			result := hexToIPv6(test.hexStr)

			if !test.expectErr {
				//No error
				actualIP := net.ParseIP(test.expected).To16()
				if actualIP == nil {
					t.Fatalf("invalid nil. test data(%d):\n"+
						"  actual: %v\n"+
						"expected: %v", testNo, actualIP, test.expected)
				}
				resultIP := net.ParseIP(result).To16()
				if resultIP == nil {
					t.Fatalf("invalid nil test data(%d):\n"+
						"  actual: %v\n"+
						"expected: %v", testNo, actualIP, test.expected)
				}
				if !actualIP.Equal(resultIP) {
					t.Errorf("hexToIPv6(%q) mismatch\n"+
						"  actual:   %s\n"+
						"  expected: %s",
						test.hexStr, resultIP.String(), actualIP.String())
				}
			} else if result != "" {
				t.Errorf("hexToIPv6(%q) = %v; want error", test.hexStr, result)
			}
		})
	}
}
