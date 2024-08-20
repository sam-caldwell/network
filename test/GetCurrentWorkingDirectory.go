package test

import (
	"os"
	"path/filepath"
)

// GetCurrentWorkingDirectory - return the current working directory
func GetCurrentWorkingDirectory() string {
	dir, _ := os.Getwd()
	absPath, _ := filepath.Abs(dir)
	return absPath
}
