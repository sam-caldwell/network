package core

// DevlinkPortFlavourEnum - devlink_port_flavour
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
type DevlinkPortFlavourEnum uint8

const (
	// DevlinkPortFlavourPhysical - DEVLINK_PORT_FLAVOUR_PHYSICAL - represents any kind of port physically
	// facing the user.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkPortFlavourPhysical DevlinkPortFlavourEnum = 0

	// DevlinkPortFlavourCpu - DEVLINK_PORT_FLAVOUR_CPU - CPU port
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkPortFlavourCpu DevlinkPortFlavourEnum = 1

	// DevlinkPortFlavourDsa - DEVLINK_PORT_FLAVOUR_DSA - Distributed switch architecture interconnect port.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkPortFlavourDsa DevlinkPortFlavourEnum = 2

	// DevlinkPortFlavourPciPf - DEVLINK_PORT_FLAVOUR_PCI_PF - Represents eswitch port for the PCI PF.
	// It is an internal port that faces the PCI PF.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkPortFlavourPciPf DevlinkPortFlavourEnum = 3

	// DevlinkPortFlavourPciVf - DEVLINK_PORT_FLAVOUR_PCI_VF - Represents eswitch port for the PCI VF. It is an
	// internal port that faces the PCI VF.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkPortFlavourPciVf DevlinkPortFlavourEnum = 4

	// DevlinkPortFlavourVirtual - DEVLINK_PORT_FLAVOUR_VIRTUAL - Any virtual port facing the user.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkPortFlavourVirtual DevlinkPortFlavourEnum = 5

	// DevlinkPortFlavourUnused - DEVLINK_PORT_FLAVOUR_UNUSED - Port which exists in the switch, but is not used
	// in any way.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkPortFlavourUnused DevlinkPortFlavourEnum = 6

	// DevlinkPortFlavourPciSf - DEVLINK_PORT_FLAVOUR_PCI_SF - Represents eswitch port for the PCI SF. It is an
	// internal port that faces the PCI SF.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
	DevlinkPortFlavourPciSf DevlinkPortFlavourEnum = 7
)
