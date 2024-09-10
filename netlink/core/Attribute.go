package core

import "encoding/binary"

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

// Uint64 returns the uint64 value of the Netlink attribute's value, taking into account the `NET_BYTEORDER` flag.
// Netlink attributes may have their values in either host byte order (native to the system) or network byte order
// (big-endian, commonly used for network communication). The `NLA_F_NET_BYTEORDER` flag in the attribute's type
// indicates whether the value is stored in network byte order.
//
// This function checks if the `NLA_F_NET_BYTEORDER` flag is set in the attribute's type and uses the appropriate
// byte order to decode the uint64 value from the attribute's value.
//
// Reference:
//   - The `NLA_F_NET_BYTEORDER` flag is defined in the Linux kernel's `netlink.h` and indicates that the value
//     of the attribute should be interpreted in network byte order (big-endian).
//     https://github.com/torvalds/linux/blob/master/include/uapi/linux/netlink.h
//
// Usage:
// - If `NLA_F_NET_BYTEORDER` is set, the value will be interpreted as big-endian.
// - Otherwise, the value will be interpreted in the system's native byte order (little-endian or big-endian, depending on the architecture).
//
// Fields:
// - `Type`: This field in the `Attribute` struct may contain the `NLA_F_NET_BYTEORDER` flag.
// - `Value`: A byte slice representing the data. In this case, it should be at least 8 bytes long to represent a uint64.
//
// Example Usage:
// ```go
//
//	attr := &Attribute{
//	    Type:  NLA_F_NET_BYTEORDER,
//	    Value: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00}, // Represents 256 in big-endian
//	}
//
// value := attr.Uint64() // value will be 256
// ```
//
// In this example, the `NLA_F_NET_BYTEORDER` flag is set, so the byte slice is interpreted in big-endian format,
// resulting in the uint64 value of 256.
//
// Returns:
// - `uint64`: The decoded uint64 value from the attribute's value field.
//
// Notes:
// - This function assumes that the `Value` slice contains exactly 8 bytes, as required for uint64 representation.
// - If the `Value` slice is not 8 bytes long, the function will panic.
//
// Panics:
// - The function panics if `Value` is `nil` or not exactly 8 bytes long, as these would result in undefined behavior.
func (attr *Attribute) Uint64() uint64 {
	// Ensure the value is exactly 8 bytes, as required for uint64.
	if attr.Value == nil || len(attr.Value) != 8 {
		panic("invalid attribute value: must be exactly 8 bytes")
	}

	// Check if the attribute type has the `NLA_F_NET_BYTEORDER` flag set.
	if attr.Type&NlaFNetByteorder != 0 {
		// If the flag is set, interpret the value in network byte order (big-endian).
		return binary.BigEndian.Uint64(attr.Value)
	} else {
		// Otherwise, interpret the value in the system's native byte order.
		return NativeEndian.Uint64(attr.Value)
	}
}

// Uint32 returns the uint32 value of the Netlink attribute's value, taking into account the `NET_BYTEORDER` flag.
// Netlink attributes may have their values in either host byte order (native to the system) or network byte order
// (big-endian, commonly used for network communication). The `NLA_F_NET_BYTEORDER` flag in the attribute's type
// indicates whether the value is stored in network byte order.
//
// This function checks if the `NLA_F_NET_BYTEORDER` flag is set in the attribute's type and uses the appropriate
// byte order to decode the uint32 value from the attribute's value.
//
// Reference:
//   - The `NLA_F_NET_BYTEORDER` flag is defined in the Linux kernel's `netlink.h` and indicates that the value
//     of the attribute should be interpreted in network byte order (big-endian).
//     https://github.com/torvalds/linux/blob/master/include/uapi/linux/netlink.h
//
// Usage:
// - If `NLA_F_NET_BYTEORDER` is set, the value will be interpreted as big-endian.
// - Otherwise, the value will be interpreted in the system's native byte order (little-endian or big-endian, depending on the architecture).
//
// Fields:
// - `Type`: This field in the `Attribute` struct may contain the `NLA_F_NET_BYTEORDER` flag.
// - `Value`: A byte slice representing the data. In this case, it should be at least 4 bytes long to represent a uint32.
//
// Example Usage:
// ```go
//
//	attr := &Attribute{
//	    Type:  NLA_F_NET_BYTEORDER,
//	    Value: []byte{0x00, 0x00, 0x01, 0x00}, // Represents 256 in big-endian
//	}
//
// value := attr.Uint32() // value will be 256
// ```
//
// In this example, the `NLA_F_NET_BYTEORDER` flag is set, so the byte slice is interpreted in big-endian format,
// resulting in the uint32 value of 256.
//
// Returns:
// - `uint32`: The decoded uint32 value from the attribute's value field.
//
// Note:
// - This function assumes that the value slice contains at least 4 bytes, as required for uint32 representation.
// - If the `Value` slice is smaller than 4 bytes, the result will be undefined behavior and should be handled before calling this function.
func (attr *Attribute) Uint32() uint32 {
	if attr.Value == nil || len(attr.Value) != 4 {
		panic("invalid attribute value")
	}
	// Check if the attribute type has the `NLA_F_NET_BYTEORDER` flag set.
	if attr.Type&NlaFNetByteorder != 0 {
		// If the flag is set, the value is in network byte order (big-endian).
		return binary.BigEndian.Uint32(attr.Value)
	} else {
		// Otherwise, the value is in the system's native byte order.
		return NativeEndian.Uint32(attr.Value)
	}
}
