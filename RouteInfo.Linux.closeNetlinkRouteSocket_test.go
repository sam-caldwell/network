//go:build linux

package network

import (
	"golang.org/x/sys/unix"
	"testing"
)

func TestRouteInfo_closeNetlinkRouteSocket(t *testing.T) {
	if fd, err := unix.Socket(unix.AF_NETLINK, unix.SOCK_RAW, unix.NETLINK_ROUTE); err != nil {
		t.Fatal(err)
	} else {
		closeNetlinkRouteSocket(fd)
	}
}
