package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}

// Pointers
//
// A pointer is simply a variable whose contents are the address
// where another variable is stored.
//
// Its zero value is nil.
func main() {
	x := "Hello"
	// & os the address operator.
	// It returns a pointer to the variable.
	pointerToX := &x
	// * is the indirection operator.
	// It returns the value a pointer points to.
	// This is called dereferencing.
	// Dereferencing a nil pointer panics.
	fmt.Println(pointerToX)
	fmt.Println(*pointerToX)
	// New returns a pointer to the zero value of the indicated type.
	y := new(int)
	fmt.Println(y, *y)
	// Thought that's not particularly idiomatic.
	// For primitive types, declare a variable then point to it.
	var z string
	a := &z
	fmt.Println(a)
	// For structs, derefence a struct literal to create a pointer.
	b := &Person{}
	fmt.Println(b)
}
