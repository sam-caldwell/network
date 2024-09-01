package core

// CnMsgOp - represents connector message operation, combining message identification, sequence, acknowledgment,
// and operational flags with an additional operation-specific field.
type CnMsgOp struct {
	CnMsg
	// here we differ from the C header
	Op uint32
}
