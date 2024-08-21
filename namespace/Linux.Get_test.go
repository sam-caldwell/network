package namespace

import (
	"github.com/sam-caldwell/network/test"
	"testing"
)

func TestGet(t *testing.T) {
	const testDockerImage = "network-test:latest"

	t.Run("build container", func(t *testing.T) {
		if err := test.BuildTestContainer(testDockerImage); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("run test container", func(t *testing.T) {
		if err := test.RunContainer(testDockerImage, "TestCreateGetSetNamespace"); err != nil {
			t.Fatal(err)
		}
	})
}
