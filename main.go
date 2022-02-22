package main

type Status int

// Define an enumeration to hold the status codes.
const (
	InvalidLogin Status = iota + 1
	NotFound
)

// Define a struct to hold the error information.
type StatusErr struct {
	Status  Status
	Message string
}

// And a method to guarantee we neet the interface.
func (se StatusErr) Error() string {
	return se.Message
}

// Now, you can return errors with more information about the failure.
// Even when implementing a custom error type,
// always return the standard error.
