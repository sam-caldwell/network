package core

import "fmt"

// String - Returns a string representation of the BridgeVlanInfo struct.
// It outputs the field names and their values in a key-value format, making it easier to debug and inspect
// the content of the BridgeVlanInfo struct.
//
// Example output:
// BridgeVlanInfo{Flags: 0x0100, Vid: 100}
//
// This method is useful when logging or printing the contents of a BridgeVlanInfo object.
func (bridge *BridgeVlanInfo) String() string {

	return fmt.Sprintf("BridgeVlanInfo{Flags: 0x%04x, Vid: %d}", bridge.Flags, bridge.Vid)

}
