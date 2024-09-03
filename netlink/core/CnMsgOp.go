package core

// CnMsgOp - represents connector message operation, combining message identification, sequence, acknowledgment,
// and operational flags with an additional operation-specific field.
//
// Note that this EXTENDS cn_msg found in the Linux source tree and adds an Op field.
//
// https://github.com/torvalds/linux/blob/master/include/linux/connector.h
type CnMsgOp struct {
	CnMsg
	// here we differ from the C header
	Op uint32
}
