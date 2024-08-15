//go:build linux

package network

import (
	"bytes"
	"fmt"
	"golang.org/x/sys/unix"
	"net"
)

// addIPv6Route - add an IPv6 route using rtnetlink (Linux routing socket)
//
// See https://man7.org/linux/man-pages/man7/rtnetlink.7.html
//
//	Rtnetlink allows the kernel's routing tables to be read and
//	altered.  It is used within the kernel to communicate between
//	various subsystems, though this usage is not documented here, and
//	for communication with user-space programs.  Network routes, IP
//	addresses, link parameters, neighbor setups, queueing
//	disciplines, traffic classes and packet classifiers may all be
//	controlled through NETLINK_ROUTE sockets.  It is based on netlink
//	messages; see netlink(7) for more information.
//
// RTM_NEWROUTE
// RTM_DELROUTE
// RTM_GETROUTE
//
//	Create, remove, or receive information about a network
//	route.  These messages contain an rtmsg structure with an
//	optional sequence of rtattr structures following.  For
//	RTM_GETROUTE, setting rtm_dst_len and rtm_src_len to 0
//	means you get all entries for the specified routing table.
//	For the other fields, except rtm_table and rtm_protocol, 0
//	is the wildcard.
func (route *RouteInfo) addIPv6Route() (err error) {
	// buffer - binary representation of the complete netlink message
	var buffer bytes.Buffer
	var socketFileDescriptor int

	// Create a socket for netlink communication
	if socketFileDescriptor, err = createNetlinkRouteSocket(); err != nil {
		return fmt.Errorf("socket error: %v", err)
	}
	defer func() { _ = unix.Close(socketFileDescriptor) }()

	// Netlink message header
	if err = createNlMsgHdrForNewRouteCreate(&buffer); err != nil {
		return fmt.Errorf("createNlMsgHdrForNewRouteCreate: %v", err)
	}

	// Define IPv4 route message
	if err := createIpRouteMessage(false, &buffer); err != nil {
		return fmt.Errorf("binary write error: %v", err)
	}

	// Add destination address attribute
	if err := createNetlinkRtAttr(&buffer, unix.RTA_DST, net.IPv6len, route.Gateway.To16()); err != nil {
		return err
	}

	// Add gateway address attribute
	if err := createNetlinkRtAttr(&buffer, unix.RTA_GATEWAY, net.IPv6len, route.Gateway.To16()); err != nil {
		return err
	}

	// Send the message
	if err = sendMessage(socketFileDescriptor, &buffer); err != nil {
		return err
	}

	// Handle ACK response (optional but recommended)
	if err = handleAckResponse(socketFileDescriptor); err != nil {
		return err
	}
	return nil
}
