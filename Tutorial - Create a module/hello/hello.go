package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	// a flag to disable printing the time, source file, and line number.
	// log.SetFlags(0)

	message, err := greetings.Hello("Yao")
	fmt.Println(message)

	names := []string{"xiao", "hui"}
	message1, err1 := greetings.Hellos(names)
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Println(message1)

	message2, err2 := greetings.Hello("")
	if err2 != nil {
		log.Fatal(err)
	}
	fmt.Println(message2)
}
