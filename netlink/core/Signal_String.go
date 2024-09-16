package core

// String - return the string representation of the signal value
func (sig *Signal) String() string {
	switch *sig {
	case SIGABRT:
		return "SIGABRT" //Signal(unix.SIGABRT) // 0x6
	case SIGALRM:
		return "SIGALRM" // Signal(unix.SIGALRM) // 0xe
	case SIGFPE:
		return "SIGFPE" // Signal(unix.SIGFPE)  // 0x8
	case SIGHUP:
		return "SIGHUP" // Signal(unix.SIGHUP)  // 0x1
	case SIGILL:
		return "SIGILL" // Signal(unix.SIGILL)  // 0x4
	case SIGINT:
		return "SIGINT" // Signal(unix.SIGINT)  // 0x2
	case SIGIOT:
		return "SIGIOT" // Signal(unix.SIGIOT)  // 0x6
	case SIGKILL:
		return "SIGKILL" // Signal(unix.SIGKILL) // 0x9
	case SIGPIPE:
		return "SIGPIPE" // Signal(unix.SIGPIPE) // 0xd
	case SIGQUIT:
		return "SIGQUIT" // Signal(unix.SIGQUIT) // 0x3
	case SIGSEGV:
		return "SIGSEGV" // Signal(unix.SIGSEGV) // 0xb
	case SIGTERM:
		return "SIGTERM" // Signal(unix.SIGTERM) // 0xf
	case SIGTRAP:
		return "SIGTRAP" // Signal(unix.SIGTRAP) // 0x5
	default:
		return "undefined"
	}
}
