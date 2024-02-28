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
	names := []string{"tom", "dick", "harry"}
	message, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}