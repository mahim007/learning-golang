package main

import (
	"fmt"
	"maps"
)

func maps_main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map: ", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len: ", len(m))

	delete(m, "k2")
	fmt.Println("map: ", m)

	clear(m)
	fmt.Println("map: ", m)

	m["k2"] = 20
	v, prs := m["k2"]
	fmt.Println("v: ", v, " prs: ", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("new map: ", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}

	n3 := map[string]int{}
	fmt.Println("n3:", n3)

	n4 := make(map[string]int)
	fmt.Println("n4:", n4)

}
