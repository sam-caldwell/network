package core

import (
	"net"
)

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

func (p *TcPedit) SetIPv4Src(ip net.IP) {
	u32 := NativeEndian.Uint32(ip[:4])

	tKey := TcPeditKey{}
	tKeyEx := TcPeditKeyEx{}

	tKey.Val = u32
	tKey.Off = 12
	tKeyEx.HeaderType = TCA_PEDIT_KEY_EX_HDR_TYPE_IP4
	tKeyEx.Cmd = TCA_PEDIT_KEY_EX_CMD_SET

	p.Keys = append(p.Keys, tKey)
	p.KeysEx = append(p.KeysEx, tKeyEx)
	p.Sel.NKeys++
}

func (p *TcPedit) SetIPv4Dst(ip net.IP) {
	u32 := NativeEndian.Uint32(ip[:4])

	tKey := TcPeditKey{}
	tKeyEx := TcPeditKeyEx{}

	tKey.Val = u32
	tKey.Off = 16
	tKeyEx.HeaderType = TCA_PEDIT_KEY_EX_HDR_TYPE_IP4
	tKeyEx.Cmd = TCA_PEDIT_KEY_EX_CMD_SET

	p.Keys = append(p.Keys, tKey)
	p.KeysEx = append(p.KeysEx, tKeyEx)
	p.Sel.NKeys++
}

// SetDstPort only tcp and udp are supported to set port
func (p *TcPedit) SetDstPort(dstPort uint16, protocol uint8) {
	tKey := TcPeditKey{}
	tKeyEx := TcPeditKeyEx{}

	switch protocol {
	case unix.IPPROTO_TCP:
		tKeyEx.HeaderType = TCA_PEDIT_KEY_EX_HDR_TYPE_TCP
	case unix.IPPROTO_UDP:
		tKeyEx.HeaderType = TCA_PEDIT_KEY_EX_HDR_TYPE_UDP
	default:
		return
	}

	tKeyEx.Cmd = TCA_PEDIT_KEY_EX_CMD_SET

	tKey.Val = uint32(ConvertUint16ToBigEndian(dstPort)) << 16
	tKey.Mask = 0x0000ffff
	p.Keys = append(p.Keys, tKey)
	p.KeysEx = append(p.KeysEx, tKeyEx)
	p.Sel.NKeys++
}

// SetSrcPort only tcp and udp are supported to set port
func (p *TcPedit) SetSrcPort(srcPort uint16, protocol uint8) {
	tKey := TcPeditKey{}
	tKeyEx := TcPeditKeyEx{}

	switch protocol {
	case unix.IPPROTO_TCP:
		tKeyEx.HeaderType = TCA_PEDIT_KEY_EX_HDR_TYPE_TCP
	case unix.IPPROTO_UDP:
		tKeyEx.HeaderType = TCA_PEDIT_KEY_EX_HDR_TYPE_UDP
	default:
		return
	}

	tKeyEx.Cmd = TCA_PEDIT_KEY_EX_CMD_SET

	tKey.Val = uint32(ConvertUint16ToBigEndian(srcPort))
	tKey.Mask = 0xffff0000
	p.Keys = append(p.Keys, tKey)
	p.KeysEx = append(p.KeysEx, tKeyEx)
	p.Sel.NKeys++
}
