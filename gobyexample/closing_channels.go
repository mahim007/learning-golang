package main

import "fmt"

func cosing_channels_main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job:", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for i := 1; i <= 5; i++ {
		jobs <- i
		fmt.Println("sent job:", i)
	}

	close(jobs)
	fmt.Println("sent all jobs")

	<-done

	_, ok := <-jobs
	fmt.Println("received more jobs:", ok)
}