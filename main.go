package main

import (
	"fmt"
)

// Switch statements.
func main() {
	// Best used when there is a relationship between the variables.
	// Cases DO NOT fall through,
	// and empty cases do nothing.
	words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}
	for _, v := range words {
		switch size := len(v); size {
		case 1, 2, 3, 4:
			fmt.Println(v, "is a short word")
		case 5:
			fmt.Println(v, "is just right")
		case 6, 7, 8, 9:
		default:
			fmt.Println(v, "is a long word")
		}
	}

	// Switches can be cases on any expression (blank)
	// Here's Fizzbuzz re-written with a blank switch.
	for i := 1; i < 25; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}
