package netlink

import (
	"fmt"
	"golang.org/x/sys/unix"
	"time"
)

// SetSocketTimeout - set the send/receive timeout for each socket in the netlink handle.
//
// Although the socket timeout has granularity of one microsecond, the effective granularity is floored by the kernel
// timer tick, which default value is four milliseconds.
func (h *Handle) SetSocketTimeout(to time.Duration) error {
	if to < time.Microsecond {
		return fmt.Errorf("invalid timeout, minimal value is %s", time.Microsecond)
	}
	tv := unix.NsecToTimeval(to.Nanoseconds())
	for _, sh := range h.sockets {
		if err := sh.Socket.SetSendTimeout(&tv); err != nil {
			return err
		}
		if err := sh.Socket.SetReceiveTimeout(&tv); err != nil {
			return err
		}
	}
	return nil
}
