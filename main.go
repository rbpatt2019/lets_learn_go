package main

import "fmt"

// Variadic Inputs
//
// Must be last (or only) parameters.
// Can be multiple values.
func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

// Functions.
func main() {
	fmt.Println(addTo(3))
	fmt.Println(addTo(3, 2, 3, 6))
	// If you pass a slice as a variadic argument, you must add the trailing dots.
	fmt.Println(addTo(3, []int{2, 3, 6}...))
}
