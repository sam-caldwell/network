package core

// TcPriority represents logical priority bands that don't rely on a specific packet scheduler.
// These bands map to real traffic classes by the scheduler if no precise classification exists.
// These values generally coincide with older IPv6 classes, but they have no strict meaning here.
//
// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
type TcPriority uint8

const (
	// TcPrioBestEffort - TC_PRIO_BESTEFFORT - General-purpose, default priority traffic.
	TcPrioBestEffort TcPriority = 0

	// TcPrioFiller - TC_PRIO_FILLER - Lowest priority, used for background traffic.
	TcPrioFiller TcPriority = 1

	// TcPrioBulk - TC_PRIO_BULK - Bulk data transfers, such as file downloads.
	TcPrioBulk TcPriority = 2

	_ = 3 // Skipping unused value

	// TcPrioInteractiveBulk - TC_PRIO_INTERACTIVE_BULK - Interactive bulk transfers.
	TcPrioInteractiveBulk TcPriority = 4

	_ = 5 // Skipping unused value

	// TcPrioInteractive - TC_PRIO_INTERACTIVE - High-priority, interactive traffic, such as SSH or telnet.
	TcPrioInteractive TcPriority = 6

	// TcPrioControl - TC_PRIO_CONTROL - Control traffic, such as routing or management traffic.
	TcPrioControl TcPriority = 7
)

const (
	// TcPrioMax - TC_PRIO_MAX - Maximum priority value.
	TcPrioMax TcPriority = 15
)
