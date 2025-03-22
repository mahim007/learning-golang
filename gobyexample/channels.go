package main

import "fmt"

func channels_main() {
	messages := make(chan string)

	go func() {
		sum := 0
		for i := 1; i <= 3; i++ {
			sum += i
		}

		msg := fmt.Sprintf("the result is %v", sum)
		messages <- msg
	}()

	msg := <- messages
	fmt.Println("Output:", msg)
}