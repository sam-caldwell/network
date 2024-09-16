package core

import (
	"encoding/json"
)

// MarshalJSON provides custom marshaling to JSON.
func (req *NetlinkRequest) MarshalJSON() ([]byte, error) {
	// Marshal the NetlinkRequest using default JSON marshaling
	type Alias NetlinkRequest
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(req),
	})
}
