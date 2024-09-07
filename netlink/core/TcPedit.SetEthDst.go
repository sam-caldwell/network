package core

import "net"

// SetEthDst sets the Ethernet destination MAC address for the pedit action.
// This function modifies packet headers by creating pedit keys and extended pedit keys
// for the Ethernet destination address (MAC). It splits the MAC address into two parts
// (first 4 bytes as uint32 and last 2 bytes as uint16) and sets the appropriate values
// and commands for each part.
//
// Parameters:
//
//	mac - The Ethernet destination MAC address to set (net.HardwareAddr).
func (p *TcPedit) SetEthDst(mac net.HardwareAddr) {
	// Convert the first 4 bytes of the MAC address to uint32.
	u32 := NativeEndian.Uint32(mac)

	// Convert the last 2 bytes of the MAC address to uint16.
	u16 := NativeEndian.Uint16(mac[4:])

	// Create a new TcPeditKey and TcPeditKeyEx for the first 4 bytes.
	tKey := TcPeditKey{}
	tKeyEx := TcPeditKeyEx{}

	// Set the value for the first 4 bytes of the MAC address.
	tKey.Val = u32

	// Set the extended key type to Ethernet and command to "Set".
	tKeyEx.HeaderType = TcaPeditKeyExHdrTypeEth
	tKeyEx.Cmd = PeditCmdSet

	// Append the new key and extended key to the TcPedit struct.
	p.Keys = append(p.Keys, tKey)
	p.KeysEx = append(p.KeysEx, tKeyEx)

	// Increment the number of keys in the TcPedit selection.
	p.Sel.NKeys++

	// Create a new TcPeditKey and TcPeditKeyEx for the last 2 bytes.
	tKey = TcPeditKey{}
	tKeyEx = TcPeditKeyEx{}

	// Set the value for the last 2 bytes of the MAC address, cast to uint32.
	tKey.Val = uint32(u16)

	// Set the mask to apply only to the last 2 bytes (0xffff0000).
	tKey.Mask = 0xffff0000

	// Set the offset for the last 2 bytes to 4 (starting at byte 5 in the MAC address).
	tKey.Off = 4

	// Set the extended key type to Ethernet and command to "Add".
	tKeyEx.HeaderType = TcaPeditKeyExHdrTypeEth
	tKeyEx.Cmd = PeditCmdAdd

	// Append the new key and extended key to the TcPedit struct.
	p.Keys = append(p.Keys, tKey)
	p.KeysEx = append(p.KeysEx, tKeyEx)

	// Increment the number of keys in the TcPedit selection.
	p.Sel.NKeys++
}
