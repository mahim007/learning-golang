package main

import "fmt"

func range_over_channels_main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	// queue <- "three"
	close(queue)

	for ele := range queue {
		fmt.Println(ele)
	}
}