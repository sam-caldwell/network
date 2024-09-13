package main

import (
	"fmt"
	"github.com/sam-caldwell/network/netlink/core"
	"golang.org/x/sys/unix"
	"log"
)

func main() {
	req := &core.NetlinkRequest{
		NlMsghdr: unix.NlMsghdr{
			Seq: 12345,
		},
	}

	err := req.ExecuteIter(unix.SOCK_RAW, unix.NLMSG_DONE, func(msg []byte) bool {
		fmt.Printf("Received message: %x\n", msg)
		return true
	})

	if err != nil {
		log.Fatalf("Error in ExecuteIter: %v", err)
	} else {
		fmt.Println("Netlink message processed successfully.")
	}
}
