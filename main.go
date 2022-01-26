package main

import (
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// User-defined types can have methods.
//
// These are declared very much like functions,
// with the addition of a receiver specification.
// Idiomatically, this is a one-ish letter abbreviation of the type.
//
// The receiver should be a pointer if:
// 1) The method modifies the receiver.
// 2) The method must handle nil instances.
// If either of the above are true, then normal practice is to have all methods
// for that type be pointer receivers.
//
// Also, avoid getter and setter methods.
//
// And methods should follow their type.
func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

func main() {
	p := Person{
		FirstName: "Ryan",
		LastName:  "Patterson-Cross",
		Age:       101,
	}
	fmt.Println(p.String())
}
