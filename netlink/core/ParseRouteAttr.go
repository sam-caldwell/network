package core

import (
	"golang.org/x/sys/unix"
)

func ParseRouteAttr(b []byte) ([]NetlinkRouteAttr, error) {

	var attrs []NetlinkRouteAttr

	for len(b) >= unix.SizeofRtAttr {
		a, vbuf, alen, err := netlinkRouteAttrAndValue(b)
		if err != nil {
			return nil, err
		}
		ra := NetlinkRouteAttr{
			Attr:  *a,
			Value: vbuf[:int(a.RtAttr.Len)-unix.SizeofRtAttr],
		}
		attrs = append(attrs, ra)
		b = b[alen:]
	}

	return attrs, nil

}
