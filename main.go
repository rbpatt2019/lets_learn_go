package main

import "fmt"

// Const declares things immutable
// However, they only give names to literals
// This means they can only hold values known at COMPILE time
// A variable calculated at runtime cannot be declared immutable.
const x int64 = 10

func main() {
	fmt.Printf("%T, %d", x, x)
}
