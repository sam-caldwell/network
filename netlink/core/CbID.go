package core

// CbID - cb_id -  represents an identifier used in callback or tracking mechanisms.
//
// idx and val are unique identifiers which are used for message routing and must be registered in connector.h for
// in-kernel usage.
//
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/connector.h#L65
type CbID struct {
	Idx uint32
	Val uint32
}
