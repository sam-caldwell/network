package core

import (
	"golang.org/x/sys/unix"
)

// NetlinkRequest - Represents a netlink request message that can be sent via netlink sockets.
//
// This structure is used to build and send netlink messages to the kernel. It wraps the netlink message header (NlMsghdr)
// and includes the data payload (Data), raw data (RawData), and associated socket handles (Sockets).
//
// References:
// - Linux kernel netlink source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/netlink.h
// - Netlink man page: https://man7.org/linux/man-pages/man7/netlink.7.html
type NetlinkRequest struct {

	// NlMsghdr - Represents the netlink message header (struct nlmsghdr).
	//
	// This header defines the type, flags, sequence number, and PID for the message.
	//
	// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/netlink.h#L101-L114
	unix.NlMsghdr

	// Data - The payload data to include in the message.
	//
	// This field represents the body or payload of the netlink message. It typically contains
	// information such as network device attributes, routing information, etc.
	//
	// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/netlink.h#L85
	Data []NetlinkRequestData

	// RawData - A byte slice representing raw data that can be added to the message.
	//
	// This field allows you to include additional data that may not conform to standard netlink
	// message structures, giving flexibility for custom implementations.
	//
	// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/netlink.h#L85
	RawData []byte

	// Sockets - A map of IP protocol to socket handles.
	//
	// This field holds a mapping of IP protocol types (e.g., TCP, UDP) to corresponding socket
	// handles used for communication. Socket handles facilitate interactions with the kernel over netlink.
	//
	// Reference: https://man7.org/linux/man-pages/man7/netlink.7.html
	Sockets map[IpProtocol]*SocketHandle
}
