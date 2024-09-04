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

	//SizeofIfInfoMsg -  0x16 // bytes as derived from unix.SizeofIfInfomsg
	SizeofIfInfoMsg = int(unsafe.Sizeof(IfInfoMsg{}))

	// SizeofCnMsgOp - size of CnMsgOp struct
	SizeofCnMsgOp = int(unsafe.Sizeof(CnMsgOp{}))

	// SizeofBridgeVlanInfo - size of BridgeVlanInfo
	SizeofBridgeVlanInfo = int(unsafe.Sizeof(BridgeVlanInfo{}))

	// SizeofNfgenmsg - Track the message sizes for the correct serialization/deserialization
	SizeofNfgenmsg = int(unsafe.Sizeof(Nfgenmsg{}))

	// SizeofVfMac - Size of VfMac struct
	SizeofVfMac = int(unsafe.Sizeof(VfMac{}))

	// SizeofVfVlan - Size of VfVlan struct
	SizeofVfVlan = int(unsafe.Sizeof(VfVlan{}))

	// SizeofVfVlanInfo - Size of IfLaVfVlanInfoStruct struct
	SizeofVfVlanInfo = int(unsafe.Sizeof(IfLaVfVlanInfoStruct{}))

	// SizeofNlMsghdr - Size of unix.NlMsghdr struct
	SizeofNlMsghdr = int(unsafe.Sizeof(unix.NlMsghdr{}))

	// SizeofUnixRtAttr - Size of unix.RtAttr struct (note this is smaller than our core.RtAttr which wraps and extends
	// the unix.RtAttr.
	SizeofUnixRtAttr = int(unsafe.Sizeof(unix.RtAttr{}))

	//SizeofGenlmsg - size of Genlmsg struct
	SizeofGenlmsg = int(unsafe.Sizeof(Genlmsg{}))

	// SizeofVfRate - Size of VfRate struct
	SizeofVfRate = int(unsafe.Sizeof(VfRate{}))

	// SizeofVfSpoofchk - Size of VfSpoofchk
	SizeofVfSpoofchk = int(unsafe.Sizeof(VfSpoofchk{}))

	// SizeofVfTxRate - Size of VfTxRate struct
	SizeofVfTxRate = int(unsafe.Sizeof(VfTxRate{}))

	// SizeofVfLinkState - Size of VfLinkState struct
	SizeofVfLinkState = int(unsafe.Sizeof(VfLinkState{}))

	// SizeofVfRssQueryEn - Size of VfRssQueryEn struct
	SizeofVfRssQueryEn = int(unsafe.Sizeof(VfRssQueryEn{}))

	// SizeofVfTrust - Size of VfTrust struct
	SizeofVfTrust = int(unsafe.Sizeof(VfTrust{}))

	// SizeofVfGUID - Size of VfGUID struct
	SizeofVfGUID = int(unsafe.Sizeof(VfGUID{}))

	// SizeofNfattr -
	SizeofNfattr = 4

	// SizeofNfConntrack -
	SizeofNfConntrack = 376

	// SizeofNfctTupleHead -
	SizeofNfctTupleHead = 52

	// SizeofRtNexthop -
	SizeofRtNexthop = int(unsafe.Sizeof(RtNexthop{}))
)
