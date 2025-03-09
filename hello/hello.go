package main

import (
	"fmt"
	"go-test/greetings"
	"log"
)

func main()  {
	log.SetPrefix("Greetings: ")
	log.SetFlags(0)

	names := []string{
		"Mahim",
		"Nadia",
	}


	message, err := greetings.Hellos(names)
	
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
	for _, name := range names {
		fmt.Println(name + " : " + message[name])
	}
}