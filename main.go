package main

import (
	"errors"
	"fmt"
	"os"
)

// Basic error handling.
// Create by calling `errors.New`.
// Should not be capitalised, and should not end with punctuation.
// Return nil if successful.
func calcRemainderAndMod(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("denominator is 0")
	}
	return numerator / denominator, numerator % denominator, nil
}

// Check for errors by comparing err to nil.
func main() {
	r, m, err := calcRemainderAndMod(1, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(r, m)
}
