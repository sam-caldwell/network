package core

// MacVlanMacAddrMode -
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
type MacVlanMacAddrMode uint8

const (

	// MacvlanMacaddrAdd - MACVLAN_MACADDR_ADD -
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	MacvlanMacaddrAdd MacVlanMacAddrMode = 0

	// MacvlanMacaddrDel - MACVLAN_MACADDR_DEL -
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	MacvlanMacaddrDel MacVlanMacAddrMode = 1

	// MacvlanMacaddrFlush - MACVLAN_MACADDR_FLUSH -
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	MacvlanMacaddrFlush MacVlanMacAddrMode = 2

	// MacvlanMacaddrSet - MACVLAN_MACADDR_SET -
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	MacvlanMacaddrSet MacVlanMacAddrMode = 3
)
