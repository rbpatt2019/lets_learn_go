package main

import "fmt"

// Capacity vs length.
func main() {
	// First note:
	// Empty slices from var are declared but not assigned. Hence, nil.
	// Empty slices from make are declared and assigned.
	// They are initialised with the zero value of their type.
	var a []int
	b := make([]int, 5, 10)
	fmt.Println("a ", a, len(a), cap(a))
	fmt.Println("b ", b, len(b), cap(b))
	// Length is the nuber of elements in the array.
	// Capacity is the number of elemets that could be in the array.
	// If an append would make len > cap,
	// a new slice is created with 2x cap (if cap < 1_024)
	// or 1.25*cap (if cap > 1024).
	c := []int{1, 2, 3, 4, 5}
	b = append(b, c...)
	fmt.Println("b ", b, len(b), cap(b))
	b = append(b, 10)
	fmt.Println("b ", b, len(b), cap(b))

	// Given all this, what's the best way to declare a slice?
	// 1) If it might stay nil (ie possibly no return value in a function), use:
	// var x []int
	// 2) If there are starting values, or the values won't change, use:
	//  y := []int{1, 2, 3}
	// 3) If you know the size, but not the values, use:
	//  z := make([]int, 1, 5)
	// Use a nonzero length when using a slice as a buffer
	// If you 100% know the length, then use a non-zero length
	// Oterwise, use a zero length and a non-zero capacity.
}
