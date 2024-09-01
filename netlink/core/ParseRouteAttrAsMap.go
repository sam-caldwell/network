package core

// ParseRouteAttrAsMap - parses provided buffer that contains raw RtAttrs and returns a map of parsed
// attributes indexed by attribute type or error if occurred.
func ParseRouteAttrAsMap(b []byte) (map[uint16]NetlinkRouteAttr, error) {
	attrMap := make(map[uint16]NetlinkRouteAttr)

	attrs, err := ParseRouteAttr(b)

	if err != nil {
		return nil, err
	}

	for _, attr := range attrs {
		attrMap[attr.Attr.Type] = attr
	}

	return attrMap, nil

}
