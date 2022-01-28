package main

import (
	"fmt"
)

type MailCategory int

// Generally, use iota for interaml purposes only,
// where constants are referred to by name rather than value.
//
// Or put another way:
// only use iota when you want to differentiate between values,
// but don't care what those values are.
// If the actual value matters,
// specify it.
const (
	Uncategorised MailCategory = iota
	Personal
	Spam
	Social
	Advertisements
)

func main() {
	fmt.Println(Uncategorised)
}
