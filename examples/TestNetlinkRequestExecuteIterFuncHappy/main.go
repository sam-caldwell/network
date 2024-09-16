package main

import (
	"fmt"
	"github.com/sam-caldwell/network/netlink/core"
	"golang.org/x/sys/unix"
	"log"
)

func main() {
	req := &core.NetlinkRequest{
		NetlinkMessageHeader: core.NetlinkMessageHeader{
			Seq:   12345,
			Type:  core.RtmGetLink,                  // Request for network link information
			Flags: core.NlmFDump | core.NlmFRequest, // Flags indicating a dump request
		},
	}

	// Set the Data field of the request using create an IfInfomsg
	// Create an IfInfomsgData
	ifmsgData := &core.IfInfoMsg{
		IfInfomsg: unix.IfInfomsg{
			Family: core.AfUnspec, // Address family
		},
	}

	// Add the IfInfomsgData to the request
	req.Data = append(req.Data, *ifmsgData)

	// Set the length of the netlink message
	req.NetlinkMessageHeader.Len = uint32(core.NetlinkMessageHeaderSize + len(req.Data))

	// Proceed with the rest of the code
	log.Println("Calling GetNetlinkSocket() for pre-flight")
	socket, err := core.GetNetlinkSocket(unix.NETLINK_ROUTE)
	if err != nil {
		log.Fatalf("GetNetlinkSocket() Error: %v", err)
	}
	defer func() { _ = socket.Close() }()
	log.Println("Returned from GetNetlinkSocket()")

	log.Println("calling ExecuteIter()")
	err = req.ExecuteIter(unix.NETLINK_ROUTE, 0, func(msg []byte) bool {
		fmt.Printf("Received message: %x\n", msg)
		return true
	})
	log.Println("Returned from ExecuteIter()")

	if err != nil {
		log.Fatalf("Error in ExecuteIter:\n"+
			"Error: %v\n"+
			"Request: %v", err, req.String())
	} else {
		fmt.Println("Netlink message processed successfully.")
	}
}
