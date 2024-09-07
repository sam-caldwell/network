package core

// Len - Returns the size of the TcTunnelKey structure in bytes.
// This method provides the size of the TcTunnelKey object to help with serialization or buffer allocation.
//
// This method is particularly useful when dealing with netlink messages or similar systems that
// need to know the size of the structure when sending data.
//
// Returns:
// - The size of the TcTunnelKey structure, which is typically defined as a constant.
//
// See the Linux source code for more information:
// - https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_tunnel_key.h
func (msg *TcTunnelKey) Len() int {
	return SizeOfTcTunnelKey
}
