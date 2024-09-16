package core

// GenlGtpAttribute - Enumerated GTP Attributes
//
// This enumeration defines various attributes used in the GTP (GPRS Tunneling Protocol) subsystem
// for managing Packet Data Protocol (PDP) contexts via Generic Netlink. These attributes provide
// the necessary information for the GTP commands to operate correctly.
// See https://docs.kernel.org/networking/gtp.html
type GenlGtpAttribute uint8

const (
	// GenlGtpAttrUnspec - Unspecified attribute.
	// This is a placeholder for an unspecified or unknown attribute.
	// See https://docs.kernel.org/networking/gtp.html
	GenlGtpAttrUnspec GenlGtpAttribute = 0

	// GenlGtpAttrLink - The network link associated with the GTP tunnel.
	// This attribute specifies the network interface (link) over which the GTP tunnel is established.
	// See https://docs.kernel.org/networking/gtp.html
	GenlGtpAttrLink GenlGtpAttribute = 1

	// GenlGtpAttrVersion - The GTP version.
	// This attribute indicates the version of the GTP protocol being used (e.g., GTPv1 or GTPv2).
	// See https://docs.kernel.org/networking/gtp.html
	GenlGtpAttrVersion GenlGtpAttribute = 2

	// GenlGtpAttrTid - The Tunnel Identifier (TID).
	// This attribute specifies the Tunnel Identifier, which is used to identify a specific GTP tunnel.
	// See https://docs.kernel.org/networking/gtp.html
	GenlGtpAttrTid GenlGtpAttribute = 3

	// GenlGtpAttrPeerAddress - The address of the peer in the GTP tunnel.
	// This attribute specifies the IP address of the peer node in the GTP tunnel, typically the remote GTP endpoint.
	// See https://docs.kernel.org/networking/gtp.html
	GenlGtpAttrPeerAddress GenlGtpAttribute = 4

	// GenlGtpAttrMsAddress - The address of the mobile station.
	// This attribute specifies the IP address of the mobile station (MS) associated with the PDP context.
	// See https://docs.kernel.org/networking/gtp.html
	GenlGtpAttrMsAddress GenlGtpAttribute = 5

	// GenlGtpAttrFlow - Flow label or identifier for the GTP tunnel.
	// This attribute is used to specify a flow label or identifier for the traffic within the GTP tunnel.
	// See https://docs.kernel.org/networking/gtp.html
	GenlGtpAttrFlow GenlGtpAttribute = 6

	// GenlGtpAttrNetNsFd - Network namespace file descriptor.
	// This attribute specifies the file descriptor for the network namespace where the GTP tunnel operates.
	// See https://docs.kernel.org/networking/gtp.html
	GenlGtpAttrNetNsFd GenlGtpAttribute = 7

	// GenlGtpAttrITei - Incoming Tunnel Endpoint Identifier (I-TEI).
	// This attribute specifies the Tunnel Endpoint Identifier (TEI) for incoming traffic in the GTP tunnel.
	// See https://docs.kernel.org/networking/gtp.html
	GenlGtpAttrITei GenlGtpAttribute = 8

	// GenlGtpAttrOTei - Outgoing Tunnel Endpoint Identifier (O-TEI).
	// This attribute specifies the Tunnel Endpoint Identifier (TEI) for outgoing traffic in the GTP tunnel.
	// See https://docs.kernel.org/networking/gtp.html
	GenlGtpAttrOTei GenlGtpAttribute = 9

	// GenlGtpAttrPad - Padding attribute.
	// This attribute is used for padding in the message structure, ensuring proper alignment.
	// See https://docs.kernel.org/networking/gtp.html
	GenlGtpAttrPad GenlGtpAttribute = 10
)
