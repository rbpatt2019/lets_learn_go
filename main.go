package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Defer
//
// In go, resources cleanup is handled with the defer keword.
// These delay calling until the surrounding function exits,
// and guarantee that certain code runs, no matter what.
func Contents(filename string) (string, error) {
	// Then check that the file handle is valid.
	// If it is, we defer closing it.
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	// Introducing idiomatic file reading in Go
	// First create a buffer of bytes.
	// Then read to that buffer.
	// Read returns the number of bytes read, and an err.
	// If err is io.EOF, then the file has read successfully.
	var result []byte
	buf := make([]byte, 2048)
	for {
		n, err := f.Read(buf)
		result = append(result, buf[:n]...)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
	}
	return string(result), nil
}

func main() {
	// Os.Args is a slice containing parameters passed to the script
	// Os.Args[0] is the filename.
	if len(os.Args) < 2 {
		log.Fatal("No file specified")
	}

	// Once we know we have a file we open it.
	results, err := Contents(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(results)
}
