package core

// TcaNetEmEnum - Enumeration for NetEm (Network Emulation) attributes.
// These attributes are used in the Linux kernel's Traffic Control system (NetEm) to emulate various network behaviors
// such as delay, packet loss, reordering, and corruption.
//
// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
type TcaNetEmEnum uint8

const (
	// TCANetEmUnspec - TCA_NETEM_UNSPEC - Unspecified NetEm attribute.
	// Placeholder for an unspecified or unknown attribute.
	TCANetEmUnspec TcaNetEmEnum = iota

	// TCANetEmCorr - TCA_NETEM_CORR - Correlation of packet delay and loss.
	// Sets correlation between packet delay or loss events.
	TCANetEmCorr TcaNetEmEnum = iota

	// TCANetEmDelayDist - TCA_NETEM_DELAY_DIST - Distribution of packet delay.
	// Specifies a distribution model for packet delay to create more complex emulation of network delays.
	TCANetEmDelayDist TcaNetEmEnum = iota

	// TCANetEmReorder - TCA_NETEM_REORDER - Packet reordering.
	// Controls packet reordering to emulate out-of-order packet delivery.
	TCANetEmReorder TcaNetEmEnum = iota

	// TCANetEmCorrupt - TCA_NETEM_CORRUPT - Packet corruption.
	// Introduces random corruption in packets to simulate damaged network packets.
	TCANetEmCorrupt TcaNetEmEnum = iota

	// TCANetEmLoss - TCA_NETEM_LOSS - Packet loss.
	// Simulates random packet loss to model unreliable networks.
	TCANetEmLoss TcaNetEmEnum = iota

	// TCANetEmRate - TCA_NETEM_RATE - Rate limitation.
	// Limits the bandwidth of the emulated network.
	TCANetEmRate TcaNetEmEnum = iota

	// TCANetEmEcn - TCA_NETEM_ECN - Explicit Congestion Notification.
	// Enables Explicit Congestion Notification (ECN) marking in the emulated network.
	TCANetEmEcn TcaNetEmEnum = iota

	// TCANetEmRate64 - TCA_NETEM_RATE64 - 64-bit rate limitation.
	// Extends the rate limitation attribute to use 64-bit values for higher precision.
	TCANetEmRate64 TcaNetEmEnum = iota

	// TCANetEmPad - TCA_NETEM_PAD - Padding for alignment.
	// Adds padding for alignment purposes.
	TCANetEmPad TcaNetEmEnum = iota

	// TCANetEmLatency64 - TCA_NETEM_LATENCY64 - 64-bit latency value.
	// Allows for higher precision in specifying network latency.
	TCANetEmLatency64 TcaNetEmEnum = iota

	// TCANetEmJitter64 - TCA_NETEM_JITTER64 - 64-bit jitter value.
	// Introduces jitter in network latency with higher precision.
	TCANetEmJitter64 TcaNetEmEnum = iota

	// TCANetEmSlot - TCA_NETEM_SLOT - Slot mechanism for packet transmission timing.
	// Controls the transmission slot mechanism to simulate time-based network behaviors.
	TCANetEmSlot TcaNetEmEnum = iota

	// TCANetEmSlotDist - TCA_NETEM_SLOT_DIST - Slot distribution.
	// Allows the distribution of transmission slots.
	TCANetEmSlotDist TcaNetEmEnum = iota

	// TCANetEmPrngSeed - TCA_NETEM_PRNG_SEED - Pseudo-random number generator seed.
	// Sets the seed for the pseudo-random number generator used in network emulation.
	TCANetEmPrngSeed TcaNetEmEnum = iota

	// TCANetEmMax - TCA_NETEM_MAX - Maximum NetEm attribute.
	// Represents the maximum valid value for the NetEm attributes.
	TCANetEmMax = iota - 1
)
