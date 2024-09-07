package core

// Serialize - return a byte slice containing the serialized contents of the RtGenMsg struct.
func (msg *RtGenMsg) Serialize() ([]byte, error) {

	out := make([]byte, msg.Len())

	out[0] = msg.Family

	return out, nil

}
