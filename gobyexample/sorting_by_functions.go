package main

import (
	"cmp"
	"fmt"
	"slices"
)

func sorting_by_functions_main() {
	fruits := []string{"peach", "banana", "kiwi"}

	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	type Person struct {
		name string
		age int
	}

	people := []Person{
		{name: "mahim", age: 30},
		{name: "nadia", age: 26},
		{name: "masum", age: 35},
	}

	slices.SortFunc(people, func(a, b Person) int {
		return cmp.Compare(a.age, b.age)
	})

	fmt.Println(people)
}