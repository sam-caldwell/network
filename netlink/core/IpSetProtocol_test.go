package core

import (
	"testing"
)

func TestIpSetConstants(t *testing.T) {
	tests := []struct {
		name     string
		value    int
		expected int
	}{
		{"IpSetProtocol", IpSetProtocol, 7},
		{"IpSetProtocolMin", IpSetProtocolMin, 6},
		{"IpSetMaxnamelen", IpSetMaxnamelen, 32},
		{"IpSetMaxCommentSize", IpSetMaxCommentSize, 255},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.value != test.expected {
				t.Errorf("Expected %d but got %d for %s", test.expected, test.value, test.name)
			}
		})
	}
}
