package test

import (
	"log"
	"os/exec"
	"strings"
)

// DeleteContainer - delete the named container
func DeleteContainer(testName string) error {
	const command = "docker"
	args := []string{
		"rm",
		"-f",
		testName,
	}
	log.Println(command, strings.Join(args, " "))
	output, err := exec.Command(command, args...).CombinedOutput()
	log.Println("output:")
	log.Println(string(output))
	log.Println("-----------------")
	return err
}
