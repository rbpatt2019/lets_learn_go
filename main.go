package main

import "fmt"

type LogicProvider struct{}

func (lp LogicProvider) Process(data string) string {
	// Do the things.
	return fmt.Sprintf("New string: %s", data)
}

// Interfaces
//
// Type safe duck typing.
// Interfaces specify the methods a struct needs to have.
// But, they do not care how those methods are provided.
// Here, LogicProvider meets the interface.
// Yet, if we changed to an utterly un-related provider,
// it would still meet the interface as long as there was
// a Process method.
type Logic interface {
	Process(data string) string
}

// Essentially,
// only the client knows about the interface.
// Thus, we get the flexibility of duck typing
// with the bonus of type checking!
type Client struct {
	L Logic
}

func (c Client) Program(data string) string {
	return c.L.Process(data)
}

func main() {
	c := Client{
		L: LogicProvider{},
	}

	// In the real world,
	// we'd get this data from somewhere rather than hard-code it.
	data := "Hello World"
	c.Program(data)
}
