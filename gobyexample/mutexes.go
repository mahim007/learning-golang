package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mu sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string)  {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func rate_limiting_main()  {
	c := Container {
		counters: map[string]int{
			"a": 0,
			"b": 0,
		},
	}

	var wg sync.WaitGroup

	doIncrement := func (name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}

		wg.Done()

		fmt.Println(name, n, "done")
	}

	wg.Add(3)

	go doIncrement("a", 5000000)
	go doIncrement("a", 7000000)
	go doIncrement("b", 9000000)

	wg.Wait()
	fmt.Println(c.counters)

}