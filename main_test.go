package main

import (
	"fmt"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMustReturnOnePath(t *testing.T) {
	result := len(which("go", "/usr/bin/:/etc/bin", MockFileExists))
	assert.Equal(t, 1, result, "Should only return 1")

}

func TestMustReturnTwoPaths(t *testing.T) {
	result := len(which("npm", "/usr/bin/:/etc/bin", MockFileExists))
	assert.Equal(t, 2, result, "Should only return 2")

}

func MockFileExists(absolutePath string) bool {
	fileSystem := []string{
		"/etc/bin/npm",
		"/etc/bin/go",
		"/usr/bin/npm",
	}
	fmt.Println(absolutePath)

	fmt.Println(slices.Contains(fileSystem, absolutePath))
	return slices.Contains(fileSystem, absolutePath)
}
