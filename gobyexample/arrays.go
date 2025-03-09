package main

import "fmt"

func arrays_main() {
	var a [5]int
	fmt.Println("empty: ", a)

	a[4] = 100
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[4])
	fmt.Println("len: ", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b:", b)

	c := [...]int{5, 4, 3, 2, 1, 0, -1}
	fmt.Println("c:", c)

	d := [...]int{100, 3: 400, 500}
	fmt.Println("d:", d)

	var twoD [4][5] int
	for i := 0; i < 4; i++ {
		for j := 0; j < 5; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D: ", twoD)

	// var anotherTwoD [2][3]int
	anotherTwoD := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("another 2 d:", anotherTwoD)
}