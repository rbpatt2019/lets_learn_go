package main

import "fmt"

type MyInt int

// Type assertions
//
// Name the concorete type that implemented the interface.
// If your type assertion is wrong,
// the code will panic.
// This can be avoided with the 'comma, ok idioom'.
func main() {
    var i interface{}
    var mine MyInt = 20
    i = mine
    // This line is the type assertion.
    i2, ok := i.(MyInt)
    if !ok {
        fmt.Println("What were you doing?")
    }
    fmt.Println(i2 + 1)
}
