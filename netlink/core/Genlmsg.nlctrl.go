package core

// Family nlctrl netlink specification
// See https://docs.kernel.org/networking/netlink_spec/nlctrl.html
const (

	// GenlCtrlName - https://docs.kernel.org/networking/netlink_spec/nlctrl.html
	GenlCtrlName = "nlctrl"

	// GenlCtrlVersion - https://docs.kernel.org/networking/netlink_spec/nlctrl.html
	GenlCtrlVersion = 2

	// GenlCtrlCmdGetFamily - https://docs.kernel.org/networking/netlink_spec/nlctrl.html
	GenlCtrlCmdGetFamily = 3
)
