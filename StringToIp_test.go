package network

import (
	"net"
	"testing"
)

func TestStringToIP(t *testing.T) {
	tests := []struct {
		ipStr    string
		expected net.IP
	}{
		{"192.168.1.1", net.ParseIP("192.168.1.1")},
		{"2001:db8::1", net.ParseIP("2001:db8::1")},
		{"invalid_ip", nil},
		{"256.256.256.256", nil},
		{"", nil},
	}

	for _, test := range tests {
		t.Run(test.ipStr, func(t *testing.T) {
			result := StringToIP(test.ipStr)
			if !result.Equal(test.expected) {
				t.Errorf("StringToIP(%q) = %v; want %v", test.ipStr, result, test.expected)
			}
		})
	}
}
