package core

// AddRawData - add raw bytes to the end of the NetlinkRequest object during serialization
func (req *NetlinkRequest) AddRawData(data []byte) {

	req.RawData = append(req.RawData, data...)

}
