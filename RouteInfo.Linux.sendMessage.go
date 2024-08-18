//go:build linux

package network

import (
	"bytes"
	"fmt"
	"golang.org/x/sys/unix"
	"log"
)

// sendMessage - send the netlink message
func sendMessage(socketFileDescriptor int, buffer *bytes.Buffer) error {

	log.Printf("sendMessage(): calling unix.Sendto()")
	if err := unix.Sendto(socketFileDescriptor, buffer.Bytes(), 0, &unix.SockaddrNetlink{Family: unix.AF_NETLINK}); err != nil {
		log.Printf("sendMessage(): unix.Sendto() failed: %v", err)
		return fmt.Errorf("sendto error: %v", err)
	}
	log.Println("sendMessage(): sent")

	return nil

}
