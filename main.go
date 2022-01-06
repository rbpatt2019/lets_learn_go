package main

import (
	"fmt"
	"strconv"
)

// Functions can also be defined as types...
type opFuncType func(int, int) int

// Functions are values.
//
// This means they can be passed around,
// allowing for some fun behaviour.
func add(i, j int) int {
	return i + j
}

func sub(i, j int) int {
	return i - j
}

func mul(i, j int) int {
	return i * j
}

func div(i, j int) int {
	return i / j
}

var opMap = map[string]opFuncType{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

// Always check your errors!
func main() {
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "3"},
		{"5"},
	}
	for _, exp := range expressions {
		if len(exp) != 3 {
			fmt.Println("Invalid expression:", exp)
			continue
		}

		// Use strconv.Atoi to convert to an integer.
		p1, err := strconv.Atoi(exp[0])
		if err != nil {
			fmt.Println(err)
			continue
		}

		op := exp[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("Unsupported operator:", op)
			continue
		}

		p2, err := strconv.Atoi(exp[2])
		if err != nil {
			fmt.Println(err)
			continue
		}

		result := opFunc(p1, p2)
		fmt.Println(result)
	}
}
