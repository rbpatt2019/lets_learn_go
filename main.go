package main

import (
	"fmt"
	"math/rand"
)

// Blocks.
// Variables are scoped to the blocks they are defined in.
func main() {
	// So n is only accessible in this if block!
	if n := rand.Intn(10); n == 0 {
		fmt.Println("Low")
	} else if n > 5 {
		fmt.Println("High")
	} else {
		fmt.Println("Good")
	}
	// Also, idiomatic go prefers short if statements that return early.
	// You'd normally write this as a switch.
	// In fact, I'm going to have to commit with no-verify to show this :P.
}
