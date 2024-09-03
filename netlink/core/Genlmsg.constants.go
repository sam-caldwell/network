package core

import (
	"golang.org/x/sys/unix"
)

const (
	// GenlAdminPerm - GENL_ADMIN_PERM - This constant represents the permission required for administrative commands
	//                 in the Generic Netlink (Genl) framework, ensuring that only privileged users can execute certain
	//                 commands.
	GenlAdminPerm = unix.GENL_ADMIN_PERM

	// GenlCmdCapDo - GENL_CMD_CAP_DO - This constant indicates that a Generic Netlink command has the capability to
	//                perform an action (typically an administrative or control operation).
	GenlCmdCapDo = unix.GENL_CMD_CAP_DO

	// GenlCmdCapDump - GENL_CMD_CAP_DUMP - This constant signifies that a Generic Netlink command has the capability
	//                  to dump data, usually for retrieving information in bulk.
	GenlCmdCapDump = unix.GENL_CMD_CAP_DUMP

	// GenlCmdCapHaspol - GENL_CMD_CAP_HASPOL - This constant specifies that a Generic Netlink command has an
	//                    associated policy, which defines the expected attributes and constraints for that command.
	GenlCmdCapHaspol = unix.GENL_CMD_CAP_HASPOL

	// GenlHdrlen - GENL_HDRLEN - This constant defines the length of the Generic Netlink header, which is added to
	//              all Generic Netlink messages.
	GenlHdrlen = unix.GENL_HDRLEN

	// GenlIdCtrl - GENL_ID_CTRL - This constant represents the identifier for the Generic Netlink control family,
	//              which is used to manage and query other Generic Netlink families.
	GenlIdCtrl = unix.GENL_ID_CTRL
)
