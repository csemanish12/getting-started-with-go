package greetings

import (
	"fmt"
	"errors"
	"math/rand"

)

// hello returns greetings for named parameter
// function name starting with Capital letter can be called from function which are not in same package
func Hello(name string) (string, error) {
	// returns error if name is empty
	if name == "" {
		return "", errors.New("Empty name")
	}
	// returns a greetings that embeds name in the message
	message := fmt.Sprintf(randomFormat(), name)
	// := declare and initialize variable in one line
	return message, nil
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }
	return formats[rand.Intn(len(formats))]
}