package network

import (
	"fmt"
	"net"
	"testing"
)

func TestRouteInfo_Add(t *testing.T) {
	tests := []struct {
		name          string
		route         RouteInfo
		expectedError error
	}{
		{
			name: "Valid IPv4 route",
			route: RouteInfo{
				Network: net.IPv4(192, 168, 1, 1),
			},
			expectedError: nil, // Assuming addIPv4Route() returns nil on success
		},
		{
			name: "Valid IPv6 route",
			route: RouteInfo{
				Network: net.ParseIP("2001:db8::1"),
			},
			expectedError: nil, // Assuming addIPv6Route() returns nil on success
		},
		{
			name: "Invalid IP version",
			route: RouteInfo{
				Network: net.ParseIP("::/128"), // Assuming /128 is part of the IP address for some reason
			},
			expectedError: fmt.Errorf("unsupported IP version"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Call the Add method
			err := test.route.Add()

			// Check if the error matches the expected error
			if err != nil && test.expectedError != nil && err.Error() != test.expectedError.Error() {
				t.Errorf("Add() error = %v, want %v", err, test.expectedError)
			} else if err == nil && test.expectedError != nil {
				t.Errorf("Add() error = %v, want %v", err, test.expectedError)
			}
		})
	}
}
