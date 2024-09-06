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

	// SizeOfTcNetemRate is the size of the TcNetemRate structure, defining rate options in netem.
	SizeOfTcNetemRate = int(unsafe.Sizeof(TcNetemRate{}))

	// SizeofTcActionMsg is the size of the TcActionMsg structure, which represents actions taken on packets in TC.
	SizeofTcActionMsg = int(unsafe.Sizeof(TcActionMsg{}))

	// SizeofTcConnmark is the size of the TcConnmark structure, used for connection marking in TC.
	SizeofTcConnmark = int(unsafe.Sizeof(TcConnmark{}))

	// SizeofTcCsum is the size of the TcCsum structure, representing checksum actions in traffic control.
	SizeofTcCsum = int(unsafe.Sizeof(TcCsum{}))

	// SizeofTcGen is the size of the TcGen structure, representing general traffic control options.
	SizeofTcGen = int(unsafe.Sizeof(TcGen{}))

	// SizeofTcHtbCopt is the size of the TcHtbCopt structure, representing options for hierarchical token bucket (HTB) classes in TC.
	SizeofTcHtbCopt = int(unsafe.Sizeof(TcHtbCopt{}))

	// SizeofTcHtbGlob is the size of the TcHtbGlob structure, representing global parameters for HTB.
	SizeofTcHtbGlob = int(unsafe.Sizeof(TcHtbGlob{}))

	// SizeofTcInfoMsg -  0x16 // bytes as derived from unix.SizeofIfInfomsg
	SizeofIfInfoMsg = int(unsafe.Sizeof(IfInfoMsg{}))

	// SizeofTcMsg is the size of the TcMsg structure, which is used to configure traffic control (TC) in Linux.
	SizeofTcMsg = int(unsafe.Sizeof(TcMsg{}))

	// SizeofTcMirred is the size of the TcMirred structure, representing "mirred" (mirror/redirect) actions in TC.
	SizeofTcMirred = int(unsafe.Sizeof(TcMirred{}))

	// SizeofTcNetemCorr is the size of the TcNetemCorr structure, representing correlation options in netem.
	SizeofTcNetemCorr = int(unsafe.Sizeof(TcNetemCorr{}))

	// SizeofTcNetemCorrupt is the size of the TcNetemCorrupt structure, representing options for packet corruption in netem.
	SizeofTcNetemCorrupt = int(unsafe.Sizeof(TcNetemCorrupt{}))

	// SizeofTcNetemQopt is the size of the TcNetemQopt structure, used for network emulation (netem) options in traffic control.
	SizeofTcNetemQopt = int(unsafe.Sizeof(TcNetemQopt{}))

	// SizeofTcNetemReorder is the size of the TcNetemReorder structure, used to configure packet reordering in netem.
	SizeofTcNetemReorder = int(unsafe.Sizeof(TcNetemReorder{}))

	// SizeofTcPolice is the size of the TcPolice structure, representing police actions in TC.
	SizeofTcPolice = int(unsafe.Sizeof(TcPolice{}))

	// SizeofTcPrioMap is the size of the TcPrioMap structure, used to map priorities for traffic control.
	SizeofTcPrioMap = int(unsafe.Sizeof(TcPrioMap{}))

	// SizeofTcRateSpec is the size of the TcRateSpec structure, which defines the rate specification for traffic control.
	SizeofTcRateSpec = int(unsafe.Sizeof(TcRateSpec{}))

	// SizeofTcSfqQopt is the size of the TcSfqQopt structure, representing options for SFQ.
	SizeofTcSfqQopt = int(unsafe.Sizeof(TcSfqQopt{}))

	// SizeofTcSfqQoptV1 is the size of the TcSfqQoptV1 structure, representing options for stochastic fair queuing (SFQ) version 1.
	SizeofTcSfqQoptV1 = int(unsafe.Sizeof(TcSfqQoptV1{}))

	// SizeofTcSfqRedStats is the size of the TcSfqRedStats structure, representing RED (random early detection) statistics in SFQ.
	SizeofTcSfqRedStats = int(unsafe.Sizeof(TcSfqRedStats{}))

	// SizeofTcSkbEdit is the size of the TcSkbEdit structure, representing modifications to skb (socket buffer) in TC.
	SizeofTcSkbEdit = int(unsafe.Sizeof(TcSkbEdit{}))

	// SizeofTcTbfQopt is the size of the TcTbfQopt structure, which represents options for token bucket filter (TBF) in traffic control.
	SizeofTcTbfQopt = int(unsafe.Sizeof(TcTbfQopt{}))

	// SizeofTcTunnelKey is the size of the TcTunnelKey structure, representing tunnel key actions in TC.
	SizeofTcTunnelKey = int(unsafe.Sizeof(TcTunnelKey{}))

	// SizeofTcf is the size of the Tcf structure, representing a filter in traffic control.
	SizeofTcf = int(unsafe.Sizeof(Tcf{}))

	// SizeofUint32Bitfield is the size of a 32-bit bitfield used in various traffic control structures.
	SizeofUint32Bitfield = int(unsafe.Sizeof(Uint32Bitfield{}))

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

	// SizeofXfrmAddress is the size of the XfrmAddress structure, used to represent addresses in XFRM (IPsec framework).
	SizeofXfrmAddress = int(unsafe.Sizeof(XfrmAddress{}))

	// SizeofXfrmAlgo is the size of the XfrmAlgo structure, representing cryptographic algorithms for IPsec SAs.
	SizeofXfrmAlgo = int(unsafe.Sizeof(XfrmAlgo{}))

	// SizeofXfrmAlgoAEAD is the size of the XfrmAlgoAEAD structure, representing AEAD (authenticated encryption) algorithms for IPsec.
	SizeofXfrmAlgoAEAD = int(unsafe.Sizeof(XfrmAlgoAEAD{}))

	// SizeofXfrmAlgoAuth is the size of the XfrmAlgoAuth structure, representing authentication algorithms for IPsec SAs.
	SizeofXfrmAlgoAuth = int(unsafe.Sizeof(XfrmAlgoAuth{}))

	// SizeofXfrmEncapTmpl is the size of the XfrmEncapTmpl structure, representing encapsulation templates for IPsec.
	SizeofXfrmEncapTmpl = int(unsafe.Sizeof(XfrmEncapTmpl{}))

	// SizeofXfrmId is the size of the XfrmId structure, representing the identifier for IPsec Security Associations.
	SizeofXfrmId = int(unsafe.Sizeof(XfrmId{}))

	// SizeofXfrmLifetimeCfg is the size of the XfrmLifetimeCfg structure, representing lifetime configuration for IPsec SAs.
	SizeofXfrmLifetimeCfg = int(unsafe.Sizeof(XfrmLifetimeCfg{}))

	// SizeofXfrmLifetimeCur is the size of the XfrmLifetimeCur structure, representing current lifetime usage of IPsec SAs.
	SizeofXfrmLifetimeCur = int(unsafe.Sizeof(XfrmLifetimeCur{}))

	// SizeofXfrmMark is the size of the XfrmMark structure, used for packet classification in IPsec SAs.
	SizeofXfrmMark = int(unsafe.Sizeof(XfrmMark{}))

	// SizeofXfrmReplayState is the size of the XfrmReplayState structure, representing replay state for IPsec SAs.
	SizeofXfrmReplayState = int(unsafe.Sizeof(XfrmReplayState{}))

	// SizeofXfrmReplayStateEsn is the size of the XfrmReplayStateEsn structure, representing replay state with extended sequence numbers.
	SizeofXfrmReplayStateEsn = int(unsafe.Sizeof(XfrmReplayStateEsn{}))

	// SizeofXfrmStats is the size of the XfrmStats structure, representing IPsec framework statistics.
	SizeofXfrmStats = int(unsafe.Sizeof(XfrmStats{}))

	// SizeofXfrmUserExpire is the size of the XfrmUserExpire structure, used for XFRM (IPsec framework) expiration events.
	SizeofXfrmUserExpire = int(unsafe.Sizeof(XfrmUserExpire{}))

	// SizeofXfrmUsersaFlush is the size of the XfrmUsersaFlush structure, representing requests to flush Security Associations in XFRM.
	SizeofXfrmUsersaFlush = int(unsafe.Sizeof(XfrmUsersaFlush{}))

	// SizeofXfrmUsersaId is the size of the XfrmUsersaId structure, representing user-level Security Association identifiers.
	SizeofXfrmUsersaId = int(unsafe.Sizeof(XfrmUsersaId{}))

	// SizeofXfrmUsersaInfo is the size of the XfrmUsersaInfo structure, representing user-level Security Association information.
	SizeofXfrmUsersaInfo = int(unsafe.Sizeof(XfrmUsersaInfo{}))

	// SizeofXfrmUserSpiInfo is the size of the XfrmUserSpiInfo structure, used for Security Parameter Index (SPI) information.
	SizeofXfrmUserSpiInfo = int(unsafe.Sizeof(XfrmUserSpiInfo{}))
)
