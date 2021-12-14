package main

import (
	"fmt"
)

// Primitive types and Declarations.
func main() {
	// Variables can be declared 2 ways:
	// With the `var` keyword, or with :=
	// Within functions, := predominates
	// Exceptions include: 1) initialising to the zero val (requires `var`).
	var x int
	fmt.Printf("%d", x)
	// 2) Assigning an untyped literal to a var where the default type doesn't match the desired.
	var y byte = 20
	fmt.Printf("%T", y)
	// 3) To avoid shadowing existing variables (more on that later).

	// Numerics
	// Can be ints or floats and signed or unsigned
	// Zero value is, well, 0
	// Some special notes:
	// Byte (as above) is a synonym for uint8 - but byte is idiomatic
	// When choosing which form of int to use:
	// 1) if you are working on a system with a specific integer size, use that
	// 2) if you're library that should work with both, use int64 and uint64
	// 3) otherwise, use int
	// Ints support the usual ops.
	x = 20
	x *= 4
	fmt.Printf("%d\n", x)

	// Strings/Runes
	// Runes represent literal single characters and are marked with single quotes
	// Under the hood, these are int32's.
	z := 'a'
	fmt.Printf("%T, %c\n", z, z)
	// Strings are represented with double quotes
	// and contain 0 or more runes.
	a := "Hello world"
	fmt.Printf("%T, %s\n", a, a)

	// Booleans
	// Have a zero value of False
	// "Truthiness/Falsiness" is not allowed.
	var b bool
	fmt.Printf("%T, %t", b, b)
}
