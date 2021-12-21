package main

import (
	"fmt"
)

// For loop.
// These are the only flow loops in go.
func main() {
	// Complete for loop.
	// Idiomatic go uses short if statements.
	// The continue keyword helps achieve this.
	// EG the classic fizzbuzz problem in idiomatic go,
	// Where continue help prevent a deeply nested if/else block.
	for i := 1; i < 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}
		if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}
		fmt.Println(i)
	}
	// Condition only - this is very similar to while in many languages.
	i := 1
	for i < 100 {
		fmt.Println(i)
		i *= 2
	}
	// Infinite for loop.
	// Will fun forever without the use of break.
	i = 100
	for {
		fmt.Println(i)
		if i%5 != 0 {
			break
		}
		i /= 5
	}
	// For-range statement.
	// This is probably the most used,
	// as it is the method for iterating over types that allow iteration.
	// It gives two variables - the position and the value.
	evens := []int{2, 4, 6, 8, 10}
	for _, v := range evens {
		fmt.Println(v)
	}
}
