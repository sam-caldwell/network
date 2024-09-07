package core

import (
	"github.com/sam-caldwell/convert"
	"github.com/sam-caldwell/network"
)

// SetDstPort only tcp and udp are supported to set port
func (p *TcPedit) SetDstPort(dstPort network.PortNumber, protocol IpProtocol) {
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

	tKey.Val = uint32(convert.Uint16ToBigEndian(uint16(dstPort))) << 16
	tKey.Mask = 0x0000ffff
	p.Keys = append(p.Keys, tKey)
	p.KeysEx = append(p.KeysEx, tKeyEx)
	p.Sel.NKeys++
}
