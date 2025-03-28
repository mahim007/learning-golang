package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func waitgroups_main() {
	var wg sync.WaitGroup

	for i := 1; i <=5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()
}