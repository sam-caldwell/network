package core

// DeserializeCnMsgOp - Deserialize []byte into CnMsgOp
func DeserializeCnMsgOp(b []byte) (*CnMsgOp, error) {
	var o CnMsgOp
	if err := o.Deserialize(b); err != nil {
		return nil, err
	}
	return &o, nil
}
