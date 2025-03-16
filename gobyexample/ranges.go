package main

import "fmt"

func ranges_main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}

	fmt.Println("sum:", sum)

	for index, v := range nums {
		if v == 3 {
			fmt.Println("index:", index)
		}
	}

	kvs := map[string]string{
		"a": "apple",
		"b": "banana",
	}

	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Printf("key %s\n", k)
	}

	for i, ch := range "abcABC" {
		fmt.Println(i, ch)
	}
}