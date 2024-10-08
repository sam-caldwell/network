package core

import (
	"golang.org/x/sys/unix"
	"testing"
)

func TestNetlinkSocket_SetExtAck(t *testing.T) {
	tests := []struct {
		name        string
		enable      bool
		expectedErr error
	}{
		{
			name:        "enable ext ack",
			enable:      true,
			expectedErr: nil,
		},
		{
			name:        "disable ext ack",
			enable:      false,
			expectedErr: nil,
		},
	}

	// Create a new socket for testing
	fd, err := unix.Socket(unix.AF_NETLINK, unix.SOCK_RAW, unix.NETLINK_GENERIC)
	if err != nil {
		t.Fatalf("failed to create socket: %v", err)
	}
	defer func() { _ = unix.Close(fd) }()

	mockSocket := &MockNetlinkSocket{fd: fd}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := mockSocket.SetExtAck(tt.enable)

			// Check for errors
			if (err != nil) != (tt.expectedErr != nil) {
				t.Errorf("SetExtAck() error = %v, expectedErr %v", err, tt.expectedErr)
			}
		})
	}
}
