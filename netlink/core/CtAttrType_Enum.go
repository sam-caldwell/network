package core

// CtAttrType - ctattr_type -  constant in the Linux Netfilter framework that defines the type of connection tracking
// attribute, categorizing the different attributes (such as IP addresses, ports, and protocol information) associated
// with a tracked connection within the connection tracking system.
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/nfnetlink_conntrack.h
type CtAttrType uint8

const (
	// CtaUnspec - CTA_UNSPEC - constant in the Linux Netfilter framework used to indicate that no specific
	// connection tracking attribute is provided or applicable within the connection tracking system. It typically
	// serves as a placeholder or default value.
	CtaUnspec CtAttrType = 0

	// CtaTupleOrig - CTA_TUPLE_ORIG - constant in the Linux Netfilter framework that represents the original tuple
	// of a connection, which includes the source and destination IP addresses, ports, and protocol numbers, as seen
	// in the initial packet that established the connection within the connection tracking system.
	CtaTupleOrig CtAttrType = 1

	// CtaTupleReply - CTA_TUPLE_REPLY - constant in the Linux Netfilter framework that represents the reply tuple
	// of a connection, which includes the source and destination IP addresses, ports, and protocol numbers, as seen
	// in the response packets of the connection within the connection tracking system.
	CtaTupleReply CtAttrType = 2

	// CtaStatus - CTA_STATUS - constant in the Linux Netfilter framework used to represent the status flags of a
	// tracked connection within the connection tracking system, indicating various states and conditions of the
	// connection, such as whether it is established, related, assured, or in need of more processing.
	CtaStatus CtAttrType = 3

	// CtaProtoInfo - CTA_PROTOINFO - constant in the Linux Netfilter framework that represents protocol-specific
	// information for a tracked connection, such as details related to TCP, IpProtoUDP, SCTP, or DCCP protocols, within
	// the connection tracking system.
	CtaProtoInfo CtAttrType = 4

	// CtaHelp - CTA_HELP - constant in the Linux Netfilter framework used to represent the connection helper
	// information associated with a tracked connection, which typically involves assigning specific handling
	// routines or modules (helpers) for certain types of connections, such as FTP or SIP, within the connection
	// tracking system.
	CtaHelp CtAttrType = 5

	// CtaNatSrc - CTA_NAT_SRC - constant in the Linux Netfilter framework that represents the source Network Address
	// Translation (NAT) configuration or information for a tracked connection within the connection tracking system.
	// It is used to track or modify the source address of packets as they are translated by NAT.
	CtaNatSrc CtAttrType = 6

	// CtaNat - CTA_NAT - backwards compatibility - constant in the Linux Netfilter framework that represents the
	// Network Address Translation (NAT) information associated with a tracked connection, including details about
	// how the source or destination addresses and ports are translated within the connection tracking system.
	CtaNat CtAttrType = CtaNatSrc

	// CtaTimeout - CTA_TIMEOUT - constant in the Linux Netfilter framework that represents the timeout value for a
	// tracked connection, indicating how long the connection should be kept in the connection tracking system before
	// it is considered stale and removed if no further packets are seen.
	CtaTimeout CtAttrType = 7

	// CtaMark - CTA_MARK - constant in the Linux Netfilter framework used to represent a mark (a user-defined label)
	// that can be applied to a tracked connection, allowing for customized processing or filtering decisions within
	// the connection tracking system based on that mark.
	CtaMark CtAttrType = 8

	// CtaCountersOrig - CTA_COUNTERS_ORIG - constant in the Linux Netfilter framework that represents the packet
	// and byte counters for the original direction of a tracked connection, providing statistics on the number of
	// packets and the total amount of data exchanged in that direction within the connection tracking system.
	CtaCountersOrig CtAttrType = 9

	// CtaCountersReply - CTA_COUNTERS_REPLY - constant in the Linux Netfilter framework that represents the packet
	// and byte counters for the reply direction of a tracked connection, indicating the number of packets and the
	// total amount of data exchanged in the reply direction within the connection tracking system.
	CtaCountersReply CtAttrType = 10

	// CtaUse - CTA_USE - constant in the Linux Netfilter framework that represents the reference count of a tracked
	// connection, indicating how many times the connection is being used or referenced within the connection tracking
	// system.
	CtaUse CtAttrType = 11

	// CtaId - CTA_ID - constant in the Linux Netfilter framework that represents a unique identifier for a tracked
	// connection within the connection tracking system, used to distinguish this connection from others.
	CtaId CtAttrType = 12

	// CtaNatDst - CTA_NAT_DST - constant in the Linux Netfilter framework that represents the destination Network
	// Address Translation (NAT) configuration or information for a tracked connection, used to track or modify the
	// destination address of packets as they are translated by NAT within the connection tracking system.
	CtaNatDst CtAttrType = 13

	// CtaTupleMaster - CTA_TUPLE_MASTER - constant in the Linux Netfilter framework that represents the master tuple
	// of a tracked connection, typically used for related or associated connections, linking them to the primary or
	// original connection within the connection tracking system.
	CtaTupleMaster CtAttrType = 14

	// CtaSeqAdjOrig - CTA_SEQ_ADJ_ORIG -  constant in the Linux Netfilter framework that represents sequence number
	// adjustments for the original direction of a tracked TCP connection, used to keep track of modifications
	// (such as NAT) that alter the sequence numbers in the original packets within the connection tracking system.
	CtaSeqAdjOrig CtAttrType = 15

	// CtaNatSeqAdjOrig - CTA_NAT_SEQ_ADJ_ORIG - constant in the Linux Netfilter framework that represents sequence
	// number adjustments specifically related to NAT for the original direction of a tracked TCP connection,
	// ensuring that the sequence numbers remain consistent after NAT modifications within the connection tracking
	// system.
	CtaNatSeqAdjOrig CtAttrType = CtaSeqAdjOrig

	// CtaSeqAdjReply - CTA_SEQ_ADJ_REPLY - constant in the Linux Netfilter framework that represents sequence number
	// adjustments for the reply direction of a tracked TCP connection, used to track modifications (such as those
	// made by NAT) that alter the sequence numbers in the reply packets within the connection tracking system.
	CtaSeqAdjReply CtAttrType = 16

	// CtaNatSeqAdjReply - CTA_NAT_SEQ_ADJ_REPLY - constant in the Linux Netfilter framework that represents sequence
	// number adjustments specifically related to NAT for the reply direction of a tracked TCP connection, ensuring
	// that the sequence numbers remain consistent after NAT modifications within the connection tracking system.
	CtaNatSeqAdjReply CtAttrType = CtaSeqAdjReply

	// CtaSecmark - CTA_SECMARK - obsolete - constant in the Linux Netfilter framework that represents a security
	// mark associated with a tracked connection, used to label connections with a security context or policy for
	// further processing by security modules within the connection tracking system.
	CtaSecmark CtAttrType = 17

	// CtaZone - CTA_ZONE - constant in the Linux Netfilter framework that represents the security zone or context
	// associated with a tracked connection, used to segregate connections into different zones for isolated
	// processing within the connection tracking system.
	CtaZone CtAttrType = 18

	// CtaSeccTx - CTA_SECCTX - constant in the Linux Netfilter framework that represents the security context
	// associated with a tracked connection, typically used to store and manage security attributes, such as SELinux
	// labels, within the connection tracking system.
	CtaSeccTx CtAttrType = 19

	// CtaTimestamp - CTA_TIMESTAMP - constant in the Linux Netfilter framework that represents the timestamp
	// information associated with a tracked connection, typically recording when the connection was first seen and
	// last seen within the connection tracking system.
	CtaTimestamp CtAttrType = 20

	// CtaMarkMask - CTA_MARK_MASK - constant in the Linux Netfilter framework used to represent a mask applied to
	// the connection tracking mark (CTA_MARK) for matching or modifying specific bits in the mark, allowing for
	// more granular control in processing or filtering within the connection tracking system.
	CtaMarkMask CtAttrType = 21

	// CtaLabels - CTA_LABELS - constant in the Linux Netfilter framework that represents a set of labels associated
	// with a tracked connection, typically used for tagging or categorizing connections within the connection
	// tracking system.
	CtaLabels CtAttrType = 22

	// CtaLabelsMask - CTA_LABELS_MASK - constant in the Linux Netfilter framework that represents a mask used in
	// conjunction with CTA_LABELS to selectively match or modify specific labels associated with a tracked connection
	// within the connection tracking system.
	CtaLabelsMask CtAttrType = 23

	// CtaSynProxy - CTA_SYNPROXY - constant in the Linux Netfilter framework that represents SYN proxy-related
	// information for a tracked connection, used to mitigate SYN flood attacks by offloading the TCP handshake to
	// the firewall within the connection tracking system.
	CtaSynProxy CtAttrType = 24

	// CtaFilter - CTA_FILTER - constant in the Linux Netfilter framework that represents filtering rules or criteria
	// applied to a tracked connection, used to enforce specific actions or policies within the connection tracking
	// system.
	CtaFilter CtAttrType = 25

	// CtaStatusMask - CTA_STATUS_MASK - constant in the Linux Netfilter framework that represents a mask used to
	// selectively match or modify specific status flags of a tracked connection within the connection tracking system.
	CtaStatusMask CtAttrType = 26

	// CtaMax - CTA_MAX - constant in the Linux Netfilter framework that represents the maximum value for ctattr_type
	// enumeration, indicating the upper limit or the total number of connection tracking attribute types defined
	// within the connection tracking system. It is typically used to validate or iterate over the defined attributes.
	CtaMax CtAttrType = CtaStatusMask
)
