package main

import (
	"os"
)

// Generic map functions to run a function over one array and return another.
// Not supported by default in go.
func Map[T any, R any](s []T, f func(T) R) []R {
	result := make([]R, len(s))
	for i, v := range s {
		result[i] = f(v)
	}
	return result
}

type FileExistsDelegate func(absolutePath string) bool

func FileExists(absolutePath string) bool {
	_, err := os.Stat(absolutePath)
	return err == nil
}
