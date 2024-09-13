package core

import (
	"golang.org/x/sys/unix"
	"testing"
)

// TestGetNetlinkSocket tests the GetNetlinkSocket function.
func TestGetNetlinkSocket(t *testing.T) {
	tests := []struct {
		name     string
		protocol IpProtocol
		wantErr  bool
	}{
		{
			name:     "Valid protocol",
			protocol: unix.NETLINK_ROUTE, // Example valid protocol
			wantErr:  false,
		},
	}

	t.Skip("disabled")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sock, err := GetNetlinkSocket(tt.protocol)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNetlinkSocket() error = %v, wantErr %v", err, tt.wantErr)
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
