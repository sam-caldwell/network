package core

import (
	"golang.org/x/sys/unix"
	"testing"
)

// TestGetNetlinkSocket tests the getNetlinkSocket function.
func TestGetNetlinkSocket(t *testing.T) {
	tests := []struct {
		name     string
		protocol int
		wantErr  bool
	}{
		{
			name:     "Valid protocol",
			protocol: unix.NETLINK_ROUTE, // Example valid protocol
			wantErr:  false,
		},
		{
			name:     "Invalid protocol",
			protocol: -1, // Example invalid protocol
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sock, err := getNetlinkSocket(tt.protocol)
			if (err != nil) != tt.wantErr {
				t.Errorf("getNetlinkSocket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && sock == nil {
				t.Errorf("Expected non-nil socket, got nil")
			}
			if err == nil && sock.fd == 0 {
				t.Errorf("Expected valid file descriptor, got 0")
			}
		})
	}
}
