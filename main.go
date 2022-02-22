package main

import (
	"errors"
	"fmt"
	"os"
)

func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in fileChecker: %w", err)
	}
	defer f.Close()
	return nil
}

// Use `errors.Is` to check if any error in the error chain matches a sentinel error.
func main() {
	err := fileChecker("whoops")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("File does not exist")
		}
	}
}
