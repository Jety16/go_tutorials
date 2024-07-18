package greetings

import "fmt"

func Hello(name string) string {
	// the := operator is a shortcut for declaring and initializing a variable in one line
	// (Go uses the value on the right to determine the variable's type). 
	// var message string is the other way

	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}