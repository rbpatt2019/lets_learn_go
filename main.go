package main

import "fmt"

type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

// Embedding
//
// Is not inheritance!
// Employee has no name assigned to it,
// Making it an embedded field.
// Any method declared on an embedded field
// is promoted to the containing struct.
type Manager struct {
	Employee
	Reports []Employee
}

// Which makes the following valid!
func main() {
	m := Manager{
		Employee: Employee{
			Name: "Bob",
			ID:   "123",
		},
		Reports: []Employee{},
	}
	fmt.Println(m.ID)
}
