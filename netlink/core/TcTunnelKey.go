package core

// TcTunnelKey - Struct representing a tunnel key action in Linux traffic control.
// This structure is used to encapsulate generic tunnel key parameters in the Linux traffic control system.
// It is often used with netlink messages to configure tunnels for traffic redirection.
//
// Fields:
// - `TcGen`: A general structure that contains common fields for traffic control actions.
// - `Action`: Specifies the action to take with the tunnel key (e.g., redirect traffic, modify headers).
//
// See more details:
// - https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_tunnel_key.h
type TcTunnelKey struct {
	// TcGen - General traffic control action fields.
	TcGen
	// Action - The action to take with the tunnel key.
	Action int32
}
