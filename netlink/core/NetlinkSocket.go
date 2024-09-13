package core

import (
	"errors"
	"fmt"
	"golang.org/x/sys/unix"
	"sync"
	"sync/atomic"
)

// NetlinkSocket - sockaddr_nl - represents a netlink socket with an associated file descriptor (fd) and sockaddr
// structure (lsa). The socket is protected by a mutex to ensure thread-safe operations, preventing concurrent access
// issues during send, receive, or other socket operations.
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netlink.h
type NetlinkSocket struct {

	// Mutex to ensure thread-safe operations on the socket.
	sync.Mutex

	// fd - File descriptor for the netlink socket.
	fd int32

	// lsa - Socket address information specific to netlink.
	lsa unix.SockaddrNetlink
}

// Close - close the NetlinkSocket by safely setting its file descriptor (fd) to -1 and then closing the original fd.
//
// Thread Safety:
//
//	The method uses an atomic operation (atomic.SwapInt32) to ensure that the file descriptor is safely set to -1
//	before closing the original file descriptor. This prevents race conditions where multiple goroutines might attempt
//	to close the socket simultaneously. The atomic swap ensures that only the first call to Close will actually close
//	the file descriptor, while subsequent calls will safely do nothing, as the fd will already be set to -1.
//
// Returns:
//
//	An error, if any, encountered while closing the socket.
func (s *NetlinkSocket) Close() error {
	fd := int(atomic.SwapInt32(&s.fd, -1))
	return unix.Close(fd)
}

// GetFd - Return the NetlinkSocket file descriptor
func (s *NetlinkSocket) GetFd() int {

	return int(atomic.LoadInt32(&s.fd))

}

// GetPid retrieves the PID (Process ID) associated with the netlink socket.
//
// This function uses the socket's file descriptor to obtain the socket address (Sockaddr)
// associated with the netlink socket. If the socket is of type SockaddrNetlink, it returns
// the PID stored in the socket address structure. Otherwise, it returns an error indicating
// that the socket type is unsupported.
func (s *NetlinkSocket) GetPid() (uint32, error) {

	socketFd := int(atomic.LoadInt32(&s.fd))

	lsa, err := unix.Getsockname(socketFd)

	if err != nil {
		return 0, err
	}

	switch v := lsa.(type) {
	case *unix.SockaddrNetlink:
		// Return the Pid from the SockaddrNetlink structure.
		return v.Pid, nil
	}

	return 0, errors.New("unsupported socket type")
}

// SetSendTimeout - set send timeout on the socket
//
// Set send timeout of SOCKET_SEND_TIMEOUT, this will allow the Send to periodically unblock and avoid some routine
// from hanging on send on a closed fd.
func (s *NetlinkSocket) SetSendTimeout(timeout *unix.Timeval) error {

	return unix.SetsockoptTimeval(int(s.fd), unix.SOL_SOCKET, unix.SO_SNDTIMEO, timeout)

}

// Receive - read a message from the netlink socket and returns the parsed NetlinkMessage slice, the sender's
// SockaddrNetlink, and an error if any occurred during the process.
func (s *NetlinkSocket) Receive() ([]NetlinkMessage, *unix.SockaddrNetlink, error) {

	var (
		err             error
		ok              bool
		numberBytesRecd int
		from            unix.Sockaddr
		fromAddr        *unix.SockaddrNetlink
		nl              []NetlinkMessage
		rb              [ReceiveBufferSize]byte
	)

	if socketFd := int(atomic.LoadInt32(&s.fd)); socketFd < 0 {
		return nil, nil, errors.New("cannot receive. socket closed")
	} else {
		numberBytesRecd, from, err = unix.Recvfrom(socketFd, rb[:], 0)
	}

	if err != nil {
		return nil, nil, err
	}

	if fromAddr, ok = from.(*unix.SockaddrNetlink); !ok {
		return nil, nil, errors.New("error (converting to netlink sockaddr)")
	}

	if numberBytesRecd < unix.NLMSG_HDRLEN {
		return nil, nil, errors.New("error (short response from netlink)")
	}

	msgLength := nlmAlignOf(numberBytesRecd)

	rb2 := make([]byte, msgLength)

	copy(rb2, rb[:msgLength])

	if nl, err = ParseNetlinkMessage(rb2); err != nil {
		return nil, nil, err
	}

	return nl, fromAddr, nil
}

// Send - send a given NetlinkRequest via the current socket file descriptor to netlink
func (s *NetlinkSocket) Send(request *NetlinkRequest) (err error) {

	if request == nil {
		return fmt.Errorf("nil request")
	}

	socketFd := int(atomic.LoadInt32(&s.fd))

	if socketFd < 0 {

		return fmt.Errorf("send called on a closed socket")

	}

	var requestBytes []byte
	if requestBytes, err = request.Serialize(); err != nil {
		return err
	}

	if err = unix.Sendto(socketFd, requestBytes, 0, &s.lsa); err != nil {
		return err
	}

	return err

}

// SetExtAck - request error messages to be reported on the socket
func (s *NetlinkSocket) SetExtAck(enable bool) error {

	var n int

	if enable {
		n = 1
	}

	return unix.SetsockoptInt(int(s.fd), unix.SOL_NETLINK, unix.NETLINK_EXT_ACK, n)

}

// SetReceiveBufferSize - set receive buffer size on the socket
func (s *NetlinkSocket) SetReceiveBufferSize(size int, force bool) error {

	opt := unix.SO_RCVBUF

	if force {
		opt = unix.SO_RCVBUFFORCE
	}

	return unix.SetsockoptInt(int(s.fd), unix.SOL_SOCKET, opt, size)

}

// SetReceiveTimeout - set receive timeout on the socket
//
// Set a read timeout of SOCKET_READ_TIMEOUT, this will allow the Read to periodically unblock and avoid that a routine
// remains stuck on a recvmsg on a closed fd
func (s *NetlinkSocket) SetReceiveTimeout(timeout *unix.Timeval) error {

	return unix.SetsockoptTimeval(int(s.fd), unix.SOL_SOCKET, unix.SO_RCVTIMEO, timeout)

}
