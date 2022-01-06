package main

import (
	"errors"
	"fmt"
	"os"
)

// Multiple return values.
//
// Unlike, say, python, these are actually multiple values,
// not a single tuple that can be unpacked.
//
// Also, return values can be named.
// When named, they are initilaised to their 0 value.
// There is some debate as to the utility of this,
// due to shadowing risks, and some potential odd return behaviour.
// Your choice.
func divAndRemainder(numerator, denominator int) (result, remainder int, err error) {
	if denominator == 0 {
		err = errors.New("Cannot divide by 0")
		return result, remainder, err
	}
	result, remainder = numerator/denominator, numerator%denominator
	return result, remainder, err
}

// Functions.
//
// Also, we see here a simple example of the idiomatic way to handle errors.
// They are the last value returned from a a function,
// and you retunr nil if there is no error.
func main() {
	result, remainder, err := divAndRemainder(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder)
}
