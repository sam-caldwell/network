package core

// GtpRoleEnum - Enumeration for GTP (GPRS Tunneling Protocol) roles.
//
// This enumeration defines the roles within a GTP setup, such as GGSN (Gateway GPRS Support Node) and SGSN (Serving GPRS Support Node),
// which are key components in mobile networks.
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
type GtpRoleEnum uint8

const (

	// GtpRoleGgsn - GTP_ROLE_GGSN -
	// Represents the GGSN (Gateway GPRS Support Node) role in a GTP context.
	GtpRoleGgsn GtpRoleEnum = 0

	// GtpRoleSgsn - GTP_ROLE_SGSN -
	// Represents the SGSN (Serving GPRS Support Node) role in a GTP context.
	GtpRoleSgsn GtpRoleEnum = 1
)
