package core

// XfrmMsgType - Enum for XFRM message types used in IPsec configuration and management.
// These message types are part of the Linux kernel's XFRM subsystem, which is responsible
// for managing Security Associations (SA), Security Policies (SP), and related components
// in IPsec.
//
// For more information, refer to the Linux kernel source code:
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
type XfrmMsgType uint8

const (
	// XfrmMsgBase - XFRM_MSG_BASE - Base message type for XFRM messages.
	// Used as the starting value for other XFRM message types.
	XfrmMsgBase XfrmMsgType = 0x10

	// XfrmMsgNewSA - XFRM_MSG_NEWSA - Create or replace a new Security Association (SA).
	// Security Associations are used in IPsec to define the encryption and integrity algorithms
	// for protecting data packets.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
	XfrmMsgNewSA XfrmMsgType = 0x10

	// XfrmMsgDelSA - XFRM_MSG_DELSA - Delete an existing Security Association (SA).
	// This message type is used to remove an SA from the system.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
	XfrmMsgDelSA XfrmMsgType = 0x11

	// XfrmMsgGetSA - XFRM_MSG_GETSA - Retrieve information about a Security Association (SA).
	// This message is used to query the state of a particular SA.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
	XfrmMsgGetSA XfrmMsgType = 0x12

	// XfrmMsgNewPolicy - XFRM_MSG_NEWPOLICY - Create or replace a new Security Policy (SP).
	// Security Policies define which traffic is protected and how it is handled in IPsec.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
	XfrmMsgNewPolicy XfrmMsgType = 0x13

	// XfrmMsgDelPolicy - XFRM_MSG_DELPOLICY - Delete an existing Security Policy (SP).
	// This message type is used to remove a security policy from the system.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
	XfrmMsgDelPolicy XfrmMsgType = 0x14

	// XfrmMsgGetPolicy - XFRM_MSG_GETPOLICY - Retrieve information about a Security Policy (SP).
	// This message is used to query the details of a particular security policy.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
	XfrmMsgGetPolicy XfrmMsgType = 0x15

	// XfrmMsgAllocSPI - XFRM_MSG_ALLOCSPI - Allocate a Security Parameter Index (SPI).
	// SPIs are used in IPsec to uniquely identify a Security Association.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
	XfrmMsgAllocSPI XfrmMsgType = 0x16

	// XfrmMsgAcquire - XFRM_MSG_ACQUIRE - Trigger acquisition of an SA.
	// This message is used when a security policy requires an SA, but no SA is available.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
	XfrmMsgAcquire XfrmMsgType = 0x17

	// XfrmMsgExpire - XFRM_MSG_EXPIRE - Notify about an expiring Security Association (SA).
	// This message type is used when an SA is about to expire due to time or byte limits.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
	XfrmMsgExpire XfrmMsgType = 0x18

	// XfrmMsgUpdPolicy - XFRM_MSG_UPDPOLICY - Update an existing Security Policy (SP).
	// This message is used to modify a security policy.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
	XfrmMsgUpdPolicy XfrmMsgType = 0x19

	// XfrmMsgUpdSA - XFRM_MSG_UPDSA - Update an existing Security Association (SA).
	// This message is used to modify an SA.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
	XfrmMsgUpdSA XfrmMsgType = 0x1a

	// XfrmMsgPolExpire - XFRM_MSG_POLEXPIRE - Notify about an expiring Security Policy (SP).
	// This message type is used when a security policy is about to expire.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
	XfrmMsgPolExpire XfrmMsgType = 0x1b

	// XfrmMsgFlushSA - XFRM_MSG_FLUSHSA - Flush all Security Associations (SAs).
	// This message type removes all SAs from the system.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
	XfrmMsgFlushSA XfrmMsgType = 0x1c

	// XfrmMsgFlushPolicy - XFRM_MSG_FLUSHPOLICY - Flush all Security Policies (SPs).
	// This message type removes all security policies from the system.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
	XfrmMsgFlushPolicy XfrmMsgType = 0x1d

	// XfrmMsgNewAE - XFRM_MSG_NEWAE - Create or update a new Anti-Replay window (AE).
	// Anti-Replay windows are used to detect and prevent replay attacks in IPsec.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
	XfrmMsgNewAE XfrmMsgType = 0x1e

	// XfrmMsgGetAE - XFRM_MSG_GETAE - Retrieve Anti-Replay window (AE) information.
	// This message is used to query the status of an anti-replay window.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
	XfrmMsgGetAE XfrmMsgType = 0x1f

	// XfrmMsgReport - XFRM_MSG_REPORT - Send reports or notifications.
	// This message type is used for reporting events related to IPsec.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
	XfrmMsgReport XfrmMsgType = 0x20

	// XfrmMsgMigrate - XFRM_MSG_MIGRATE - Migrate Security Associations (SAs) between interfaces.
	// This message type is used when IPsec SAs need to be moved to a new interface.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
	XfrmMsgMigrate XfrmMsgType = 0x21

	// XfrmMsgNewSadInfo - XFRM_MSG_NEWSADINFO - Create or update new Security Association Database (SAD) info.
	// The SAD holds information about all the active SAs in the system.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
	XfrmMsgNewSadInfo XfrmMsgType = 0x22

	// XfrmMsgGetSadInfo - XFRM_MSG_GETSADINFO - Retrieve Security Association Database (SAD) info.
	// This message is used to query information about the SAD.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
	XfrmMsgGetSadInfo XfrmMsgType = 0x23

	// XfrmMsgNewSpdInfo - XFRM_MSG_NEWSPDINFO - Create or update new Security Policy Database (SPD) info.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
	// The SPD holds information about all the security policies in the system.
	XfrmMsgNewSpdInfo XfrmMsgType = 0x24

	// XfrmMsgGetSpdInfo - XFRM_MSG_GETSPDINFO - Retrieve Security Policy Database (SPD) info.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
	// This message is used to query information about the SPD.
	XfrmMsgGetSpdInfo XfrmMsgType = 0x25

	// XfrmMsgMapping - XFRM_MSG_MAPPING - Handle SA-to-interface mapping.
	// This message is used for managing the mapping between SAs and network interfaces.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
	XfrmMsgMapping XfrmMsgType = 0x26

	// XfrmMsgMax - XFRM_MSG_MAX - The highest defined message type in the XFRM subsystem.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
	XfrmMsgMax XfrmMsgType = 0x26
)
