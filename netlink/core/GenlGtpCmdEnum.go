package core

// GenlGtpCmdEnum - Enumerated GTP Commands
//
// This enumeration defines the various commands used in the GTP (GPRS Tunneling Protocol) subsystem
// for managing Packet Data Protocol (PDP) contexts via Generic Netlink.
// See https://docs.kernel.org/networking/gtp.html
type GenlGtpCmdEnum uint8

const (
	// GenlGtpCmdNewpdp - GTP command to create a new PDP context.
	// This command is used to establish a new Packet Data Protocol (PDP) context in the GTP subsystem.
	GenlGtpCmdNewpdp GenlGtpCmdEnum = 0

	// GenlGtpCmdDelpdp - GTP command to delete an existing PDP context.
	// This command is used to remove a previously established PDP context from the GTP subsystem.
	GenlGtpCmdDelpdp GenlGtpCmdEnum = 1

	// GenlGtpCmdGetpdp - GTP command to retrieve information about an existing PDP context.
	// This command is used to query the GTP subsystem for details regarding a specific PDP context.
	GenlGtpCmdGetpdp GenlGtpCmdEnum = 2
)
