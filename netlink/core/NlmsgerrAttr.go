package core

// NlmsgerrAttr - nlmsgerr attributes
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netlink.h
//
// enum nlmsgerr_attrs - nlmsgerr attributes
//
//	@NLMSGERR_ATTR_UNUSED: unused
//	@NLMSGERR_ATTR_MSG: error message string (string)
//	@NLMSGERR_ATTR_OFFS: offset of the invalid attribute in the original message, counting from the beginning
//	                     of the header (u32)
//	@NLMSGERR_ATTR_COOKIE: arbitrary subsystem specific cookie to be used - in the success case - to
//	                       identify a created object or operation or similar (binary)
//	@NLMSGERR_ATTR_POLICY: policy for a rejected attribute
//	@NLMSGERR_ATTR_MISS_TYPE: type of missing required attribute, %NLMSGERR_ATTR_MISS_NEST will not be present
//	                          if the attribute was missing at the message level
//	@NLMSGERR_ATTR_MISS_NEST: offset of the nest where attribute was missing
//	@NLMSGERR_ATTR_MAX: highest attribute number
type NlmsgerrAttr int

const (
	// NlmsgerrAttrUnused - NLMSGERR_ATTR_UNUSED - Unused attribute in NLMSGERR netlink error.
	NlmsgerrAttrUnused NlmsgerrAttr = 0

	// NlmsgerrAttrMsg - NLMSGERR_ATTR_MSG - Netlink attribute for an error message string.
	NlmsgerrAttrMsg NlmsgerrAttr = 1

	// NlmsgerrAttrOffs - NLMSGERR_ATTR_OFFS - offset of the invalid attribute in the original message, counting
	//   from the beginning of the header (u32)
	NlmsgerrAttrOffs NlmsgerrAttr = 2

	// NlmsgerrAttrCookie - NLMSGERR_ATTR_COOKIE - arbitrary subsystem specific cookie to be used - in the success
	//  case - to identify a created object or operation or similar (binary)
	NlmsgerrAttrCookie NlmsgerrAttr = 3

	// NlmsgerrAttrPolicy - NLMSGERR_ATTR_POLICY - policy for a rejected attribute
	NlmsgerrAttrPolicy NlmsgerrAttr = 4

	// NlmsgerrAttrMissType - NLMSGERR_ATTR_MISS_TYPE - type of missing required attribute, NlmsgerrAttrMissNest
	//   will not be present if the attribute was missing at the message level.
	NlmsgerrAttrMissType NlmsgerrAttr = 5

	// NlmsgerrAttrMissNest - NLMSGERR_ATTR_MISS_NEST - will not be present if the attribute was missing at the
	//   message level
	NlmsgerrAttrMissNest NlmsgerrAttr = 6

	// NlmsgerrAttrMax - NLMSGERR_ATTR_MAX - highest attribute number
	NlmsgerrAttrMax NlmsgerrAttr = NlmsgerrAttrMissNest
)
