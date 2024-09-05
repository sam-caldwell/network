package core

// Len - Returns the size of the TcConnmark structure in bytes.
// This method calculates the total size of the TcConnmark object,
// which includes the embedded TcGen structure's size and the size
// of the Zone field.
//
// The `SizeofTcConnmark` is assumed to be a constant defined elsewhere,
// representing the total byte size of the structure.
//
// This method is particularly useful for serialization and deserialization
// of TcConnmark objects in netlink messages or other low-level networking operations.
func (msg *TcConnmark) Len() int {
	return SizeofTcConnmark
}
