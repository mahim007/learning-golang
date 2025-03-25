package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	// fmt.Println("starting worker:", id)
	for job := range jobs {
		fmt.Println("worker", id, " started job", job)
		time.Sleep(1 * time.Second)
		fmt.Println("worker", id, " finished job", job)
		results <- job * 2
	}

	// fmt.Println("returning from worker", id)
}

func worker_pool_main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 1; i <= numJobs; i++ {
		// fmt.Println("work:", i , " committed with:", <-results)
		<-results
	}

	// time.Sleep(5 * time.Second)
}