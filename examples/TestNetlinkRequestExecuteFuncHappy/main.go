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
			Seq: 12345,
		},
	}
	func() {
		log.Println("Calling GetNetlinkSocket() for pre-flight")
		socket, err := core.GetNetlinkSocket(unix.NETLINK_ROUTE)
		if err != nil {
			log.Fatalf("GetNetlinkSocket() Error: %v", err)
		}
		defer func() { _ = socket.Close() }()
		log.Println("Returned from GetNetlinkSocket()")
	}()
	func() {

		log.Println("calling ExecuteIter()")
		err := req.ExecuteIter(unix.NETLINK_ROUTE, unix.NLMSG_DONE, func(msg []byte) bool {
			fmt.Printf("Received message: %x\n", msg)
			return true
		})
		log.Println("Returned from ExecuteIter()")

		if err != nil {
			log.Fatalf("Error in ExecuteIter: %v", err)
		} else {
			fmt.Println("Netlink message processed successfully.")
		}
	}()
}
