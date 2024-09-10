package core

// enableErrorMessageReporting is the default error message reporting configuration for the new netlink sockets
var enableErrorMessageReporting bool = false

// EnableErrorMessageReporting - enable error message reporting for new netlink sockets
func EnableErrorMessageReporting() {
	enableErrorMessageReporting = true
}

// DisableErrorMessageReporting - disable error message reporting for new netlink sockets
func DisableErrorMessageReporting() {
	enableErrorMessageReporting = false
}
