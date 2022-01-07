package main

import "fmt"

type person struct {
	name string
}

func modifyNot(i int, s string, p person) {
	i *= 2
	s = "Nope"
	p.name = "Bob"
}

func modifyMostly(m map[int]string, l []int) {
	m[2] = "yes"
	m[3] = "totally"
	delete(m, 1)

	for k, v := range l {
		l[k] = v * 2
	}
	l = append(l, 10)
}

// Go is call by value
//
// When you supply a variable for a parameter,
// Go always makes a copy.
// This effectively makes variables immutable,
// unless pointers are used (more on that later).
//
// This also means that maps and slices behave slightly differently,
// as they are implemented with pointers.
func main() {
	p := person{}
	i := 2
	s := "Yes"
	modifyNot(i, s, p)
	fmt.Println(i, s, p) // None will change.

	m := map[int]string{
		1: "a",
		2: "b",
	}
	l := []int{1, 2, 3}
	modifyMostly(m, l)
	fmt.Println(m, l)
	// Maps are fully mutable (pointer)
	// Slices cannot be lengthened.
}
