package core

func ParseRouteAttr(b []byte) ([]NetlinkRouteAttr, error) {

	var attrs []NetlinkRouteAttr

	for len(b) >= SizeOfUnixRtAttr {
		a, vbuf, alen, err := netlinkRouteAttrAndValue(b)
		if err != nil {
			return nil, err
		}
		ra := NetlinkRouteAttr{
			Attr:  *a,
			Value: vbuf[:int(a.RtAttr.Len)-SizeOfUnixRtAttr],
		}
		attrs = append(attrs, ra)
		b = b[alen:]
	}

	return attrs, nil

}
