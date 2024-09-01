//go:build linux

package core

// TCPState - tcp connection state tracking
// From https://git.netfilter.org/libnetfilter_conntrack/tree/include/libnetfilter_conntrack/libnetfilter_conntrack_tcp.h
type TCPState uint8

const (
	TCP_CONNTRACK_NONE        TCPState = 0
	TCP_CONNTRACK_SYN_SENT    TCPState = 1
	TCP_CONNTRACK_SYN_RECV    TCPState = 2
	TCP_CONNTRACK_ESTABLISHED TCPState = 3
	TCP_CONNTRACK_FIN_WAIT    TCPState = 4
	TCP_CONNTRACK_CLOSE_WAIT  TCPState = 5
	TCP_CONNTRACK_LAST_ACK    TCPState = 6
	TCP_CONNTRACK_TIME_WAIT   TCPState = 7
	TCP_CONNTRACK_CLOSE       TCPState = 8
	TCP_CONNTRACK_LISTEN      TCPState = 9 /* obsolete */
	TCP_CONNTRACK_SYN_SENT2   TCPState = TCP_CONNTRACK_LISTEN
	TCP_CONNTRACK_MAX         TCPState = 10
	TCP_CONNTRACK_IGNORE      TCPState = 11
)
