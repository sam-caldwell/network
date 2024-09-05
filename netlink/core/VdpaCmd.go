package core

// VdpaCmd represents the different command types for VDPA interface.
// For reference, see: https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/tree/include/uapi/linux/vdpa.h
type VdpaCmd uint8

const (
	// vdpaCmdUnspec - VDPA_CMD_UNSPEC - Unspecified command.
	vdpaCmdUnspec VdpaCmd = iota

	// vdpaCmdMgmtDevNew - VDPA_CMD_MGMTDEV_NEW - Create a new management device.
	vdpaCmdMgmtDevNew

	// vdpaCmdMgmtDevGet - VDPA_CMD_MGMTDEV_GET - Get a management device (can dump).
	vdpaCmdMgmtDevGet

	// vdpaCmdDevNew - VDPA_CMD_DEV_NEW - Create a new VDPA device.
	vdpaCmdDevNew

	// vdpaCmdDevDel - VDPA_CMD_DEV_DEL - Delete a VDPA device.
	vdpaCmdDevDel

	// vdpaCmdDevGet - VDPA_CMD_DEV_GET - Get a VDPA device (can dump).
	vdpaCmdDevGet

	// vdpaCmdDevConfigGet - VDPA_CMD_DEV_CONFIG_GET - Get VDPA device configuration (can dump).
	vdpaCmdDevConfigGet

	// vdpaCmdDevVStatsGet - VDPA_CMD_DEV_VSTATS_GET - Get VDPA device statistics.
	vdpaCmdDevVStatsGet
)
