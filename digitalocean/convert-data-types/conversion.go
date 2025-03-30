package main

import (
	"fmt"
	"strconv"
)

func conversion_main() {
	fmt.Printf("%T %d\n", 8, 8)

	var small int8 = 8
	var big int32 = int32(small)

	fmt.Printf("%T %d\n", small, small)
	fmt.Printf("%T %d\n", big, big)


	var anotherBig int = 129
	var anotherSmall int8 = int8(anotherBig)

	fmt.Println(anotherSmall)

	a := 5.00
	b := 2
	res := a / float64(b)
	fmt.Println(res)

	fmt.Printf("%q\n",strconv.Itoa(int(a)))

	fmt.Println(fmt.Sprint(421.034))

	f := 5524.53
	fmt.Println(fmt.Sprint(f))

	fmt.Println(strconv.ParseInt("100", 10, 32))
	fmt.Println(strconv.Atoi("100.21"))
}