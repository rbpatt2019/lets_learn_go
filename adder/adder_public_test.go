package adder_test

import (
	"testing"

	"github.com/rbpatt2019/lets_learn_go/adder"
)

// Testing your public api.
// 1) package name is `pkg_test`.
// 2) module must be imported.
// 3) no underscore in test function name.
//
// This forces you to interact with your package as black box.
// Can coincide with standard tests, too.
func TestAddNumbers(t *testing.T) {
	result := adder.AddNumbers(2, 3)
	if result != 5 {
		t.Error("incorrect result: expected 5, got", result)
	}
}
