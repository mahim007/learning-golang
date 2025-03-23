package main

import "fmt"

func channel_buffering_main() {
	messages := make(chan string, 2)
	messages <- "mahim"
	messages <- "nadia"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}