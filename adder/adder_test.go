package adder

import (
	"os"
	"testing"
)

// Cleanup is, in most ways a test specific defer.
// It runs when the calling test completes.
// If there are multiple, they run in a
// "last added, first called" order.
// They are most useful when a helper function creates temporary resources.
func createFile(t *testing.T) (string, error) {
	f, err := os.Create("tempfile")
	if err != nil {
		return "", err
	}
	// Do something to f...
	t.Cleanup(func() {
		os.Remove(f.Name())
	})
	return f.Name(), nil
}

func Test_addNumbers(t *testing.T) {
	// createFile would then be used here.
	result := addNumbers(2, 3)
	if result != 5 {
		t.Error("incorrect result: expected 5, got", result)
	}
}
