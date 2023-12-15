package main

import (
	"fmt"
	"log"
	"tutorial/greetings"

	"rsc.io/quote"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Elliott", "Frodo", "Bilbo"}

	// Request greeting messages
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	// Print all the greeting messages
	for _, message := range messages {
		fmt.Println(message)
	}

	// Send a quote
	fmt.Println(quote.Go())
}
