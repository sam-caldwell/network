package core

import (
	"golang.org/x/sys/unix"
)

func netlinkRouteAttrAndValue(b []byte) (*RtAttr, []byte, int, error) {

	var attr RtAttr
	if err := attr.Deserialize(b); err != nil {
		return nil, nil, 0, err
	}
	if int(attr.Len()) < unix.SizeofRtAttr || int(attr.Len()) > len(b) {
		return nil, nil, 0, unix.EINVAL
	}

	return &attr, b[unix.SizeofRtAttr:], rtaAlignOf(int(attr.Len())), nil

}
