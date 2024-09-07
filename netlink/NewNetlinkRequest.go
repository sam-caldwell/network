package netlink

import (
	"github.com/sam-caldwell/network/netlink/core"
)

// NewNetlinkRequest creates and returns a new NetlinkRequest object.
// A Netlink request is used to communicate with the kernel, sending commands related to network configuration,
// such as adding routes, modifying addresses, or interacting with other network-related subsystems.
//
// The Netlink protocol (e.g., `RTM_GETADDR` for retrieving addresses or `RTM_NEWLINK` for modifying interfaces)
// is specified through the `proto` parameter, while the `flags` parameter controls the behavior of the request.
// Netlink flags might include request types like `NLM_F_REQUEST` (indicating that this is a request to the kernel)
// or `NLM_F_ACK` (indicating that the kernel should send an acknowledgment).
//
// This method uses the `Handle`'s socket map (if initialized) to create a request specific to a given network namespace
// or Netlink protocol. If no sockets are provided, it falls back to the package-level variable `nextSeqNr`.
//
// Relevant C sources:
//   - Netlink message construction and handling occur in the Linux kernel's `netlink.c`:
//     https://github.com/torvalds/linux/blob/master/net/netlink/af_netlink.c
//   - The specific request protocols, such as `RTM_NEWADDR`, are defined in `rtnetlink.c` for routing and addressing:
//     https://github.com/torvalds/linux/blob/master/net/core/rtnetlink.c
//
// Parameters:
//   - `proto`: The Netlink protocol identifier, such as `RTM_NEWLINK` or `RTM_GETADDR`, indicating the type of request.
//     These constants are defined in the Netlink protocol and control what kind of request is being sent to the kernel.
//   - `flags`: The Netlink message flags, such as `NLM_F_REQUEST`, which control how the kernel should handle the request.
//
// Returns:
//   - `req`: A pointer to the newly created `NetlinkRequest` object. This object can then be used to send the request
//     to the kernel and process the response.
//
// Example Usage:
// ```go
//
//	handle := &Handle{
//	    sockets: make(map[int]*core.SocketHandle),
//	}
//
// req := handle.NewNetlinkRequest(RTM_GETADDR, NLM_F_REQUEST)
// // Further configure and send the request
// ```
//
// In this example, a new Netlink request for retrieving addresses (`RTM_GETADDR`) is created with the `NLM_F_REQUEST` flag.
// The `Handle`'s socket map is used for managing the connection.
//
// Notes:
// - If the `Handle` has an initialized socket map (`sockets`), it will use this map for the Netlink request.
// - If no sockets are provided, the `nextSeqNr` (a package-level sequence number) is used for the request.
//
// See Also:
// - https://man7.org/linux/man-pages/man7/netlink.7.html for general information on the Netlink protocol and its flags.
func (h *Handle) NewNetlinkRequest(proto, flags int) (req *core.NetlinkRequest) {
	// Create a new NetlinkRequest object using the specified protocol and flags.
	req = core.NewNetlinkRequest(proto, flags)

	// If the Handle's `sockets` map is initialized, associate it with the request.
	// This allows the request to use the Handle's managed sockets for communication with the kernel.
	if h.sockets != nil {
		req.Sockets = h.sockets
	}

	// Return the configured NetlinkRequest object.
	return req
}
