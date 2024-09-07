package netlink

import (
	"fmt"
	"github.com/sam-caldwell/network/netlink/core"
	"golang.org/x/sys/unix"
	"time"
)

// SetSocketTimeout configure timeout for default netlink sockets
func SetSocketTimeout(to time.Duration) error {
	if to < time.Microsecond {
		return fmt.Errorf("invalid timeout, minimal value is %s", time.Microsecond)
	}

	core.SocketTimeoutTv = unix.NsecToTimeval(to.Nanoseconds())
	return nil
}
