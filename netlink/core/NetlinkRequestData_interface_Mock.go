package core

import (
	"errors"
)

// MockNetlinkRequestData - Mock struct implementing NetlinkRequestData for testing purposes.
type MockNetlinkRequestData struct {
	data []byte
}

// Len - Returns the length of the data.
func (m *MockNetlinkRequestData) Len() int {
	return len(m.data)
}

// Serialize - Serializes the data into a byte slice.
func (m *MockNetlinkRequestData) Serialize() ([]byte, error) {
	if len(m.data) == 0 {
		return nil, errors.New("no data to serialize")
	}
	return m.data, nil
}
