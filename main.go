package main

import "fmt"

// Since the copy of a pointer still points to the same memory,
// this updates the contents.
func update(px *int) {
	*px = 20
}

// Call by value - so this changes the copy, not the original.
// This actually fails lints!
func failUpdate(px *int) {
	x := 20
	px = &x
}

// Pointers
//
// Since go is call by value,
// a passed pointer is copied.
// However, this means that nil pointers cannot be updated by functions.
//
// Additionally,
// if you want the assigned value of a pointer to be there after exiting a function,
// then you must dereference the pointer on assignment.
func main() {
	x := 10
	failUpdate(&x)
	fmt.Println(x)
	update(&x)
	fmt.Println(x)
}
