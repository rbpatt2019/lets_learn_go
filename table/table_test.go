package table

import (
	"testing"
)

// Table testing
// Helps to abstract away repetitive code for testing multiple
// similar situations.
// 1) Create a slice of anonymous structs.
// Fields must be name, parameters, expected, error (if any)
// 2) Call `t.Run` for each element of the slice.
// `t.Run` gets two arguments - a name, and the test function.
func TestDoMathTable(t *testing.T) {
	data := []struct { //nolint:govet // Priortise readability for testing.
		name     string
		num1     int
		num2     int
		op       string
		expected int
		errMsg   string
	}{
		{"add", 2, 3, "+", 5, ""},
		{"subtract", 3, 2, "-", 1, ""},
		{"multiply", 2, 3, "*", 6, ""},
		{"divide", 4, 2, "/", 2, ""},
		{"zero_division", 4, 0, "/", 0, "division by zero"},
		{"default", 2, 3, "$", 0, "unknown operator $"},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result, err := DoMath(d.num1, d.num2, d.op)
			if result != d.expected {
				t.Errorf("Expected %d, got %d", d.expected, result)
			}
			// This is a very fragile way to test for errors.
			// Mayhaps the package could use a custom type,
			// Then check for with errors.Is or errors.As?
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != d.errMsg {
				t.Errorf("Expected error %s, got error %s", d.errMsg, errMsg)
			}
		})
	}
}
