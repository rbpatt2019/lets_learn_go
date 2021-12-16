package main

import "fmt"

// When you have related data that you want to group together,
// Use a struct!
// These are NOT objects/classes as they lack inheritance.
type person struct {
	name string
	age  int
}

// Structs.
func main() {
	// The zero value of a struct has all fields set to that field's 0-value.
	var fred person
	fmt.Println(fred)
	// Structs can be created as literals.
	// Fields not specified get their 0 value.
	beth := person{
		age:  30,
		name: "Beth",
	}
	fmt.Println(beth)
	// Anonymous structs - creating a struct type without a name -
	// are a thing, but they are mostly used for data marshalling
	// and testing.
	pet := struct {
		name string
		kind string
	}{
		name: "fido",
		kind: "cat",
	}
	fmt.Printf("pet: %T, %v\n", pet, pet)
	// Structs are ony comparable if al their fields are comparable
	// And type conversions are only possible if all fields are in the same order, name, and type.
	fmt.Println(beth == fred)
}
