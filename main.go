package main

import (
	"errors"
	"fmt"
	"net/http"
)

func LogOutput(message string) {
	fmt.Println(message)
}

type SimpleDataStore struct {
	userData map[string]string
}

func (sds SimpleDataStore) UserNameForid(id string) (string, bool) {
	name, ok := sds.userData[id]
	return name, ok
}

// Factory function for creating a new SDS.
// In the real world,
// this data would be retireved from somewhere,
// perhaps a database.
func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		userData: map[string]string{
			"1": "Fred",
			"2": "Mary",
			"3": "Joe",
		},
	}
}

// To make sure that our logic does not depend on specific implementations,
// we create interfaces for the desired functionality.
type DataStore interface {
	UserNameForid(id string) (string, bool)
}

type Logger interface {
	Log(message string)
}

// In order for a function to meet the logger interfaces,
// we must implement an apdaptor type...
type LoggerAdapter func(message string)

// And an appropriate method.
// Now, any func that has the signature func(message string)
// can be used to emet the interface.
func (lg LoggerAdapter) Log(message string) {
	lg(message)
}

// Implement business logic
// Notice that there are no dependencies here,
// as we require interfaces!
type SimpleLogic struct {
	l  Logger
	ds DataStore
}

func (sl SimpleLogic) Hello(id string) (string, error) {
	sl.l.Log("in Hello for " + id)
	name, ok := sl.ds.UserNameForid(id)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "Hello " + name, nil
}

func (sl SimpleLogic) Goodbye(id string) (string, error) {
	sl.l.Log("in Goodbye for " + id)
	name, ok := sl.ds.UserNameForid(id)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "Goodbye " + name, nil
}

// Factory function for creating a SimpleLogic instances.
func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic {
	return SimpleLogic{
		l:  l,
		ds: ds,
	}
}

// In order to make sure the controller does not depend on specific implementations,
// we implement an interface.
type Logic interface {
	Hello(id string) (string, error)
}

// Implement controller.
// Again, notice that there are no dependencies,
// as we specify interfaces.
type Controller struct {
	l     Logger
	logic Logic
}

// Controller handles communicating with server.
// Do not use query parameters in real life for authentication!
func (c Controller) SayHello(w http.ResponseWriter, r *http.Request) {
	c.l.Log("In SayHello")
	id := r.URL.Query().Get("user_id")
	message, err := c.logic.Hello(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(message))
}

// Factory function for creating a controller.
func NewController(l Logger, logic Logic) Controller {
	return Controller{
		l:     l,
		logic: logic,
	}
}

// Wire it all up.
// ACCEPT INTERFACES, RETURN STRUCTS!
// Because of this principle,
// this is the only place that needs to change if/when
// we swap implementations.
func main() {
	l := LoggerAdapter(LogOutput)
	ds := NewSimpleDataStore()
	logic := NewSimpleLogic(l, ds)
	c := NewController(l, logic)
	http.HandleFunc("/hello", c.SayHello)
	http.ListenAndServe(":8080", nil)
}
