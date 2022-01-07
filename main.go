package main

import (
	"fmt"
	"sort"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// Closures - functions as parameters
//
// These can reference local variables and then be passed elsewhere.
func main() {
	people := []Person{
		{"Pat", "Patterson", 25},
		{"Tracy", "Bobbert", 23},
		{"Fred", "Fredson", 18},
	}
	fmt.Println("Before sorting", people)

	// Can sort any slice by the passed functions
	// Which can be anonymous (think lambdas)
	// Note how `people` is captured by the closure.
	sort.Slice(people, func(i, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println("After sorting", people)
}
