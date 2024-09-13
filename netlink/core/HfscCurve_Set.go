package core

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
