package main

import "fmt"

func main() {
	messages := make(chan string, 2)
	messages <- "mahim"
	messages <- "nadia"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}