package core

import "github.com/sam-caldwell/network"

// SetSrcPort only tcp and udp are supported to set port
func (p *TcPedit) SetSrcPort(srcPort network.PortNumber, protocol IpProtocol) {
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
