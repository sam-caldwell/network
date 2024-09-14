package core

// NlaType is an enumeration of Netlink attribute types.
//
// This enumeration defines the different types of attributes that can be used in Netlink messages.
// Netlink is a communication protocol between the Linux kernel and user-space processes, and these attribute types
// help define the structure of messages.
//
// See https://github.com/torvalds/linux/blob/master/include/net/netlink.h
type NlaType int

const (
	// NlaUnspec - NLA_UNSPEC - Unspecified attribute.
	// This is a placeholder for an unspecified or unknown attribute type.
	//
	// See https://github.com/torvalds/linux/blob/master/include/net/netlink.h
	NlaUnspec NlaType = iota

	// NlaU8 - NLA_U8 - 8-bit unsigned integer.
	// This attribute type represents an 8-bit unsigned integer.
	//
	// See https://github.com/torvalds/linux/blob/master/include/net/netlink.h
	NlaU8

	// NlaU16 - NLA_U16 - 16-bit unsigned integer.
	// This attribute type represents a 16-bit unsigned integer.
	//
	// See https://github.com/torvalds/linux/blob/master/include/net/netlink.h
	NlaU16

	// NlaU32 - NLA_U32 - 32-bit unsigned integer.
	// This attribute type represents a 32-bit unsigned integer.
	//
	// See https://github.com/torvalds/linux/blob/master/include/net/netlink.h
	NlaU32

	// NlaU64 - NLA_U64 - 64-bit unsigned integer.
	// This attribute type represents a 64-bit unsigned integer.
	//
	// See https://github.com/torvalds/linux/blob/master/include/net/netlink.h
	NlaU64

	// NlaString - NLA_STRING - Null-terminated string.
	// This attribute type represents a null-terminated string.
	//
	// See https://github.com/torvalds/linux/blob/master/include/net/netlink.h
	NlaString

	// NlaFlag - NLA_FLAG - Flag attribute.
	// This attribute type represents a flag, which is a boolean value typically used as a simple on/off switch.
	//
	// See https://github.com/torvalds/linux/blob/master/include/net/netlink.h
	NlaFlag

	// NlaMsecs - NLA_MSECS - Milliseconds.
	// This attribute type represents a time value in milliseconds.
	//
	// See https://github.com/torvalds/linux/blob/master/include/net/netlink.h
	NlaMsecs

	// NlaNested - NLA_NESTED - Nested attributes.
	// This attribute type represents a set of nested attributes, allowing for complex hierarchical structures
	// in messages.
	//
	// See https://github.com/torvalds/linux/blob/master/include/net/netlink.h
	NlaNested

	// NlaNestedArray - NLA_NESTED_ARRAY - Array of nested attributes.
	// This attribute type represents an array of nested attributes.
	//
	// See https://github.com/torvalds/linux/blob/master/include/net/netlink.h
	NlaNestedArray

	// NlaNulString - NLA_NUL_STRING - Null-terminated string with a maximum length.
	// This attribute type represents a null-terminated string that is limited to a certain length.
	//
	// See https://github.com/torvalds/linux/blob/master/include/net/netlink.h
	NlaNulString

	// NlaBinary - NLA_BINARY - Binary data.
	// This attribute type represents raw binary data.
	//
	// See https://github.com/torvalds/linux/blob/master/include/net/netlink.h
	NlaBinary

	// NlaS8 - NLA_S8 - 8-bit signed integer.
	// This attribute type represents an 8-bit signed integer.
	//
	// See https://github.com/torvalds/linux/blob/master/include/net/netlink.h
	NlaS8

	// NlaS16 - NLA_S16 - 16-bit signed integer.
	// This attribute type represents a 16-bit signed integer.
	//
	// See https://github.com/torvalds/linux/blob/master/include/net/netlink.h
	NlaS16

	// NlaS32 - NLA_S32 - 32-bit signed integer.
	// This attribute type represents a 32-bit signed integer.
	//
	// See https://github.com/torvalds/linux/blob/master/include/net/netlink.h
	NlaS32

	// NlaS64 - NLA_S64 - 64-bit signed integer.
	// This attribute type represents a 64-bit signed integer.
	//
	// See https://github.com/torvalds/linux/blob/master/include/net/netlink.h
	NlaS64

	// NlaBitfield32 - NLA_BITFIELD32 - 32-bit bitfield.
	// This attribute type represents a 32-bit bitfield, used to store multiple boolean flags in a single 32-bit value.
	//
	// See https://github.com/torvalds/linux/blob/master/include/net/netlink.h
	NlaBitfield32

	// NlaReject - NLA_REJECT - Rejection attribute.
	// This attribute type represents a rejection, indicating that a certain value or operation is not acceptable.
	//
	// See https://github.com/torvalds/linux/blob/master/include/net/netlink.h
	NlaReject

	// NlaBe16 - NLA_BE16 - 16-bit big-endian integer.
	// This attribute type represents a 16-bit big-endian integer.
	//
	// See https://github.com/torvalds/linux/blob/master/include/net/netlink.h
	NlaBe16

	// NlaBe32 - NLA_BE32 - 32-bit big-endian integer.
	// This attribute type represents a 32-bit big-endian integer.
	//
	// See https://github.com/torvalds/linux/blob/master/include/net/netlink.h
	NlaBe32

	// NlaSint - NLA_SINT - Signed integer.
	// This attribute type represents a signed integer, where the size is not explicitly defined.
	//
	// See https://github.com/torvalds/linux/blob/master/include/net/netlink.h
	NlaSint

	// NlaUint - NLA_UINT - Unsigned integer.
	// This attribute type represents an unsigned integer, where the size is not explicitly defined.
	//
	// See https://github.com/torvalds/linux/blob/master/include/net/netlink.h
	NlaUint

	// NlaTypeMax - NLA_TYPE_MAX - Maximum attribute type.
	// This constant represents the maximum valid value for Netlink attribute types, used for validation and
	// boundary checks.
	//
	// See https://github.com/torvalds/linux/blob/master/include/net/netlink.h
	NlaTypeMax = iota - 1
)
