package core

// MacVlanMode - macvlan_mode -
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
type MacVlanMode uint8

const (
	_ = iota
	// MacvlanModePrivate - MACVLAN_MODE_PRIVATE - don't talk to other mac vlans
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	MacvlanModePrivate MacVlanMode = 1

	// MacvlanModeVepa - MACVLAN_MODE_VEPA - talk to other ports through ext bridge
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	MacvlanModeVepa MacVlanMode = 2

	// MacvlanModeBridge - MACVLAN_MODE_BRIDGE - talk to bridge ports directly
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	MacvlanModeBridge MacVlanMode = 4

	// MacvlanModePassthru - MACVLAN_MODE_PASSTHRU - take over the underlying device
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	MacvlanModePassthru MacVlanMode = 8

	// MacvlanModeSource - MACVLAN_MODE_SOURCE - use source MAC address list to assign
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	MacvlanModeSource MacVlanMode = 16
)
