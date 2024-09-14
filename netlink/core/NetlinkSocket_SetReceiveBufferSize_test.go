package core

import (
	"golang.org/x/sys/unix"
	"testing"
)

func TestNetlinkSocket_SetReceiveBufferSize(t *testing.T) {
	tests := []struct {
		name        string
		size        int
		force       bool
		expectedErr error
	}{
		{
			name:        "set receive buffer size",
			size:        4096,
			force:       false,
			expectedErr: nil,
		},
		{
			name:        "force set receive buffer size",
			size:        8192,
			force:       true,
			expectedErr: nil,
		},
	}

	t.Skip("disabled.  move to docker-based test")

	// Create a new socket for testing
	fd, err := unix.Socket(unix.AF_NETLINK, unix.SOCK_RAW, unix.NETLINK_GENERIC)
	if err != nil {
		t.Fatalf("failed to create socket: %v", err)
	}
	defer func() { _ = unix.Close(fd) }()

	mockSocket := &NetlinkSocket{fd: int32(fd)}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := mockSocket.SetReceiveBufferSize(tt.size, tt.force)

			// Check for errors
			if (err != nil) != (tt.expectedErr != nil) {
				t.Errorf("SetReceiveBufferSize() error = %v, expectedErr %v", err, tt.expectedErr)
			}
		})
	}
}
