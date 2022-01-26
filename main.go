package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Using slices as buffers
//
// Just because Go has a garbage collector doesn't mean we shouldn't
// try to reduce it's workload!
// Here, we create a single buffer slice of bytes
// and repeatedly read to that!
func readFile(name string) error {
	file, err := os.Open(name)
	if err != nil {
		return err
	}
	defer file.Close()

	buf := make([]byte, 100)
	for {
		count, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if count == 0 {
			return nil
		}
		fmt.Println(buf[:count])
	}
	return nil
}

func main() {
	err := readFile("lorem.txt")
	if err != nil {
		log.Fatal(err)
	}
}
