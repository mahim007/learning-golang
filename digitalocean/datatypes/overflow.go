package main

import "fmt"

func overflow_main() {
	var maxUint32 uint32 = 4294967295 + 1
	fmt.Println(maxUint32)
}