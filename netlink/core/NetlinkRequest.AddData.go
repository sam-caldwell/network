package core

// AddData - append the NetlinkRequest data with the given NetlinkRequestData
func (req *NetlinkRequest) AddData(data NetlinkRequestData) {

	req.Data = append(req.Data, data)

}
