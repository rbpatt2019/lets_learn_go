package main

import "fmt"

func main() {
	// Arrays can be declared by:
	// 1) using zero values of the specified type.
	var x [3]int
	fmt.Printf("%T, %v\n", x, x)
	// 2) specifying the values for an array literal.
	y := [3]int{1, 2, 3}
	fmt.Printf("%T, %v\n", y, y)
	// 3) defining a sparse array.
	z := [10]int{1, 5: 2, 9: 3}
	fmt.Printf("%T, %v\n", z, z)

	// However, arrays are very rare for direct usage in Go
	// As the length is part of the type, there are some odd limitations
	// the size cannot be defined with a variable (types must be known at runtime)
	// arrays of different size cannot be converted to identical types
	// They exist mostly to allow slices.

	// Slices
	// These are specified the same, but with without size
	// Their 0 value is `nil`.
	var a []int
	fmt.Printf("%T, %v\n", a, a)
	// The only allowed comparisons for slices is to check for nil.
	fmt.Println(a == nil)
	// Slices can be grown
	// To append a slice, we use the variadic operator.
	a = append(a, 50)
	b := []int{1, 2, 3}
	a = append(a, b...)
	fmt.Printf("%T, %v\n", a, a)
}
