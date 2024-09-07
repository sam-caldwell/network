package core

// TcPedit represents the pedit action used in the Linux kernel's traffic control (TC).
// The pedit action allows packet header modifications by specifying a selection (Sel),
// a set of keys (Keys) for basic operations, and extended keys (KeysEx) for more granular control.
//
// The `Extend` field can be used to enable additional functionality if necessary.
//
// For more details, refer to the Linux Kernel Source:
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_pedit.h
type TcPedit struct {
	// Sel defines the selection criteria for the pedit operation, including the number of keys and flags.
	Sel TcPeditSel

	// Keys is the set of keys used to apply basic modifications to the packet.
	Keys []TcPeditKey

	// KeysEx is the set of extended keys that allow for more specific modifications, including commands and header types.
	KeysEx []TcPeditKeyEx

	// Extend is a flag that can be used to extend the pedit operation or enable additional features.
	Extend uint8
}

// SetDstPort only tcp and udp are supported to set port
func (p *TcPedit) SetDstPort(dstPort IpProtocol, protocol IpProtocol) {
	tKey := TcPeditKey{}
	tKeyEx := TcPeditKeyEx{}

	switch protocol {
	case IpProtoTCP:
		tKeyEx.HeaderType = TcaPeditKeyExHdrTypeTCP
	case IpProtoUDP:
		tKeyEx.HeaderType = TcaPeditKeyExHdrTypeUDP
	default:
		return
	}

	tKeyEx.Cmd = PeditCmdSet

	tKey.Val = uint32(ConvertUint16ToBigEndian(uint16(dstPort))) << 16
	tKey.Mask = 0x0000ffff
	p.Keys = append(p.Keys, tKey)
	p.KeysEx = append(p.KeysEx, tKeyEx)
	p.Sel.NKeys++
}

// SetSrcPort only tcp and udp are supported to set port
func (p *TcPedit) SetSrcPort(srcPort IpProtocol, protocol IpProtocol) {
	tKey := TcPeditKey{}
	tKeyEx := TcPeditKeyEx{}

	switch protocol {
	case IpProtoTCP:
		tKeyEx.HeaderType = TcaPeditKeyExHdrTypeTCP
	case IpProtoUDP:
		tKeyEx.HeaderType = TcaPeditKeyExHdrTypeUDP
	default:
		return
	}

	tKeyEx.Cmd = PeditCmdSet

	tKey.Val = uint32(ConvertUint16ToBigEndian(uint16(srcPort)))
	tKey.Mask = 0xffff0000
	p.Keys = append(p.Keys, tKey)
	p.KeysEx = append(p.KeysEx, tKeyEx)
	p.Sel.NKeys++
}
