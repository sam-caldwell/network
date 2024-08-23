package core

// CnMsg represents a connector message.
type CnMsg struct {
	ID     CbID   // Contains Idx and Val fields to identify the message or its source/destination.
	Seq    uint32 // Sequence number to track the order of messages.
	Ack    uint32 // Acknowledgment number to confirm the receipt or manage synchronization.
	Length uint16 // Length of the payload that follows the CnMsg header.
	Flags  uint16 // Flags for message properties, such as indicating a multi-part message.
}
