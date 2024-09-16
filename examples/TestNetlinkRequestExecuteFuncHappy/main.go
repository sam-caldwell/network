package main

import (
	"github.com/sam-caldwell/network/netlink/core"
	"golang.org/x/sys/unix"
	"log"
)

func main() {
	// Create a valid NetlinkRequest for network link information
	req := &core.NetlinkRequest{
		NetlinkMessageHeader: core.NetlinkMessageHeader{
			Seq:   12345,
			Type:  core.RtmGetLink,                  // Request for network link information
			Flags: core.NlmFDump | core.NlmFRequest, // Flags indicating a dump request
		},
	}

	// Create an IfInfomsg and wrap it
	ifmsgData := &core.IfInfoMsg{
		IfInfomsg: unix.IfInfomsg{
			Family: unix.AF_UNSPEC, // Address family
		},
	}

	// Add the IfInfomsgData to the request
	req.Data = append(req.Data, *ifmsgData)

	// Set the length of the netlink message
	req.NetlinkMessageHeader.Len = uint32(core.NetlinkMessageHeaderSize + ifmsgData.Len())

	// Call the Execute function
	res, err := req.Execute(unix.NETLINK_ROUTE, 0)
	if err != nil {
		log.Fatalf("Execute() returned an error: %v", err)
	}

	// Check that we got some responses
	if len(res) == 0 {
		log.Fatal("Execute() returned no messages")
	}

	// Print the received messages for debugging purposes
	for i, msg := range res {
		log.Printf("Message %d: %x\n", i, msg)
	}

	log.Println("Execute() succeeded and returned valid messages.")
}
