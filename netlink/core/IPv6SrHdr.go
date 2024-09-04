package core

import (
	"net"
)

type IPv6SrHdr struct {
	nextHdr      uint8
	hdrLen       uint8
	routingType  uint8
	segmentsLeft uint8
	firstSegment uint8
	flags        uint8
	reserved     uint16

	Segments []net.IP
}
