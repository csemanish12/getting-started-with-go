package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main(){
	// set properties for predefined logger
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	message, err := greetings.Hello("Manish")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}