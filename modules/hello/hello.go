package main

import ( 
	"fmt"
	"example.com/greetings"
)

func main() {
	message:= greetings.Hello("Juancito")
	fmt.Println(message)
}