package main

import "fmt"

func zeroval(val int)  {
	val = 0
}

func zeroptr(val *int)  {
	*val = 0
}

func pointers_main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println(i)

	fmt.Println("pointer:", &i)

}