//go:build linux

package network

import (
	"bytes"
	"fmt"
	"golang.org/x/sys/unix"
)

// sendMessage - send the netlink message
func sendMessage(socketFileDescriptor int, buffer *bytes.Buffer) error {

	if err := unix.Sendto(socketFileDescriptor, buffer.Bytes(), 0, &unix.SockaddrNetlink{Family: unix.AF_NETLINK}); err != nil {
		return fmt.Errorf("sendto error: %v", err)
	}

	return nil

}
