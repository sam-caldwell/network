package core

import (
	"errors"
)

// VlanIdType represents a VLAN ID (Virtual LAN Identifier) as defined in IEEE 802.1Q.
// VLAN IDs are used to identify and separate network traffic at the data link layer (Layer 2).
// The valid range for VLAN IDs is 1 to 4094, as per the IEEE 802.1Q standard.
// Reference: IEEE 802.1Q (https://standards.ieee.org/standard/802_1Q-2018.html)
type VlanIdType uint16

// ToInt converts the VlanIdType to an int type, which may be useful in contexts where
// integer-based operations or comparisons are needed.
//
// Returns:
// - int: The VLAN ID as an integer.
func (id *VlanIdType) ToInt() int {
	return int(*id)
}

// FromInt sets the VlanIdType from an integer value. It validates the VLAN ID to ensure
// it is within the valid range, which is between 1 and 4094, according to the IEEE 802.1Q standard.
//
// If the integer value is outside the valid range, an error is returned.
//
// Arguments:
// - i (int): The integer value to set as the VLAN ID.
//
// Returns:
// - error: An error is returned if the integer value is outside the valid VLAN ID range (1-4094).
//
// Reference:
// - IEEE 802.1Q Standard: VLAN IDs range from 1 to 4094.
// - VLAN ID 0 is reserved for priority tagging and VLAN ID 4095 is reserved.
func (id *VlanIdType) FromInt(i int) error {
	if i < 1 || i > 4094 {
		return errors.New("invalid VLAN ID")
	}
	*id = VlanIdType(i)
	return nil
}
