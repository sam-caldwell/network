package test

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// GoToRootDirectory navigates upwards until it finds a directory with the specified name.
func GoToRootDirectory(name string, maxLevels int) (err error) {
	currentDir := GetCurrentWorkingDirectory()
	log.Printf("Starting Directory: %s", currentDir)

	for level := 0; level < maxLevels; level++ {
		if _, err := os.Stat(filepath.Join(currentDir, name)); err == nil {
			// The root directory with the specified name is found.

			return os.Chdir(filepath.Join(currentDir, name))
		}

		// Move up one directory.
		parentDir := filepath.Dir(currentDir)
		if parentDir == currentDir {
			return fmt.Errorf("directory '%s' not found", name)
		}

		currentDir = parentDir
		err = os.Chdir(currentDir)
		if err != nil {
			return fmt.Errorf("failed to change directory to %s: %v", currentDir, err)
		}
	}
	return fmt.Errorf("directory '%s' not found", name)
}
