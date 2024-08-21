package namespace

import (
	"github.com/sam-caldwell/network/test"
	"testing"
)

func TestGetFromName(t *testing.T) {
	t.Skip("Disabled for debugging")
	const testDockerImage = "network-test:latest"

	t.Run("build container", func(t *testing.T) {
		if err := test.BuildTestContainer(testDockerImage); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("run test container", func(t *testing.T) {
		if err := test.RunContainer(testDockerImage, "TestCreateGetFromName"); err != nil {
			t.Fatal(err)
		}
	})
}
