package person

import "time"

type Person struct {
	DateAdded time.Time
	Name      string
	Age       int
}

func CreatePerson(name string, age int) Person {
	return Person{
		Name:      name,
		Age:       age,
		DateAdded: time.Now(),
	}
}
