package core

// Genlmsg represents a generic netlink message with a command and version.
type Genlmsg struct {
	// Command identifier
	Command uint8
	// Version of the interface
	Version uint8
}
