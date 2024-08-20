package test

import (
	"log"
	"os/exec"
)

// RunContainer - Run a given test program inside a docker container
func RunContainer(imageName, test string) error {
	const command = "docker"
	args := []string{
		"run", "--cap-add=NET_ADMIN", "-v", GetCurrentWorkingDirectory() + ":/opt", imageName,
	}
	output, err := exec.Command(command, args...).Output()
	if err != nil {
		return err
	}
	log.Println(string(output))
	return nil
}
