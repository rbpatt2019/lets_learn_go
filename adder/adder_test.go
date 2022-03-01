package adder

import "testing"

// Test functions
// Always start with Test, followed by the name of the function.
// If the tested function is unexported, then an underscore is frequently used.
// They are part of the package as the code so tha tthey can access unexported
// functions, code, etc.
// They take one parameter `t *testing.T`.
func Test_addNumbers(t *testing.T) {
	result := addNumbers(2, 3)
	if result != 5 {
		t.Error("incorrect result: expected 5, got", result)
	}
}
