package greetings

import "fmt"

// hello returns greetings for named parameter
// function name starting with Capital letter can be called from function which are not in same package
func Hello(name string) string {
	// returns a greetings that embeds name in the message
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	// := declare and initialize variable in one line
	return message
}