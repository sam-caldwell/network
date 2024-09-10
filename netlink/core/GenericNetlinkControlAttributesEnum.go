package core

// GenericNetlinkControlAttributesEnum - Enumeration for Generic Netlink Control Attributes
//
// This enumeration defines various attributes used by the Generic Netlink (Genl) controller
// to manage and query information about different Netlink families.
// See https://wiki.linuxfoundation.org/networking/generic_netlink_howto
type GenericNetlinkControlAttributesEnum uint8

const (
	// GenericNetlinkControlAttributesUnspec - Unspecified attribute.
	// This is a placeholder for an unspecified or unknown attribute in the Generic Netlink control operations.
	// See https://wiki.linuxfoundation.org/networking/generic_netlink_howto
	GenericNetlinkControlAttributesUnspec GenericNetlinkControlAttributesEnum = 0

	// GenericNetlinkControlAttributeFamilyId - Family ID attribute.
	// This attribute specifies the identifier (ID) for a Netlink family, used to uniquely identify the family
	// within the Netlink framework.
	// See https://wiki.linuxfoundation.org/networking/generic_netlink_howto
	GenericNetlinkControlAttributeFamilyId GenericNetlinkControlAttributesEnum = 1

	// GenericNetlinkControlAttributeFamilyName - Family name attribute.
	// This attribute specifies the name of a Netlink family, which is used to identify and reference the family
	// by name.
	// See https://wiki.linuxfoundation.org/networking/generic_netlink_howto
	GenericNetlinkControlAttributeFamilyName GenericNetlinkControlAttributesEnum = 2

	// GenericNetlinkControlAttributeVersion - Family version attribute.
	// This attribute specifies the version of the Netlink family, indicating the version of the API or protocol
	// supported by the family.
	// See https://wiki.linuxfoundation.org/networking/generic_netlink_howto
	GenericNetlinkControlAttributeVersion GenericNetlinkControlAttributesEnum = 3

	// GenericNetlinkControlAttributeHdrsize - Header size attribute.
	// This attribute specifies the size of the Netlink family-specific header, used in the communication between
	// user space and kernel space.
	// See https://wiki.linuxfoundation.org/networking/generic_netlink_howto
	GenericNetlinkControlAttributeHdrsize GenericNetlinkControlAttributesEnum = 4

	// GenericNetlinkControlAttributeMaxattr - Maximum attribute number.
	// This attribute specifies the highest attribute number that the Netlink family supports, defining the range
	// of valid attributes.
	// See https://wiki.linuxfoundation.org/networking/generic_netlink_howto
	GenericNetlinkControlAttributeMaxattr GenericNetlinkControlAttributesEnum = 5

	// GenericNetlinkControlAttributeOps - Operations attribute.
	// This attribute contains information about the operations supported by the Netlink family, such as commands
	// that can be executed.
	// See https://wiki.linuxfoundation.org/networking/generic_netlink_howto
	GenericNetlinkControlAttributeOps GenericNetlinkControlAttributesEnum = 6

	// GenericNetlinkControlAttributeMcastGroups - Multicast groups attribute.
	// This attribute specifies the multicast groups associated with the Netlink family, used for group communication
	// and event notification.
	// See https://wiki.linuxfoundation.org/networking/generic_netlink_howto
	GenericNetlinkControlAttributeMcastGroups GenericNetlinkControlAttributesEnum = 7
)
