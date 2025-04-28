package main

import (
	"fmt"
	"math/rand"
	"time"
)

func even_odd_main() {
	rand.Seed(time.Now().UnixNano())

	if n := rand.Int(); n % 2 == 0 {
		fmt.Println(n, "is an even number")
	} else {
		fmt.Println(n, "is a odd number")
	}

	n := rand.Int() % 2
	if n % 2 == 0 {
		fmt.Println(n, "is an even number")
	} else {
		fmt.Println(n, "is a odd number")
	}
}