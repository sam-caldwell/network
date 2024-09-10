package core

// GenericControlAttributesOpEnum - Enumeration for Generic Netlink Control Operation Attributes
//
// This enumeration defines various attributes used in the Generic Netlink (Genl) control operations.
// These attributes are utilized by the Generic Netlink controller to manage and query the operations of different
// Netlink families.
// See https://docs.kernel.org/networking/gtp.html
type GenericControlAttributesOpEnum uint8

const (
	// GenericNetlinkControlAttributeOpUnspec - Unspecified attribute.
	// This is a placeholder for an unspecified or unknown attribute in the Generic Netlink control operation.
	//
	// See https://docs.kernel.org/networking/gtp.html
	GenericNetlinkControlAttributeOpUnspec GenericControlAttributesOpEnum = 0

	// GenericNetlinkControlAttributeOpId - Operation identifier attribute.
	// This attribute specifies the identifier for a particular operation in the Generic Netlink control.
	//
	// See https://docs.kernel.org/networking/gtp.html
	GenericNetlinkControlAttributeOpId GenericControlAttributesOpEnum = 1

	// GenericNetlinkControlAttributeOpFlags - Operation flags attribute.
	// This attribute specifies the flags associated with a particular operation, defining its behavior or
	// capabilities.
	//
	// See https://docs.kernel.org/networking/gtp.html
	GenericNetlinkControlAttributeOpFlags GenericControlAttributesOpEnum = 2
)
