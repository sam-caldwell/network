package netlink

import "github.com/sam-caldwell/network/netlink/core"

// Handle represents a Netlink handle that manages Netlink requests on a specific network namespace.
// A `Handle` allows the user to interact with the Netlink API in Linux, sending and receiving requests
// related to networking operations such as address changes, route management, and link (interface) operations.
//
// Netlink handles are typically associated with a specific Netlink family, such as `RTNETLINK`, which is responsible
// for routing, interface configuration, and addressing. A Netlink handle manages the communication socket used for
// interacting with the kernel through Netlink messages.
//
// When multiple requests are sent on the same Netlink family, the same Netlink socket is reused to improve
// performance. The socket is released when the `Handle` is closed using the `Close()` method.
//
// Relevant C sources:
//   - The Linux kernel's implementation of Netlink sockets can be found in `netlink.c`:
//     https://github.com/torvalds/linux/blob/master/net/netlink/af_netlink.c
//   - The Netlink family and communication handling is implemented in `rtnetlink.c`, which processes Netlink messages
//     related to networking operations:
//     https://github.com/torvalds/linux/blob/master/net/core/rtnetlink.c
//
// Fields:
//   - `sockets`: A map where the key is the Netlink family (such as `RTNETLINK`), and the value is the socket handle (`core.SocketHandle`).
//     This map allows the `Handle` to manage multiple Netlink families and sockets in the same namespace.
//   - `lookupByDump`: A flag that determines whether the Handle should use `Dump` requests to look up information in
//     certain operations, such as retrieving lists of addresses or routes.
type Handle struct {
	// sockets stores the socket handles for different Netlink families.
	// For each family (e.g., `RTNETLINK`, `NETLINK_ROUTE`), a socket handle is created and managed.
	// This allows the Handle to interact with different parts of the Netlink API, such as address management
	// and routing table manipulation. When the Handle is closed, all sockets are released.
	sockets map[int]*core.SocketHandle

	// lookupByDump is a flag that indicates whether the Handle should use "Dump" requests to look up data.
	// In Netlink, Dump requests are used to retrieve bulk data from the kernel, such as lists of routes,
	// interfaces, or addresses. This flag controls whether such bulk queries should be performed in certain cases.
	lookupByDump bool
}
