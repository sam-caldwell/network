package core

import "fmt"

// String returns a string representation of NetlinkRequest in JSON format.
func (req *NetlinkRequest) String() string {
	data, err := req.MarshalJSON()
	if err != nil {
		// Instead of panicking, handle the error by returning an error message
		return fmt.Sprintf("error marshaling NetlinkRequest to JSON: %v", err)
	}
	return string(data)
}
