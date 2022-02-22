package main

import (
	"errors"
)

// Use `errors.As` to check if any error in the error chain matches a specific type.
// It takes the error and a pointer to a variable of the desired error type.
// Then, if the error is present, it is assigned to the pointer.
func main() {
	err := SomethingThatErrors()
    var myErr MyErr
    if errors.As(err, &myErr) {
        // Do things with myErr
    }
}
