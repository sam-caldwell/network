package netlink

import (
	"github.com/sam-caldwell/network/netlink/core"
	"golang.org/x/sys/unix"
	"time"
)

// GetSocketTimeout - returns timeout value used by default netlink sockets
func GetSocketTimeout() time.Duration {
	nanoseconds := unix.TimevalToNsec(core.SocketTimeoutTv)
	return time.Duration(nanoseconds) * time.Nanosecond
}
