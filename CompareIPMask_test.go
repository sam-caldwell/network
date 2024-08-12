package network

import (
	"net"
	"testing"
)

// TestCompareIPMask tests the CompareIPMask function.
func TestCompareIPMask(t *testing.T) {
	tests := []struct {
		a, b     net.IPMask
		expected bool
	}{
		{net.CIDRMask(24, 32), net.CIDRMask(24, 32), true},  // Identical masks should be equal
		{net.CIDRMask(24, 32), net.CIDRMask(16, 32), false}, // Different masks should not be equal
		{nil, nil, true},                   // Both nil should be equal
		{nil, net.CIDRMask(24, 32), false}, // Nil vs non-nil should not be equal
		{net.CIDRMask(24, 32), nil, false}, // Non-nil vs nil should not be equal
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := CompareIPMask(test.a, test.b)
			if result != test.expected {
				t.Errorf("CompareIPMask(%v, %v) = %v; want %v", test.a, test.b, result, test.expected)
			}
		})
	}
}
