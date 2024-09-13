package core

import "encoding/binary"

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
