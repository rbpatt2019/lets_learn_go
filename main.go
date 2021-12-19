package main

import "fmt"

// Shadowing.
// A shadowing variable has the same name as a variable in a containing block.
func main() {
	x := 10
	if x > 5 {
		// Uses x from the main block.
		// As x has not been declared in this block.
		// Prints 10.
		fmt.Println(x)
		// Creates a new variable constrained to the `if` block.
		x := 5
		// Prints 5.
		fmt.Println(x)
	}
	// Uses x from the main block.
	// Prints 10.
	fmt.Println(x)
}
