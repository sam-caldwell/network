package core

import "net"

func (p *TcPedit) SetIPv6Dst(ip6 net.IP) {
	u32 := NativeEndian.Uint32(ip6[:4])

	tKey := TcPeditKey{}
	tKeyEx := TcPeditKeyEx{}

	tKey.Val = u32
	tKey.Off = 24
	tKeyEx.HeaderType = TcaPeditKeyExHdrTypeIP6
	tKeyEx.Cmd = PeditCmdSet

	p.Keys = append(p.Keys, tKey)
	p.KeysEx = append(p.KeysEx, tKeyEx)
	p.Sel.NKeys++

	u32 = NativeEndian.Uint32(ip6[4:8])
	tKey = TcPeditKey{}
	tKeyEx = TcPeditKeyEx{}

	tKey.Val = u32
	tKey.Off = 28
	tKeyEx.HeaderType = TcaPeditKeyExHdrTypeIP6
	tKeyEx.Cmd = PeditCmdSet

	p.Keys = append(p.Keys, tKey)
	p.KeysEx = append(p.KeysEx, tKeyEx)

	p.Sel.NKeys++

	u32 = NativeEndian.Uint32(ip6[8:12])
	tKey = TcPeditKey{}
	tKeyEx = TcPeditKeyEx{}

	tKey.Val = u32
	tKey.Off = 32
	tKeyEx.HeaderType = TcaPeditKeyExHdrTypeIP6
	tKeyEx.Cmd = PeditCmdSet

	p.Keys = append(p.Keys, tKey)
	p.KeysEx = append(p.KeysEx, tKeyEx)

	p.Sel.NKeys++

	u32 = NativeEndian.Uint32(ip6[12:16])
	tKey = TcPeditKey{}
	tKeyEx = TcPeditKeyEx{}

	tKey.Val = u32
	tKey.Off = 36
	tKeyEx.HeaderType = TcaPeditKeyExHdrTypeIP6
	tKeyEx.Cmd = PeditCmdSet

	p.Keys = append(p.Keys, tKey)
	p.KeysEx = append(p.KeysEx, tKeyEx)

	p.Sel.NKeys++
}
