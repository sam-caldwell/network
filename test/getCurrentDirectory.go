package test

import (
	"os"
	"path/filepath"
)

// GetCurrentDirectory - return the current working directory
func GetCurrentDirectory() string {
	dir, _ := os.Getwd()
	absPath, _ := filepath.Abs(dir)
	return absPath
}
