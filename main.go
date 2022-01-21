package main

import (
	"fmt"
	"log"
)

// Pointers
//
// Pointers increase the workload of the GC,
// and make data flow harder to understand,
// so don't use unless necessary.
// E.g.
// Create a struct by having a function instantiate one
// rather than passing a pointer.
type Foo struct {
	Field1 string
	Field2 int
}

func MakeFoo(field1 string, field2 int) (Foo, error) {
	f := Foo{
		Field1: field1,
		Field2: field2,
	}
	return f, nil
}

func main() {
	f, err := MakeFoo("Hello", 10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(f)
}
