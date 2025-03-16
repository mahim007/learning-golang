package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func vals() (int, int) {
	return 3, 7
}

func sums(nums ...int) int {
	fmt.Println("nums:", nums)

	total := 0
	for _, num := range nums {
		total += num
	}

	return total
}

func functions_main() {
	res := plus(1, 2)
	fmt.Println("1 + 2:", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3:", res)

	a, b := vals()
	fmt.Println("a", a, "b", b)

	_, y := vals()
	fmt.Println("y:", y)

	nums := []int{1, 2, 3, 4, 6}
	total := sums(nums...)
	fmt.Println("total:", total)
}