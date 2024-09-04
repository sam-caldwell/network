package core

// Equal - Return whether two IPv6SrHdr structs are equal
func (s1 *IPv6SrHdr) Equal(s2 IPv6SrHdr) bool {

	if len(s1.Segments) != len(s2.Segments) {
		return false
	}

	for i := range s1.Segments {
		if !s1.Segments[i].Equal(s2.Segments[i]) {
			return false
		}
	}

	return s1.nextHdr == s2.nextHdr &&
		s1.hdrLen == s2.hdrLen &&
		s1.routingType == s2.routingType &&
		s1.segmentsLeft == s2.segmentsLeft &&
		s1.firstSegment == s2.firstSegment &&
		s1.flags == s2.flags
	// reserved doesn't need to be identical.
}
