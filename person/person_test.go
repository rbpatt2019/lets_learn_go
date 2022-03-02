package person

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// The package go-cmp is particularly useful for comparing structs.
// It produces a nice diff, and allows for creating custom compares
// if there are odd cases in your data.
func TestCreatePerson(t *testing.T) {
	expected := Person{
		Name: "Dennis",
		Age:  31,
	}
	result := CreatePerson("Dennis", 31)

	// Custom comparer since we cant compare the value of time.Now().
	comparer := cmp.Comparer(func(x, y Person) bool {
		return x.Name == y.Name && x.Age == y.Age
	})

	if diff := cmp.Diff(expected, result, comparer); diff != "" {
		t.Error(diff)
	}
}
