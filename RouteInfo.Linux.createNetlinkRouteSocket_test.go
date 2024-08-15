//go:build linux

package network

import (
	"golang.org/x/sys/unix"
	"testing"
)

func TestRouteInfo_createNetlinkRouteSocket(t *testing.T) {
	fd, err := createNetlinkRouteSocket()
	if err != nil {
		t.Fatal(err)
	}
	if fd <= 0 {
		t.Fatalf("error: negative file descriptor")
	}
	t.Cleanup(func() {
		_ = unix.Close(fd)
	})
}
