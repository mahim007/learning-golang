package main

import (
	"fmt"
	"time"
)

func main() {
	if h := time.Now().Hour(); h < 12 {
		fmt.Println("Now is AM time")
	} else if h > 19 {
		fmt.Println("now is evening time")
	} else {
		fmt.Println("Now is afternoon time")
		h := h
		_ = h
	}

	someMethod()
}

func someMethod() {
	for i := range 10 {
		fmt.Println(i)
	}
}