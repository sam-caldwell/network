//go:build linux

package network

import (
	"golang.org/x/sys/unix"
)

// closeNetlinkRouteSocket - close the file descriptor
func closeNetlinkRouteSocket(fileDescriptor int) {
	_ = unix.Close(fileDescriptor)
}
