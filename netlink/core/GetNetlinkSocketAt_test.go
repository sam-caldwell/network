package core

import (
	"github.com/sam-caldwell/network/test"
	"testing"
)

// Test for GetNetlinkSocketAt function
func TestGetNetlinkSocketAt(t *testing.T) {
	t.Skip("disabled")
	t.Run("Test successful namespace switch", func(t *testing.T) {
		const testDockerImage = "network-test:latest"
		t.Run("build container", func(t *testing.T) {
			if err := test.BuildTestContainer(testDockerImage); err != nil {
				t.Fatalf("container build error: %v", err)
			}
		})
		t.Run("run test container", func(t *testing.T) {
			if err := test.RunContainer(testDockerImage, "TestGetNetlinkSocketAtHappyPath"); err != nil {
				t.Fatalf("test container failed: %v", err)
			}
		})
	})

	t.Run("Test error on setting invalid namespace", func(t *testing.T) {
		const testDockerImage = "network-test:latest"
		t.Run("build container", func(t *testing.T) {
			if err := test.BuildTestContainer(testDockerImage); err != nil {
				t.Fatalf("container build error: %v", err)
			}
		})
		t.Run("run test container", func(t *testing.T) {
			if err := test.RunContainer(testDockerImage, "TestGetNetlinkSocketAtSadPath"); err != nil {
				t.Fatalf("test container failed: %v", err)
			}
		})
	})
}
