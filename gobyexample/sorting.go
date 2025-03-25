package main

import (
	"fmt"
	"slices"
)

func sorting_main() {
	strs := []string{"c", "b", "a"}
	slices.Sort(strs)
	fmt.Println("strings:", strs)

	ints := []int{5, 4, 3, 2, 1}
	slices.Sort(ints)
	fmt.Println(ints)
	
	fmt.Println(slices.IsSorted(ints))
}