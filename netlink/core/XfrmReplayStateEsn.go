package core

// XfrmReplayStateEsn represents the `xfrm_replay_state_esn` structure in the Linux kernel's IPsec subsystem.
// This structure is used to manage the extended sequence number (ESN) replay protection mechanism in IPsec,
// which is designed to prevent replay attacks by keeping track of packet sequence numbers and windows.
//
// ESN extends the standard 32-bit sequence number to 64 bits to support larger traffic flows without replay issues.
//
// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
type XfrmReplayStateEsn struct {
	// BmpLen represents the length of the bitmap (in 32-bit words) used for replay protection.
	// The bitmap helps track packets that are within the replay window.
	BmpLen uint32

	// OSeq represents the sequence number of the outbound packet.
	// This field is used to manage the sequence numbers of the outbound IPsec traffic.
	OSeq uint32

	// Seq represents the current sequence number of the inbound packet.
	// This field helps in keeping track of the next expected packet in the replay window for inbound traffic.
	Seq uint32

	// OSeqHi represents the high-order 32 bits of the outbound sequence number.
	// Together with OSeq, it forms a 64-bit sequence number for outbound IPsec traffic.
	OSeqHi uint32

	// SeqHi represents the high-order 32 bits of the inbound sequence number.
	// Together with Seq, it forms a 64-bit sequence number for inbound IPsec traffic.
	SeqHi uint32

	// ReplayWindow represents the size of the replay window (in terms of packets) used for replay protection.
	// This field defines how many packets are tracked for replay checking.
	ReplayWindow uint32

	// Bmp is the bitmap used for replay protection.
	// It is an array of 32-bit words that helps track received packets within the replay window.
	// Each bit in the bitmap corresponds to a packet, with 1 representing a received packet and 0 representing a missed packet.
	Bmp []uint32
}
