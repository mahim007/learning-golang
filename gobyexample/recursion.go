package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}

func recursion_main() {
	fmt.Println(fact(15))

	var fib func(n int) int
	fmt.Println("fib:", fib)
	fmt.Println(fib == nil)

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		return fib(n - 2) + fib(n - 1)
	}

	fmt.Println(fib(10))
}