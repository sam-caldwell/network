package core

// Deserialize - Safely converts a byte slice into a TcSfqQopt struct.
// This method ensures that the length of the input byte slice matches the expected size of TcSfqQopt.
//
// If the length of the input byte slice is less than the size of TcSfqQopt, it returns an error.
func (msg *TcSfqQopt) Deserialize(b []byte) error {

	data, err := DeserializeTcSfqQopt(b)
	if err != nil {
		return err
	}
	*msg = *data
	return nil

}
