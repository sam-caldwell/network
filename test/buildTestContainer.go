package test

import (
	"log"
	"os/exec"
)

// BuildTestContainer - Build the docker container used to test our solution.
func BuildTestContainer(imageName string) error {
	if err := GoToRootDirectory("network", 4); err != nil {
		return err
	}
	output, err := exec.Command("docker", "build", "-f", "test/Dockerfile", "--tag", imageName, ".").
		CombinedOutput()

	log.Printf("Container build output:\n"+
		"===\n"+
		"%s\n"+
		"===\n",
		output)

	return err
}
