package main

import "fmt"

// Maps.
func main() {
	// 0 value is nil.
	// A nil map can be read, but not written.
	var x map[string]int
	fmt.Printf("%T, %v\n", x, x)
	// To create a map literal,
	// which has length 0 but can be read and wirtten...
	y := map[string]int{}
	fmt.Printf("%T, %v\n", y, y)
	// As with slices, if the size is known, but not the values, use make...
	z := make(map[string]int, 3)
	fmt.Printf("%T, %v\n", z, z)
	// Reading and writing will fell familiar.
	z["cats"] = 1
	z["dogs"] = 2
	z["people"] = 3
	fmt.Println(z)
	// But accessing a map key will always return.
	// If the key does not exist, then the 0 value of the value type is returned.
	// Use the comma-ok idiom if you need to know if the key exists.
	v, ok := z["birds"]
	if ok {
		fmt.Println(v)
		return
	}
	fmt.Println("V does not exist")
	// Deleting requires delete...
	delete(z, "people")
	fmt.Printf("%T, %v\n", z, z)
	// As go lacks sets, a map can be used instead...
	set := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		set[v] = true
	}
}
