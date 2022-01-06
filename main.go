package main

import "fmt"

// A function declaration!
// All parameters must bet typed,
// And consecutive parameters of the same type may be typed after the last one.
// NB: Go does not have named or optional parameters!
func div(numerator, denominator int) int {
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}

// Functions.
func main() {
	results := div(5, 2)
	fmt.Println(results)
}
