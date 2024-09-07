package core

// Attribute - represents a Netlink message attribute, which is part of the Netlink protocol's way of encoding
// data in a structured and extensible manner. Netlink messages consist of headers followed by a list of
// attributes. Each attribute contains a type and a value, allowing flexible data exchange between the
// kernel and user-space processes.
//
// Attributes are encoded in a Type-Length-Value (TLV) format, where:
// - `Type` identifies what kind of data is contained in the attribute (e.g., IP address, link status).
// - `Value` is the actual data (e.g., the IP address).
//
// In the Linux kernel, this mechanism is handled by the `struct nlattr` in `include/uapi/linux/netlink.h`.
// This structure defines how attributes are structured within a Netlink message.
//
// Reference:
// - Netlink attributes in Linux: https://github.com/torvalds/linux/blob/master/include/uapi/linux/netlink.h
// - Kernel-side attribute parsing: https://github.com/torvalds/linux/blob/master/net/netlink/af_netlink.c
//
// Structure:
//   - `Type`: Represents the type of the Netlink attribute. This is an integer field that indicates the meaning
//     of the value, and is typically encoded using Netlink attribute flags (`NLA_F_*`) that specify things like
//     whether the attribute is nested, has padding, or is optional.
//   - `Value`: A byte slice containing the actual data for the attribute. This could represent an IP address,
//     a device name, a route, or any other form of binary data.
//
// Fields:
// - `Type`: An identifier for the attribute type (represented by `NlaFlags`). This tells the Netlink system how to interpret the attribute.
// - `Value`: A byte slice containing the value or payload of the attribute.
//
// Example Usage:
// ```go
//
//	attr := Attribute{
//	    Type:  NLA_F_NESTED,
//	    Value: []byte{192, 168, 1, 1}, // Example IP address value
//	}
//
// ```
//
// In this example, the `Type` is set to `NLA_F_NESTED`, which indicates that the attribute is nested, and
// the value represents an IP address in a byte array format.
type Attribute struct {
	// Type represents the type of Netlink attribute, encoded using NlaFlags (Netlink Attribute Flags).
	// It determines how the Value field should be interpreted. For example, NLA_F_NESTED indicates a nested attribute.
	Type NlaFlags

	// Value holds the actual data of the attribute. It can be any kind of data, such as an IP address, route information,
	// or even a nested set of attributes. The data is typically binary-encoded and varies in length.
	Value []byte
}
