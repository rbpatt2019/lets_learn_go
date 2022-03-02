package bench

import (
	"fmt"
	"testing"
)

// Necessary for benchmarking.
var blackhole int

func TestFileLen(t *testing.T) {
	result, err := FileLen("testdata/data.txt", 10000)
	if err != nil {
		t.Fatal(err)
	}
	if result != 65204 {
		t.Error("expected 65204, got", result)
	}
}

// Benchmark convention are similar to test.
// Function name begins with Benchmark
// Takes one parameter b of type `*testing.B`.
// Must for over i < b.N.
//
// The package variable blackhole is used to prevent the GC from abstracting out the call.
func BenchmarkFileLen(b *testing.B) {
	for _, v := range []int{1, 10, 100, 10000, 100000} {
		b.Run(fmt.Sprintf("FileLen-%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result, err := FileLen("testdata/data.txt", v)
				if err != nil {
					b.Fatal(err)
				}
				blackhole = result
			}
		})
	}
}
