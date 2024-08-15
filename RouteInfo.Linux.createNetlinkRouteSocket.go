//go:build linux

package network

import (
	"golang.org/x/sys/unix"
)

// createNetlinkRouteSocket - Create a NetLink Route Socket
func createNetlinkRouteSocket() (fileDescriptor int, err error) {
	return unix.Socket(unix.AF_NETLINK, unix.SOCK_RAW, unix.NETLINK_ROUTE)
}
