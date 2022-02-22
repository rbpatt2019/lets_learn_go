package main

import (
	"errors"
	"fmt"
	"os"
)

// Error wrapping passes the error back to your code
// with additional context about where/what went wrong.
// This is achieved with the %w verb in Errorf.
func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in fileChecker: %w", err)
	}
	defer f.Close()
	return nil
}

// Example with unwrap.
// If you wanted custom, implement an unwrap method.
func main() {
	err := fileChecker("whoops")
	if err != nil {
		fmt.Println(err)
		if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
			fmt.Println(wrappedErr)
		}
	}
}
