package core

import (
	"golang.org/x/sys/unix"
	"unsafe"
)

const (
	// SizeOfIfAddressMessage     = 0x8 // bytes as derived from unix.SizeOfIfAddressMessage
	SizeOfIfAddressMessage = int(unsafe.Sizeof(IfAddressMessage{}))

	// SizeOfIfaCacheinfo - 0x16 // bytes as derived from unix.SizeofIfaCacheinfo
	SizeOfIfaCacheinfo = int(unsafe.Sizeof(IfaCacheInfo{}))

	// SizeofBridgeVlanInfo - size of BridgeVlanInfo
	SizeofBridgeVlanInfo = int(unsafe.Sizeof(BridgeVlanInfo{}))

	// SizeofCnMsgOp - size of CnMsgOp struct
	SizeofCnMsgOp = int(unsafe.Sizeof(CnMsgOp{}))

	// SizeofGenlmsg - size of Genlmsg struct
	SizeofGenlmsg = int(unsafe.Sizeof(Genlmsg{}))

	// SizeofNfgenmsg - Track the message sizes for the correct serialization/deserialization
	SizeofNfgenmsg = int(unsafe.Sizeof(Nfgenmsg{}))

	// SizeofNlMsghdr - Size of unix.NlMsghdr struct
	SizeofNlMsghdr = int(unsafe.Sizeof(unix.NlMsghdr{}))

	// SizeofRtNexthop - Size of the RtNexthop struct
	SizeofRtNexthop = int(unsafe.Sizeof(RtNexthop{}))

	// SizeOfTcNetemRate - Size of the TcNetemRate struct, defining rate options in netem.
	SizeOfTcNetemRate = int(unsafe.Sizeof(TcNetemRate{}))

	// SizeofTcActionMsg - Size of the TcActionMsg struct, which represents actions taken on packets in TC.
	SizeofTcActionMsg = int(unsafe.Sizeof(TcActionMsg{}))

	// SizeofTcConnmark - Size of the TcConnmark struct, used for connection marking in TC.
	SizeofTcConnmark = int(unsafe.Sizeof(TcConnmark{}))

	// SizeofTcCsum - Size of the TcCsum struct, representing checksum actions in traffic control.
	SizeofTcCsum = int(unsafe.Sizeof(TcCsum{}))

	// SizeofTcGen - Size of the TcGen struct, representing general traffic control options.
	SizeofTcGen = int(unsafe.Sizeof(TcGen{}))

	// SizeofTcHtbCopt - Size of the TcHtbCopt struct, representing options for hierarchical token bucket (HTB)
	// classes in TC.
	SizeofTcHtbCopt = int(unsafe.Sizeof(TcHtbCopt{}))

	// SizeofTcHtbGlob - Size of the TcHtbGlob struct, representing global parameters for HTB.
	SizeofTcHtbGlob = int(unsafe.Sizeof(TcHtbGlob{}))

	// SizeofIfInfoMsg -  Size of the IfInfoMsg struct
	SizeofIfInfoMsg = int(unsafe.Sizeof(IfInfoMsg{}))

	// SizeofTcMsg - Size of the TcMsg struct, which is used to configure traffic control (TC) in Linux.
	SizeofTcMsg = int(unsafe.Sizeof(TcMsg{}))

	// SizeofTcMirred - Size of the TcMirred struct, representing "mirred" (mirror/redirect) actions in TC.
	SizeofTcMirred = int(unsafe.Sizeof(TcMirred{}))

	// SizeofTcNetemCorr - Size of the TcNetemCorr struct, representing correlation options in netem.
	SizeofTcNetemCorr = int(unsafe.Sizeof(TcNetemCorr{}))

	// SizeofTcNetemCorrupt - Size of the TcNetemCorrupt struct, representing options for packet corruption in netem.
	SizeofTcNetemCorrupt = int(unsafe.Sizeof(TcNetemCorrupt{}))

	// SizeofTcNetemQopt - Size of TcNetemQopt struct, used for network emulation (netem) options in traffic control.
	SizeofTcNetemQopt = int(unsafe.Sizeof(TcNetemQopt{}))

	// SizeofTcNetemReorder - Size of the TcNetemReorder struct, used to configure packet reordering in netem.
	SizeofTcNetemReorder = int(unsafe.Sizeof(TcNetemReorder{}))

	// SizeOfTcPeditKey - size of the TcPeditKey struct
	SizeOfTcPeditKey = int(unsafe.Sizeof(TcPeditKey{}))

	// SizeOfTcPeditSel - size of the TcPeditSel struct
	SizeOfTcPeditSel = int(unsafe.Sizeof(TcPeditSel{}))

	// SizeofTcPolice - Size of the TcPolice struct, representing police actions in TC.
	SizeofTcPolice = int(unsafe.Sizeof(TcPolice{}))

	// SizeofTcPrioMap - Size of the TcPrioMap struct, used to map priorities for traffic control.
	SizeofTcPrioMap = int(unsafe.Sizeof(TcPrioMap{}))

	// SizeofTcRateSpec - Size of the TcRateSpec struct, which defines the rate specification for traffic control.
	SizeofTcRateSpec = int(unsafe.Sizeof(TcRateSpec{}))

	// SizeofTcSfqQopt - Size of the TcSfqQopt struct, representing options for SFQ.
	SizeofTcSfqQopt = int(unsafe.Sizeof(TcSfqQopt{}))

	// SizeofTcSfqQoptV1 - Size of TcSfqQoptV1 struct, representing options for stochastic fair queuing (SFQ) version 1.
	SizeofTcSfqQoptV1 = int(unsafe.Sizeof(TcSfqQoptV1{}))

	// SizeofTcSfqRedStats - Size of TcSfqRedStats struct, representing RED (random early detection) statistics in SFQ.
	SizeofTcSfqRedStats = int(unsafe.Sizeof(TcSfqRedStats{}))

	// SizeofTcSkbEdit - Size of the TcSkbEdit struct, representing modifications to skb (socket buffer) in TC.
	SizeofTcSkbEdit = int(unsafe.Sizeof(TcSkbEdit{}))

	// SizeofTcTbfQopt - Size of TcTbfQopt struct, represents options for token bucket filter (TBF) in traffic control.
	SizeofTcTbfQopt = int(unsafe.Sizeof(TcTbfQopt{}))

	// SizeofTcTunnelKey - Size of the TcTunnelKey struct, representing tunnel key actions in TC.
	SizeofTcTunnelKey = int(unsafe.Sizeof(TcTunnelKey{}))

	// SizeofTcf - Size of the Tcf struct, representing a filter in traffic control.
	SizeofTcf = int(unsafe.Sizeof(Tcf{}))

	// SizeofUint32Bitfield - The size of a 32-bit bitfield used in various traffic control structs.
	SizeofUint32Bitfield = int(unsafe.Sizeof(Uint32Bitfield{}))

	// SizeOfUnixRtAttr - The size of a unix.RtAttr struct
	SizeOfUnixRtAttr = int(unsafe.Sizeof(unix.RtAttr{}))

	// SizeofVfGUID - Size of VfGUID struct
	SizeofVfGUID = int(unsafe.Sizeof(VfGUID{}))

	// SizeofVfLinkState - Size of VfLinkState struct
	SizeofVfLinkState = int(unsafe.Sizeof(VfLinkState{}))

	// SizeofVfMac - Size of VfMac struct
	SizeofVfMac = int(unsafe.Sizeof(VfMac{}))

	// SizeofVfRate - Size of VfRate struct
	SizeofVfRate = int(unsafe.Sizeof(VfRate{}))

	// SizeofVfRssQueryEn - Size of VfRssQueryEn struct
	SizeofVfRssQueryEn = int(unsafe.Sizeof(VfRssQueryEn{}))

	// SizeofVfSpoofchk - Size of VfSpoofchk
	SizeofVfSpoofchk = int(unsafe.Sizeof(VfSpoofchk{}))

	// SizeofVfTrust - Size of VfTrust struct
	SizeofVfTrust = int(unsafe.Sizeof(VfTrust{}))

	// SizeofVfTxRate - Size of VfTxRate struct
	SizeofVfTxRate = int(unsafe.Sizeof(VfTxRate{}))

	// SizeofVfVlan - Size of VfVlan struct
	SizeofVfVlan = int(unsafe.Sizeof(VfVlan{}))

	// SizeofVfVlanInfo - Size of IfLaVfVlanInfoStruct struct
	SizeofVfVlanInfo = int(unsafe.Sizeof(IfLaVfVlanInfoStruct{}))

	// SizeofXfrmAddress - Size of the XfrmAddress struct, used to represent addresses in XFRM (IPsec framework).
	// Assumes largest size (ipV6)
	SizeofXfrmAddress = int(4 * 32)

	// SizeofXfrmAlgo - Size of the XfrmAlgo struct, representing cryptographic algorithms for IPsec SAs.
	SizeofXfrmAlgo = int(unsafe.Sizeof(XfrmAlgo{}))

	// SizeofXfrmAlgoAEAD - Size of XfrmAlgoAEAD struct, representing AEAD (authenticated encryption) algorithms
	// for IPsec.
	SizeofXfrmAlgoAEAD = int(unsafe.Sizeof(XfrmAlgoAEAD{}))

	// SizeofXfrmAlgoAuth - Size of the XfrmAlgoAuth struct, representing authentication algorithms for IPsec SAs.
	SizeofXfrmAlgoAuth = int(unsafe.Sizeof(XfrmAlgoAuth{}))

	// SizeofXfrmEncapTmpl - Size of the XfrmEncapTmpl struct, representing encapsulation templates for IPsec.
	SizeofXfrmEncapTmpl = int(unsafe.Sizeof(XfrmEncapTmpl{}))

	// SizeofXfrmId - Size of the XfrmId struct, representing the identifier for IPsec Security Associations.
	SizeofXfrmId = int(unsafe.Sizeof(XfrmId{}))

	// SizeofXfrmLifetimeCfg - Size of the XfrmLifetimeCfg struct, representing lifetime configuration for IPsec SAs.
	SizeofXfrmLifetimeCfg = int(unsafe.Sizeof(XfrmLifetimeCfg{}))

	// SizeofXfrmLifetimeCur - Size of the XfrmLifetimeCur struct, representing current lifetime usage of IPsec SAs.
	SizeofXfrmLifetimeCur = int(unsafe.Sizeof(XfrmLifetimeCur{}))

	// SizeofXfrmMark - Size of the XfrmMark struct, used for packet classification in IPsec SAs.
	SizeofXfrmMark = int(unsafe.Sizeof(XfrmMark{}))

	// SizeofXfrmUserPolicyId - Size of the XfrmUserpolicyId struct
	SizeofXfrmUserPolicyId = int(unsafe.Sizeof(XfrmUserPolicyId{}))

	// SizeofXfrmUserPolicyInfo - Size of the XfrmUserPolicyInfo struct
	SizeofXfrmUserPolicyInfo = int(unsafe.Sizeof(XfrmUserPolicyInfo{}))

	// SizeofXfrmReplayState - Size of the XfrmReplayState struct, representing replay state for IPsec SAs.
	SizeofXfrmReplayState = int(unsafe.Sizeof(XfrmReplayState{}))

	// SizeofXfrmReplayStateEsn - Size of XfrmReplayStateEsn struct, representing replay state with extended
	// sequence numbers.
	SizeofXfrmReplayStateEsn = int(unsafe.Sizeof(XfrmReplayStateEsn{}))

	// SizeOfXfrmSelector defines the size of the XfrmSelector structure.
	SizeOfXfrmSelector = int(unsafe.Sizeof(XfrmSelector{}))

	// SizeofXfrmStats - Size of the XfrmStats struct, representing IPsec framework statistics.
	SizeofXfrmStats = int(unsafe.Sizeof(XfrmStats{}))

	// SizeofXfrmUserExpire - Size of the XfrmUserExpire struct, used for XFRM (IPsec framework) expiration events.
	SizeofXfrmUserExpire = int(unsafe.Sizeof(XfrmUserExpire{}))

	// SizeofXfrmUsersaFlush - Size of the XfrmUsersaFlush struct, representing requests to flush
	// Security Associations in XFRM.
	SizeofXfrmUsersaFlush = int(unsafe.Sizeof(XfrmUsersaFlush{}))

	// SizeofXfrmUsersaId - Size of the XfrmUsersaId struct, representing user-level Security Association identifiers.
	SizeofXfrmUsersaId = int(unsafe.Sizeof(XfrmUsersaId{}))

	// SizeofXfrmUsersaInfo - Size of the XfrmUsersaInfo struct, representing user-level Security Association
	// information.
	SizeofXfrmUsersaInfo = int(unsafe.Sizeof(XfrmUsersaInfo{}))

	// SizeofXfrmUserSpiInfo - Size of the XfrmUserSpiInfo struct, used for Security Parameter Index (SPI) information.
	SizeofXfrmUserSpiInfo = int(unsafe.Sizeof(XfrmUserSpiInfo{}))

	SizeofXfrmUserTmpl = int(unsafe.Sizeof(XfrmUserTmpl{}))
)
