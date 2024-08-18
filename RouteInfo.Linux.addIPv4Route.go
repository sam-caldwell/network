//go:build linux

package network

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"golang.org/x/sys/unix"
	"net"
	"time"
	"unsafe"
)

// addIPv4Route - add an IPv4 route using rtnetlink (Linux routing socket)
//
// See https://man7.org/linux/man-pages/man7/rtnetlink.7.html
func (route *RouteInfo) addIPv4Route() (err error) {
	var socketFd int
	var messageBuffer bytes.Buffer

	// Create a socket
	if socketFd, err = unix.Socket(unix.AF_NETLINK, unix.SOCK_RAW, unix.NETLINK_ROUTE); err != nil {
		return fmt.Errorf("socket error: %v", err)
	}
	defer func() {
		if cerr := unix.Close(socketFd); err != nil {
			err = cerr
		}
	}()

	// Set socket receive timeout
	timeout := 5 * time.Second
	if err := unix.SetsockoptTimeval(socketFd, unix.SOL_SOCKET, unix.SO_RCVTIMEO, &unix.Timeval{
		Sec:  int64(timeout / time.Second),
		Usec: int64(timeout % time.Second),
	}); err != nil {
		return fmt.Errorf("setsockopt error: %v", err)
	}

	// Netlink message header
	req := unix.NlMsghdr{
		// Calculate total length of the message
		Len:   uint32(unix.SizeofNlMsghdr + unix.SizeofRtMsg + 2*unix.SizeofRtAttr + 2*net.IPv4len),
		Type:  unix.RTM_NEWROUTE,
		Flags: unix.NLM_F_CREATE | unix.NLM_F_REQUEST | unix.NLM_F_ACK,
		Seq:   1,
		Pid:   uint32(unix.Getpid()),
	}
	// Write the header, route message, and attributes
	if err = binary.Write(&messageBuffer, binary.LittleEndian, req); err != nil {
		return err
	}
	// Define route message
	rt := unix.RtMsg{
		Family:   unix.AF_INET,
		Dst_len:  uint8(net.IPv4len * 8), // Length in bits
		Src_len:  0,
		Tos:      0,
		Table:    unix.RT_TABLE_MAIN,
		Protocol: unix.RTPROT_BOOT,
		Scope:    unix.RT_SCOPE_UNIVERSE,
		Type:     unix.RTN_UNICAST,
	}
	if err = binary.Write(&messageBuffer, binary.LittleEndian, rt); err != nil {
		return err
	}
	dstAttr := unix.RtAttr{
		Len:  uint16(unix.SizeofRtAttr + net.IPv4len),
		Type: unix.RTA_DST,
	}
	if err = binary.Write(&messageBuffer, binary.LittleEndian, dstAttr); err != nil {
		return err
	}
	dstAddr := route.Network.To4()
	if dstAddr == nil {
		return fmt.Errorf("invalid destination address: %v", route.Network.To4().String())
	}
	if _, err = messageBuffer.Write(dstAddr); err != nil {
		return err
	}
	gwAttr := unix.RtAttr{
		Len:  uint16(unix.SizeofRtAttr + net.IPv4len),
		Type: unix.RTA_GATEWAY,
	}
	if err = binary.Write(&messageBuffer, binary.LittleEndian, gwAttr); err != nil {
		return err
	}
	gwAddr := route.Gateway.To4()
	if gwAddr == nil {
		return fmt.Errorf("invalid gateway address: %v", route.Gateway.To4().String())
	}
	if _, err = messageBuffer.Write(gwAddr); err != nil {
		return err
	}

	// Send the message
	err = unix.Sendto(socketFd, messageBuffer.Bytes(), 0, &unix.SockaddrNetlink{Family: unix.AF_NETLINK})
	if err != nil {
		return fmt.Errorf("sendto error: %v", err)
	}

	// Handle ACK response
	ack := make([]byte, 4096)
	n, _, err := unix.Recvfrom(socketFd, ack, 0)
	if err != nil {
		return fmt.Errorf("recvfrom error: %v", err)
	}

	// Check netlink message response
	if n < unix.SizeofNlMsghdr {
		return fmt.Errorf("received message too short")
	}

	nlmsg := (*unix.NlMsghdr)(unsafe.Pointer(&ack[0]))

	if nlmsg.Type == unix.NLMSG_ERROR {
		// Extract the netlink error details
		errorMsg := fmt.Sprintf("netlink error\n"+
			"details:\n"+
			"   Len: 0x%04x\n"+
			"  Type: 0x%04x\n"+
			" Flags: 0x%04x\n"+
			"   Seq: 0x%04x (%0d)\n"+
			"   Pid: 0x%04x", nlmsg.Len, nlmsg.Type, nlmsg.Flags, nlmsg.Seq, nlmsg.Seq, nlmsg.Pid)

		// Additional debug information
		fmt.Println("Netlink error response received:")
		fmt.Println(errorMsg)

		// Check for a specific netlink error code
		if nlmsg.Len > unix.SizeofNlMsghdr {
			// Extract the error details
			errorDetails := ack[unix.SizeofNlMsghdr:n]
			fmt.Printf("Error details: % x\n", errorDetails)
		}

		return fmt.Errorf(errorMsg)
	}

	// Additional handling if required
	// e.g., check for successful response

	return nil
}
