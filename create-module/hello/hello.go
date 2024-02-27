package main

import (
	"fmt"
	"example.com/greetings"
)

func main(){
	message := greetings.Hello("Manish")
	fmt.Println(message)
}