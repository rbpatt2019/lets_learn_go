package main

import (
	"fmt"
)

type Adder struct {
	start int
}

func (a Adder) AddTo(val int) int {
	return a.start + val
}

func main() {
	// Method value
	// Assigning a method to a variable
	// In this case, the method is bound to a specific object.
	myAdder := Adder{start: 10}
	f1 := myAdder.AddTo
	fmt.Println(f1(10))

	// Method Expression
	// Assigning a method from a literal to a variable
	// In this case, it is not bound to an instance,
	// and must be passed an instance of the appropriate type as the first parameter.
	f2 := Adder.AddTo
	fmt.Println(f2(myAdder, 10))
}
