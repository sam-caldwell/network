//go:build linux

package network

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"golang.org/x/sys/unix"
	"log"
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
	var (
		socketFileDescriptor int
		// buffer - binary representation of the complete netlink message
		buffer bytes.Buffer
	)

	if unix.SocketDisableIPv6 {
		return fmt.Errorf("use of ipv6 socket is disabled")
	}

	// Create a socket for netlink communication
	if socketFileDescriptor, err = unix.Socket(unix.AF_NETLINK, unix.SOCK_RAW, unix.NETLINK_ROUTE); err != nil {
		return fmt.Errorf("socket error: %v", err)
	}

	defer func() {
		if cerr := unix.Close(socketFileDescriptor); cerr != nil && err == nil {
			err = cerr
			log.Printf("Error closing socketFileDescriptor(%d): %v", socketFileDescriptor, cerr)
		}
	}()

	log.Println("Netlink message header")
	hdr := unix.NlMsghdr{
		Len:   uint32(unix.SizeofNlMsghdr + unix.SizeofRtMsg),
		Type:  unix.RTM_NEWROUTE,
		Flags: unix.NLM_F_CREATE | unix.NLM_F_REQUEST | unix.NLM_F_ACK,
		Seq:   1,
		Pid:   uint32(unix.Getpid()),
	}

	if err := binary.Write(&buffer, binary.LittleEndian, hdr); err != nil {
		return fmt.Errorf("binary write error: %v", err)
	}

	log.Printf("\ninitial state:\n"+
		"unix.NlMsghdr {\n"+
		"    Len: %d\n"+
		"   Type: %d\n"+
		"  Flags: %d\n"+
		"    Seq: %d\n"+
		"    Pid: %d\n"+
		"}\n\nbuffer: %02x\n",
		hdr.Len, hdr.Type, hdr.Flags, hdr.Seq, hdr.Pid, buffer)

	log.Println("Define IPv6 route message")
	routeMessage := unix.RtMsg{
		Family:   unix.AF_INET6,   // AF_NET / AF_NET6
		Dst_len:  net.IPv6len * 8, // Length in bits
		Src_len:  0,               // zero is okay.  We are not doing source-based routing.
		Tos:      0,
		Table:    unix.RT_TABLE_MAIN,     //254
		Protocol: unix.RTPROT_BOOT,       //0x3
		Scope:    unix.RT_SCOPE_UNIVERSE, //0x00
		Type:     unix.RTN_UNICAST,       //0x01
		Flags:    0,
	}

	log.Printf("\n\n"+
		"unix.RtMsg {\n"+
		"    Family: '%d'\n"+
		"   Dst_len: '%d'\n"+
		"   Src_len: '%d'\n"+
		"       Tos: '%d'\n"+
		"     Table: '%d'\n"+
		"  Protocol: '%d'\n"+
		"     Scope: '%d'\n"+
		"      Type: '%d'\n"+
		"     Flags: '%v'\n"+
		"}\n\n",
		routeMessage.Family, routeMessage.Dst_len, routeMessage.Src_len, routeMessage.Tos, routeMessage.Table,
		routeMessage.Protocol, routeMessage.Scope, routeMessage.Type, routeMessage.Flags)

	if err := binary.Write(&buffer, binary.LittleEndian, routeMessage); err != nil {
		return fmt.Errorf("binary write error: %v", err)
	}

	log.Println("Add destination address attribute")
	if err := createNetlinkRtAttr(&buffer, unix.RTA_DST, net.IPv4len, route.Network.To16()); err != nil {
		return err
	}

	log.Println("Add gateway address attribute")
	if err := createNetlinkRtAttr(&buffer, unix.RTA_GATEWAY, net.IPv4len, route.Gateway.To16()); err != nil {
		return err
	}

	log.Println("sendMessage()")
	if err = sendMessage(socketFileDescriptor, &buffer); err != nil {
		return err
	}

	log.Println("Handle ACK response")
	if err = handleAckResponse(socketFileDescriptor); err != nil {
		return err
	}

	return nil
}
