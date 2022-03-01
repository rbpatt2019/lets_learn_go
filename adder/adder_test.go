package adder

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var testTime time.Time

// TestMain is run once per package to provide common test setup.
// It (basically) always has the structure indicated below.
// It's most useful for either setting up data in an external repository
// ie a database,
// or for package-level variable initialisation (as here).
// But you should really try to avoid the latter anyways!
func TestMain(m *testing.M) {
	fmt.Println("Test setup")
	testTime = time.Now()

	fmt.Println("Run tests")
	exit := m.Run()

	fmt.Println("Test cleanup")
	os.Exit(exit)
}

func Test_addNumbers(t *testing.T) {
	result := addNumbers(2, 3)
	if result != 5 {
		t.Error("incorrect result: expected 5, got", result)
	}
	fmt.Println(testTime)
}
