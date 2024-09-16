package main

import (
	"fmt"
	"github.com/sam-caldwell/network/netlink/core"
	"golang.org/x/sys/unix"
	"log"
)

func main() {
	// Test for a malformed payload
	req := &core.NetlinkRequest{
		NetlinkMessageHeader: core.NetlinkMessageHeader{
			Seq:   12345,
			Type:  unix.RTM_GETLINK,                     // Request for network link information
			Flags: unix.NLM_F_DUMP | unix.NLM_F_REQUEST, // Flags indicating a dump request
		},
	}

	// Create an IfInfomsgData with an invalid field (for example, a negative Index which is not valid)
	ifmsgData := &core.IfInfoMsg{
		IfInfomsg: unix.IfInfomsg{
			Family: unix.AF_UNSPEC, // Address family
			Index:  -1,             // Invalid Index
		},
	}

	// Add the IfInfomsgData to the request (even though it's invalid)
	req.Data = append(req.Data, *ifmsgData)

	// Set the length of the netlink message deliberately incorrect
	req.NetlinkMessageHeader.Len = 1 // Intentionally incorrect length

	// Proceed with the rest of the code
	log.Println("Calling GetNetlinkSocket() for pre-flight")
	socket, err := core.GetNetlinkSocket(unix.NETLINK_ROUTE)
	if err != nil {
		log.Fatalf("GetNetlinkSocket() Error: %v", err)
	}
	defer func() { _ = socket.Close() }()
	log.Println("Returned from GetNetlinkSocket()")

	log.Println("calling ExecuteIter() with malformed payload")
	err = req.ExecuteIter(unix.NETLINK_ROUTE, 0, func(msg []byte) bool {
		fmt.Printf("Received message: %x\n", msg)
		return true
	})
	log.Println("Returned from ExecuteIter()")

	if err == nil {
		log.Fatalf("Expected error but got none.  Request: %v", req.String())
	} else {
		log.Printf("Error in ExecuteIter (expected failure due to malformed payload):\n"+
			"Error: %v\n"+
			"Request: %v", err, req.String())
	}
}
