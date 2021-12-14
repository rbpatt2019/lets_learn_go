package main

import (
	"fmt"
)

// Go requires explicit type conversion
// So ints can't be compared to floats, etc.
// This is strict enough to prevent `int` from comparing
// with say `int32`.
func main() {
	var x int = 10
	var y float64 = 30.2
	fmt.Printf("x: %T, %d\n", x, x)
	fmt.Printf("y: %T, %f\n", y, y)

	var z float64 = y + float64(x)
	var a int = x + int(y)
	fmt.Printf("z: %T, %f\n", z, z)
	fmt.Printf("a: %T, %d\n", a, a)
}
