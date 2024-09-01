package core

import "unsafe"

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

	// SizeofNfattr -
	SizeofNfattr = 4

	// SizeofNfConntrack -
	SizeofNfConntrack = 376

	// SizeofNfctTupleHead -
	SizeofNfctTupleHead = 52
)
