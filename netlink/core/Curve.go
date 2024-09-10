package core

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
