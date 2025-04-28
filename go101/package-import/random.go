package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	fmt.Printf("Hello world.\n")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("Next pseudo-random number is %v.\n", rand.Uint32())
}

func init() {
	fmt.Println("good bye world!")
}

func init() {
	fmt.Println(" this is the third init in random.go file")
}