//go:build linux

package core

// Deserialize - Deserialize byte slice into BridgeVlanInfo
func (bridge *BridgeVlanInfo) Deserialize(b []byte) (err error) {
	result, err := DeserializeBridgeVlanInfo(b)
	if err != nil {
		return err
	}
	*bridge = *result
	return nil
}
