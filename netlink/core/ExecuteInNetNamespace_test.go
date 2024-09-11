package core

import (
	"github.com/sam-caldwell/network/test"
	"testing"
)

func TestExecuteInNetNamespace(t *testing.T) {
	t.Run("Test successful namespace switch", func(t *testing.T) {
		const testDockerImage = "network-test:latest"
		t.Run("build container", func(t *testing.T) {
			if err := test.BuildTestContainer(testDockerImage); err != nil {
				t.Fatalf("container build error: %v", err)
			}
		})
		t.Run("run test container", func(t *testing.T) {
			if err := test.RunContainer(testDockerImage, "TestNetlinkExecuteInNetNamespace"); err != nil {
				t.Fatal(err)
			}
		})
	})

	t.Run("Test error on setting invalid namespace", func(t *testing.T) {

	})
}
