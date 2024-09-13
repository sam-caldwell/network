package core

// GenericNetlinkMessage represents the Generic Netlink message header (genlmsghdr) as defined in the Linux kernel.
// This header is part of the Generic Netlink protocol, which is an extension of the Netlink socket interface
// and is used to simplify communication between kernel modules and user-space applications.
//
// The structure corresponds to the `genlmsghdr` struct in the Linux kernel:
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/genetlink.h
//
// In C, the genlmsghdr structure is defined as follows:
//
//	struct genlmsghdr {
//	   __u8    cmd;    /* Command identifier */
//	   __u8    version; /* Interface version */
//	};
//
// It is used in conjunction with a `nlmsghdr` to specify commands and messages for different subsystems
// (families) in Generic Netlink communication.
//
// See Linux source: include/uapi/linux/genetlink.h
//
// Translation to Go:
// - Command: corresponds to the `cmd` field in C, used to specify the command identifier.
// - Version: corresponds to the `version` field in C, specifying the interface version.
type GenericNetlinkMessage struct {
	// Command - Command identifier for the Generic Netlink message.
	// This corresponds to the `cmd` field in the C structure.
	Command uint8

	// Version - Interface version for the Generic Netlink message.
	// This corresponds to the `version` field in the C structure.
	Version uint8
}
