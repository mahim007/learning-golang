package main

import "fmt"

func for_main() {
	var i int = 1
	count := 3

	for i <= count {
		fmt.Println(i)
		i = i + 1
	}

	fmt.Println()

	for j := 0; j <= count; j++ {
		fmt.Println(j)
	}

	fmt.Println()

	for i := range count {
		fmt.Println(i)
	}

	for {
		fmt.Println("This is a loop!")
		break
	}
	fmt.Println()

	for n := range 10 {
		if n % 2 == 0 {
			continue
		}

		fmt.Println(n)
	}
}