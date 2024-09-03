package core

// CnMsg - cn_msg - represents a connector message.
//
// Structure in the Linux kernel's connector subsystem, used for communication between user space and kernel space.
// It encapsulates a message with a header (nlmsghdr) and data (cn_msg), allowing for various types of notifications
// or commands to be sent between the kernel and user space processes.
//
// See https://github.com/torvalds/linux/blob/master/drivers/connector/cn_proc.c
type CnMsg struct {

	// ID - Contains Idx and Val fields to identify the message or its source/destination.
	//
	// See https://github.com/torvalds/linux/blob/master/drivers/connector/cn_proc.c
	ID CbID

	// Seq - Sequence number to track the order of messages.
	//
	// See https://github.com/torvalds/linux/blob/master/drivers/connector/cn_proc.c
	Seq uint32

	// Ack - Acknowledgment number to confirm the receipt or manage synchronization.
	//
	// See https://github.com/torvalds/linux/blob/master/drivers/connector/cn_proc.c
	Ack uint32

	// Length - Size of the payload that follows the CnMsg header.
	//
	// See https://github.com/torvalds/linux/blob/master/drivers/connector/cn_proc.c
	Length uint16

	// Flags - message properties, such as indicating a multipart message.
	//
	// See https://github.com/torvalds/linux/blob/master/drivers/connector/cn_proc.c
	Flags uint16
}
