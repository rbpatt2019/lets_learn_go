package main

import "fmt"

// Recovers must in deferred functions,
// as only defers run after panics.
// Idiomatically,
// It would be better to check for div by 0 and report the appropriate error.
func div60(i int) {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println(v)
		}
	}()
	fmt.Println(60 / i)
}

// Panic/recover should be used more or less only to:
// 1) log system failures to monitoring software.
// 2) prevent panics from leaking your API by turning them into errors.
func main() {
	for _, val := range []int{1, 0, 2} {
		div60(val)
	}
}
