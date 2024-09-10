package core

import (
	"golang.org/x/sys/unix"
	"unsafe"
)

// GenlMsg, genetlink constants
const (
	// GenlAdminPerm - GENL_ADMIN_PERM - This constant represents the permission required for administrative commands
	//                 in the Generic Netlink (Genl) framework, ensuring that only privileged users can execute certain
	//                 commands.
	GenlAdminPerm = unix.GENL_ADMIN_PERM

	// GenlCmdCapDo - GENL_CMD_CAP_DO - This constant indicates that a Generic Netlink command has the capability to
	//                perform an action (typically an administrative or control operation).
	GenlCmdCapDo = unix.GENL_CMD_CAP_DO

	// GenlCmdCapDump - GENL_CMD_CAP_DUMP - This constant signifies that a Generic Netlink command has the capability
	//                  to dump data, usually for retrieving information in bulk.
	GenlCmdCapDump = unix.GENL_CMD_CAP_DUMP

	// GenlCmdCapHaspol - GENL_CMD_CAP_HASPOL - This constant specifies that a Generic Netlink command has an
	//                    associated policy, which defines the expected attributes and constraints for that command.
	GenlCmdCapHaspol = unix.GENL_CMD_CAP_HASPOL

	// GenlHdrlen - GENL_HDRLEN - This constant defines the length of the Generic Netlink header, which is added to
	//              all Generic Netlink messages.
	GenlHdrlen = unix.GENL_HDRLEN

	// GenlIdCtrl - GENL_ID_CTRL - This constant represents the identifier for the Generic Netlink control family,
	//              which is used to manage and query other Generic Netlink families.
	GenlIdCtrl = unix.GENL_ID_CTRL
)

// constants.struct.size:
//
// This file defines the structure size constants used elsewhere in the program
// We do this here to make our use of "unsafe" an exception rather than the norm.
// This should be the only place where we need unsafe (outside of tests).
const (
	// SizeOfIfAddressMessage     = 0x8 // bytes as derived from unix.SizeOfIfAddressMessage
	SizeOfIfAddressMessage = int(unsafe.Sizeof(IfAddressMessage{}))

	// SizeOfIfaCacheInfo - 0x16 // bytes as derived from unix.SizeOfIfaCacheInfo
	SizeOfIfaCacheInfo = int(unsafe.Sizeof(IfaCacheInfo{}))

	// SizeOfBridgeVlanInfo - size of BridgeVlanInfo
	SizeOfBridgeVlanInfo = int(unsafe.Sizeof(BridgeVlanInfo{}))

	// SizeOfCnMsgOp - size of CnMsgOp struct
	SizeOfCnMsgOp = int(unsafe.Sizeof(CnMsgOp{}))

	// SizeOfGenlMsg - size of Genlmsg struct
	SizeOfGenlMsg = int(unsafe.Sizeof(Genlmsg{}))

	// SizeOfNfGenMsg - Track the message sizes for the correct serialization/deserialization
	SizeOfNfGenMsg = int(unsafe.Sizeof(Nfgenmsg{}))

	// SizeOfNlMsgHdr - Size of unix.NlMsghdr struct
	SizeOfNlMsgHdr = int(unsafe.Sizeof(unix.NlMsghdr{}))

	//SizeOfRtAttr - size of the RtAttrWrapper struct
	SizeOfRtAttr = int(unsafe.Sizeof(RtAttr{}))

	//SizeOfRtGenMsg - Size of the RtGenMsg struct
	SizeOfRtGenMsg = int(unsafe.Sizeof(RtGenMsg{}))

	// SizeOfRtNextHop - Size of the RtNexthop struct
	SizeOfRtNextHop = int(unsafe.Sizeof(RtNexthop{}))

	// SizeOfTcNetemRate - Size of the TcNetemRate struct, defining rate options in netem.
	SizeOfTcNetemRate = int(unsafe.Sizeof(TcNetemRate{}))

	// SizeOfTcActionMsg - Size of the TcActionMsg struct, which represents actions taken on packets in TC.
	SizeOfTcActionMsg = int(unsafe.Sizeof(TcActionMsg{}))

	// SizeOfTcConnmark - Size of the TcConnmark struct, used for connection marking in TC.
	SizeOfTcConnmark = int(unsafe.Sizeof(TcConnmark{}))

	// SizeOfTcCsum - Size of the TcCsum struct, representing checksum actions in traffic control.
	SizeOfTcCsum = int(unsafe.Sizeof(TcCsum{}))

	// SizeOfTcGen - Size of the TcGen struct, representing general traffic control options.
	SizeOfTcGen = int(unsafe.Sizeof(TcGen{}))

	// SizeOfTcHtbCopt - Size of the TcHtbCopt struct, representing options for hierarchical token bucket (HTB)
	// classes in TC.
	SizeOfTcHtbCopt = int(unsafe.Sizeof(TcHtbCopt{}))

	// SizeOfTcHtbGlob - Size of the TcHtbGlob struct, representing global parameters for HTB.
	SizeOfTcHtbGlob = int(unsafe.Sizeof(TcHtbGlob{}))

	// SizeOfIfInfoMsg -  Size of the IfInfoMsg struct
	SizeOfIfInfoMsg = int(unsafe.Sizeof(IfInfoMsg{}))

	// SizeOfTcMsg - Size of the TcMsg struct, which is used to configure traffic control (TC) in Linux.
	SizeOfTcMsg = int(unsafe.Sizeof(TcMsg{}))

	// SizeOfTcMirred - Size of the TcMirred struct, representing "mirred" (mirror/redirect) actions in TC.
	SizeOfTcMirred = int(unsafe.Sizeof(TcMirred{}))

	// SizeOfTcNetemCorr - Size of the TcNetemCorr struct, representing correlation options in netem.
	SizeOfTcNetemCorr = int(unsafe.Sizeof(TcNetemCorr{}))

	// SizeOfTcNetemCorrupt - Size of the TcNetemCorrupt struct, representing options for packet corruption in netem.
	SizeOfTcNetemCorrupt = int(unsafe.Sizeof(TcNetemCorrupt{}))

	// SizeOfTcNetemQopt - Size of TcNetemQopt struct, used for network emulation (netem) options in traffic control.
	SizeOfTcNetemQopt = int(unsafe.Sizeof(TcNetemQopt{}))

	// SizeOfTcNetemReorder - Size of the TcNetemReorder struct, used to configure packet reordering in netem.
	SizeOfTcNetemReorder = int(unsafe.Sizeof(TcNetemReorder{}))

	// SizeOfTcPeditKey - size of the TcPeditKey struct
	SizeOfTcPeditKey = int(unsafe.Sizeof(TcPeditKey{}))

	// SizeOfTcPeditSel - size of the TcPeditSel struct
	SizeOfTcPeditSel = int(unsafe.Sizeof(TcPeditSel{}))

	// SizeOfTcPolice - Size of the TcPolice struct, representing police actions in TC.
	SizeOfTcPolice = int(unsafe.Sizeof(TcPolice{}))

	// SizeOfTcPriorityMap - Size of the TcPriorityMap struct, used to map priorities for traffic control.
	SizeOfTcPriorityMap = int(unsafe.Sizeof(TcPriorityMap{}))

	// SizeOfTcRateSpec - Size of the TcRateSpec struct, which defines the rate specification for traffic control.
	SizeOfTcRateSpec = int(unsafe.Sizeof(TcRateSpec{}))

	// SizeOfTcSfqQopt - Size of the TcSfqQopt struct, representing options for SFQ.
	SizeOfTcSfqQopt = int(unsafe.Sizeof(TcSfqQopt{}))

	// SizeOfTcSfqQoptV1 - Size of TcSfqQoptV1 struct, representing options for stochastic fair queuing (SFQ) version 1.
	SizeOfTcSfqQoptV1 = int(unsafe.Sizeof(TcSfqQoptV1{}))

	// SizeOfTcSfqRedStats - Size of TcSfqRedStats struct, representing RED (random early detection) statistics in SFQ.
	SizeOfTcSfqRedStats = int(unsafe.Sizeof(TcSfqRedStats{}))

	// SizeOfTcSkbEdit - Size of the TcSkbEdit struct, representing modifications to skb (socket buffer) in TC.
	SizeOfTcSkbEdit = int(unsafe.Sizeof(TcSkbEdit{}))

	// SizeOfTcTbfQopt - Size of TcTbfQopt struct, represents options for token bucket filter (TBF) in traffic control.
	SizeOfTcTbfQopt = int(unsafe.Sizeof(TcTbfQopt{}))

	// SizeOfTcTunnelKey - Size of the TcTunnelKey struct, representing tunnel key actions in TC.
	SizeOfTcTunnelKey = int(unsafe.Sizeof(TcTunnelKey{}))

	// SizeOfTcf - Size of the Tcf struct, representing a filter in traffic control.
	SizeOfTcf = int(unsafe.Sizeof(Tcf{}))

	// SizeOfTcU32Sel - golang will calculate the wrong size if you use unsafe.SizeOf
	SizeOfTcU32Sel = 16

	// SizeOfTcU32Key - golang will calculate the wrong size if you use unsafe.SizeOf
	SizeOfTcU32Key = 16

	// SizeOfUint32Bitfield - The size of a 32-bit bitfield used in various traffic control structs.
	SizeOfUint32Bitfield = int(unsafe.Sizeof(Uint32Bitfield{}))

	// SizeOfUnixRtAttr - The size of a unix.RtAttr struct
	SizeOfUnixRtAttr = int(unsafe.Sizeof(unix.RtAttr{}))

	// SizeOfUnixRtMsg - the size of unix.RtMsg struct
	SizeOfUnixRtMsg = int(unsafe.Sizeof(unix.RtMsg{}))

	// SizeOfVfGUID - Size of VfGUID struct
	SizeOfVfGUID = int(unsafe.Sizeof(VfGUID{}))

	// SizeOfVfLinkState - Size of VfLinkState struct
	SizeOfVfLinkState = int(unsafe.Sizeof(VfLinkState{}))

	// SizeOfVfMac - Size of VfMac struct
	SizeOfVfMac = int(unsafe.Sizeof(VfMac{}))

	// SizeOfVfRate - Size of VfRate struct
	SizeOfVfRate = int(unsafe.Sizeof(VfRate{}))

	// SizeOfVfRssQueryEn - Size of VfRssQueryEn struct
	SizeOfVfRssQueryEn = int(unsafe.Sizeof(VfRssQueryEn{}))

	// SizeOfVfSpoofchk - Size of VfSpoofchk
	SizeOfVfSpoofchk = int(unsafe.Sizeof(VfSpoofchk{}))

	// SizeOfVfStats - Size of VfStats struct
	SizeOfVfStats = int(unsafe.Sizeof(VfStats{}))

	// SizeOfVfTrust - Size of VfTrust struct
	SizeOfVfTrust = int(unsafe.Sizeof(VfTrust{}))

	// SizeOfVfTxRate - Size of VfTxRate struct
	SizeOfVfTxRate = int(unsafe.Sizeof(VfTxRate{}))

	// SizeOfVfVlan - Size of VfVlan struct
	SizeOfVfVlan = int(unsafe.Sizeof(VfVlan{}))

	// SizeOfVfVlanInfo - Size of IfLaVfVlanInfoStruct struct
	SizeOfVfVlanInfo = int(unsafe.Sizeof(IfLaVfVlanInfoStruct{}))

	// SizeOfXfrmAddress - Size of the XfrmAddress struct, used to represent addresses in XFRM (IPsec framework).
	// Assumes largest size (ipV6)
	SizeOfXfrmAddress = int(4 * 32)

	// SizeOfXfrmAlgo - Size of the XfrmAlgo struct, representing cryptographic algorithms for IPsec SAs.
	SizeOfXfrmAlgo = int(unsafe.Sizeof(XfrmAlgo{}))

	// SizeOfXfrmAlgoAEAD - Size of XfrmAlgoAEAD struct, representing AEAD (authenticated encryption) algorithms
	// for IPsec.
	SizeOfXfrmAlgoAEAD = int(unsafe.Sizeof(XfrmAlgoAEAD{}))

	// SizeOfXfrmAlgoAuth - Size of the XfrmAlgoAuth struct, representing authentication algorithms for IPsec SAs.
	SizeOfXfrmAlgoAuth = int(unsafe.Sizeof(XfrmAlgoAuth{}))

	// SizeOfXfrmEncapTmpl - Size of the XfrmEncapTmpl struct, representing encapsulation templates for IPsec.
	SizeOfXfrmEncapTmpl = int(unsafe.Sizeof(XfrmEncapTmpl{}))

	// SizeOfXfrmId - Size of the XfrmId struct, representing the identifier for IPsec Security Associations.
	SizeOfXfrmId = int(unsafe.Sizeof(XfrmId{}))

	// SizeOfXfrmLifetimeCfg - Size of the XfrmLifetimeCfg struct, representing lifetime configuration for IPsec SAs.
	SizeOfXfrmLifetimeCfg = int(unsafe.Sizeof(XfrmLifetimeCfg{}))

	// SizeOfXfrmLifetimeCur - Size of the XfrmLifetimeCur struct, representing current lifetime usage of IPsec SAs.
	SizeOfXfrmLifetimeCur = int(unsafe.Sizeof(XfrmLifetimeCur{}))

	// SizeOfXfrmMark - Size of the XfrmMark struct, used for packet classification in IPsec SAs.
	SizeOfXfrmMark = int(unsafe.Sizeof(XfrmMark{}))

	// SizeOfXfrmUserPolicyId - Size of the XfrmUserpolicyId struct
	SizeOfXfrmUserPolicyId = int(unsafe.Sizeof(XfrmUserPolicyId{}))

	// SizeOfXfrmUserPolicyInfo - Size of the XfrmUserPolicyInfo struct
	SizeOfXfrmUserPolicyInfo = int(unsafe.Sizeof(XfrmUserPolicyInfo{}))

	// SizeOfXfrmReplayState - Size of the XfrmReplayState struct, representing replay state for IPsec SAs.
	SizeOfXfrmReplayState = int(unsafe.Sizeof(XfrmReplayState{}))

	// SizeOfXfrmReplayStateEsn - Size of XfrmReplayStateEsn struct, representing replay state with extended
	// sequence numbers.
	SizeOfXfrmReplayStateEsn = int(unsafe.Sizeof(XfrmReplayStateEsn{}))

	// SizeOfXfrmSelector defines the size of the XfrmSelector structure.
	SizeOfXfrmSelector = int(unsafe.Sizeof(XfrmSelector{}))

	// SizeOfXfrmStats - Size of the XfrmStats struct, representing IPsec framework statistics.
	SizeOfXfrmStats = int(unsafe.Sizeof(XfrmStats{}))

	// SizeOfXfrmUserExpire - Size of the XfrmUserExpire struct, used for XFRM (IPsec framework) expiration events.
	SizeOfXfrmUserExpire = int(unsafe.Sizeof(XfrmUserExpire{}))

	// SizeOfXfrmUserSaFlush - Size of the XfrmUserSaFlush struct, representing requests to flush
	// Security Associations in XFRM.
	SizeOfXfrmUserSaFlush = int(unsafe.Sizeof(XfrmUserSaFlush{}))

	// SizeOfXfrmUserSaId - Size of the XfrmUserSaId struct, representing user-level Security Association identifiers.
	SizeOfXfrmUserSaId = int(unsafe.Sizeof(XfrmUserSaId{}))

	// SizeOfXfrmUserSaInfo - Size of the XfrmUserSaInfo struct, representing user-level Security Association
	// information.
	SizeOfXfrmUserSaInfo = int(unsafe.Sizeof(XfrmUserSaInfo{}))

	// SizeOfXfrmUserSpiInfo - Size of the XfrmUserSpiInfo struct, used for Security Parameter Index (SPI) information.
	SizeOfXfrmUserSpiInfo = int(unsafe.Sizeof(XfrmUserSpiInfo{}))

	SizeOfXfrmUserTmpl = int(unsafe.Sizeof(XfrmUserTmpl{}))
)
