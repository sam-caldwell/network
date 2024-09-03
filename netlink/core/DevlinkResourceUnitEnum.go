package core

// DevlinkResourceUnitEnum - devlink_resource_unit - represents the devlink resource unit type.
//
// The devlink_resource_unit is typically used to define units of resources for a particular network device.
//
// In the context of a switch, suppose you want to manage the number of routing table entries. The
// devlink_resource_unit might define the unit of these entries, so the user can query and configure
// the device to allocate or deallocate entries in a meaningful way.
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/devlink.h
type DevlinkResourceUnitEnum uint8

const (
	// DevlinkResourceUnitEntry represents the DEVLINK_RESOURCE_UNIT_ENTRY.
	DevlinkResourceUnitEntry DevlinkResourceUnitEnum = 0
)
