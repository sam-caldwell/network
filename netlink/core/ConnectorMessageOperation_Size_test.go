package core

import (
	"testing"
)

func TestConnectorMessageOperation_Size(t *testing.T) {
	const expectedSizeInBytes = 24
	if ConnectorMessageOperationSize != expectedSizeInBytes {
		t.Fatal("Wrong size")
	}
}
