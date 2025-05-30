package main

import "fmt"

func non_blocking_channel_operations_main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received messages", msg)
	default:
		fmt.Println("No message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent messages:", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <- messages:
		fmt.Println("received message:", msg)
	case sig := <-  signals:
		fmt.Println("received signal:", sig)
	default:
		fmt.Println("no activity")
	}
}