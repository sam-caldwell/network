package core

import "net"

// SetIPv6Dst sets the IPv6 destination address for the pedit action.
// This function modifies the IPv6 destination address by splitting the 16-byte address into 4 parts,
// each represented by 4 bytes (uint32). It then creates pedit keys and extended pedit keys for
// each part and sets the appropriate offset and command.
//
// Parameters:
//
//	ip6 - The IPv6 destination address to set (net.IP).
func (p *TcPedit) SetIPv6Dst(ip6 net.IP) {
	// Extract the first 4 bytes (0-3) of the IPv6 address as uint32.
	u32 := NativeEndian.Uint32(ip6[:4])

	// Create a new TcPeditKey and TcPeditKeyEx for the first 4 bytes.
	tKey := TcPeditKey{}
	tKeyEx := TcPeditKeyEx{}

	// Set the value for the first 4 bytes.
	tKey.Val = u32

	// Set the offset to 24, indicating the position of the destination address in the IPv6 header.
	tKey.Off = 24

	// Set the extended key type to IPv6 and command to "Set".
	tKeyEx.HeaderType = TcaPeditKeyExHdrTypeIP6
	tKeyEx.Cmd = PeditCmdSet

	// Append the new key and extended key to the TcPedit struct.
	p.Keys = append(p.Keys, tKey)
	p.KeysEx = append(p.KeysEx, tKeyEx)

	// Increment the number of keys in the TcPedit selection.
	p.Sel.NKeys++

	// Repeat the same steps for the remaining 12 bytes (divided into 4-byte segments).

	// Extract the second 4 bytes (4-7).
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

	// Extract the third 4 bytes (8-11).
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

	// Extract the fourth 4 bytes (12-15).
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
