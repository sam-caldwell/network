package core

// Execute - Execute the request against the given sockType.
//
// Returns a list of netlink messages in serialized format, optionally filtered by resType.
//
// Depends on ExecuteIter()
func (req *NetlinkRequest) Execute(sockType int, resType uint16) ([][]byte, error) {

	var res [][]byte

	err := req.ExecuteIter(sockType, resType, func(msg []byte) bool {
		res = append(res, msg)
		return true
	})

	return res, err

}
