package core

// CnMsg - cn_msg - represents a connector message.
//
// Structure in the Linux kernel's connector subsystem, used for communication between user space and kernel space.
// It encapsulates a message with a header (nlmsghdr) and data (cn_msg), allowing for various types of notifications
// or commands to be sent between the kernel and user space processes.
//
// See https://github.com/torvalds/linux/blob/master/drivers/connector/cn_proc.c
type CnMsg struct {
	ID     CbID   // Contains Idx and Val fields to identify the message or its source/destination.
	Seq    uint32 // Sequence number to track the order of messages.
	Ack    uint32 // Acknowledgment number to confirm the receipt or manage synchronization.
	Length uint16 // Length of the payload that follows the CnMsg header.
	Flags  uint16 // Flags for message properties, such as indicating a multipart message.
}
