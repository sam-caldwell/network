package core

import (
	"golang.org/x/sys/unix"
	"syscall"
)

func ParseRouteAttr(b []byte) ([]NetlinkRouteAttr, error) {
	var attrs []syscall.NetlinkRouteAttr
	for len(b) >= unix.SizeofRtAttr {
		a, vbuf, alen, err := netlinkRouteAttrAndValue(b)
		if err != nil {
			return nil, err
		}
		ra := NetlinkRouteAttr{Attr: RtAttr(*a), Value: vbuf[:int(a.Len)-unix.SizeofRtAttr]}
		attrs = append(attrs, ra)
		b = b[alen:]
	}
	return attrs, nil
}
