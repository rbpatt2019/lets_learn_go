package main

import "fmt"

// Slice expressions.
func main() {
	// Slicing is 0-indexed, open-ended.
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x[:2], x[2:])
	// Assigning a slice expression to a var means the memory is shared.
	y := x[:2]
	y[0] = 100
	fmt.Println(x, y)
	// And things get really wonky with append,
	// Because capacity is shared!
	// Y has capacity 4, and shares memory with x.
	// So when y is appended to, its third element goes into the same memory as the third element of x!
	y = append(y, 200)
	fmt.Println(x, y)
	// To avoid, either never append to a substring,
	// or use full slice expression notations.
	// This contains a third position to indicate how much capacity is available to rhe sub-slice.
	z := x[:2:2]
	fmt.Println(z)
	z = append(z, 500)
	fmt.Println(x, z)
	// And arrays can be converted to slices,
	// but run ito the same memory sharing issues...
	a := [4]int{1, 2, 3, 4}
	b := a[:]
	fmt.Printf("a: %T, %v\nb: %T, %v\n", a, a, b, b)
	// For memory independence, use copy which, well copies the slice.
	c := make([]int, 2)
	copy(c, b)
	fmt.Println(c)
}
