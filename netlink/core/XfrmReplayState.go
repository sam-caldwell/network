package core

// XfrmReplayState represents the `xfrm_replay_state` structure in the Linux kernel's IPsec subsystem.
// This structure is used to manage replay protection in IPsec, which prevents replay attacks by tracking packet
// sequence numbers and maintaining a bitmap for received packets.
//
// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
type XfrmReplayState struct {
	// OSeq represents the sequence number of the outbound packet.
	// This field is used to manage the sequence numbers of outbound IPsec traffic.
	// It helps in ensuring that packets are sent in the correct sequence to avoid replay attacks.
	OSeq uint32

	// Seq represents the current sequence number of the inbound packet.
	// This field helps keep track of the next expected packet for inbound traffic and ensures that
	// packets are processed in the correct order.
	Seq uint32

	// BitMap is a bitmap used to track which packets have been received within the replay window.
	// The bitmap represents the state of received packets: 1 indicates a packet has been received,
	// and 0 indicates that it has not. This helps detect and prevent replayed packets.
	BitMap uint32
}
