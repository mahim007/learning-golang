package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	lines_yesterday := "50"
	lines_today := "82"

	yesterday, err := strconv.Atoi(lines_yesterday)
	if err != nil {
		log.Fatal(err)
	}

	today, err := strconv.Atoi(lines_today)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(today - yesterday)

	// a := "not a number"
	// b, err := strconv.Atoi(a)
	// fmt.Println(b)
	// fmt.Println(err)

	a := "my string"

	b := []byte(a)

	c := string(b)

	fmt.Println(a)

	fmt.Println(b)

	fmt.Println(c)

	mp := map[string]string{"name": "ashraful", "age": "31"}
	val, ok := mp["data"]
	fmt.Println(val, ok)
}