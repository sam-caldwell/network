package core

import "encoding/binary"

// Curve - A struct representing a rate curve for shaping traffic in Linux Traffic Control (TC).
//
// This structure defines two linear segments of a curve that may be used to control bandwidth shaping in
// network traffic. The curve is defined by two slopes (`m1` and `m2`) and a displacement (`d`), which is
// commonly used to define token bucket algorithms for traffic shaping.
//
// Fields:
//   - `m1`: Slope of the first segment in bits per second (bps), used for initial rate control.
//   - `d`: X-projection (distance) of the first segment in microseconds (us), representing time or distance
//     after which the rate changes.
//   - `m2`: Slope of the second segment in bits per second (bps), defining the rate after the first segment.
//
// This struct is often used in the Linux kernel's Hierarchical Token Bucket (HTB) and other traffic
// control mechanisms for controlling bandwidth allocation.
//
// See the following references for related kernel definitions and use cases:
// - https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
type Curve struct {
	// m1: Slope of the first segment in bits per second (bps).
	m1 uint32
	// d: X-projection of the first segment in microseconds (us).
	d uint32
	// m2: Slope of the second segment in bits per second (bps).
	m2 uint32
}

// Attrs - Return the attributes of the given Curve
//
// `m1`: The slope of the first segment in bits per second (bps).
// `d`: The displacement or x-projection of the first segment in microseconds (us).
// `m2`: The slope of the second segment in bits per second (bps).
// This method is typically used in traffic control scenarios where the rate needs to change
// dynamically based on network conditions or policies.
//
// See the references for more information:
// - https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
func (c *Curve) Attrs() (m1 uint32, d uint32, m2 uint32) {
	return c.m1, c.d, c.m2
}

// Set - A method for configuring the Curve struct, setting the values of m1, d, and m2 for a traffic shaping curve.
//
// Arguments:
// - `m1`: The slope of the first segment in bits per second (bps).
// - `d`: The displacement or x-projection of the first segment in microseconds (us).
// - `m2`: The slope of the second segment in bits per second (bps).
//
// This method is typically used in traffic control scenarios where the rate needs to change
// dynamically based on network conditions or policies.
//
// See the references for more information:
// - https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
func (c *Curve) Set(m1 uint32, d uint32, m2 uint32) {
	c.m1 = m1
	c.d = d
	c.m2 = m2
}

// DeserializeHfscCurve - Converts a byte slice into a Curve structure.
//
// This function reads the first 12 bytes of the input byte slice and deserializes
// them into the `m1`, `d`, and `m2` fields of the `Curve` struct. It assumes the
// data is encoded in little-endian byte order.
//
// Parameters:
// - b: The byte slice to deserialize.
//
// Returns:
// - A pointer to a `Curve` struct with the fields populated from the byte slice.
func DeserializeHfscCurve(b []byte) *Curve {
	return &Curve{
		m1: binary.LittleEndian.Uint32(b[0:4]),  // First 4 bytes represent `m1`
		d:  binary.LittleEndian.Uint32(b[4:8]),  // Next 4 bytes represent `d`
		m2: binary.LittleEndian.Uint32(b[8:12]), // Final 4 bytes represent `m2`
	}
}
