package main

import "fmt"

func strngs_main() {
	a := "Hello, 世界"
	for i, c := range a {
		fmt.Printf("%d: %s\n", i, string(c))
	}

	fmt.Println("length of 'Hello, 世界': ", len(a))

	var pi float64 = 3.14
	fmt.Println(pi)

	names := [3]string{"tom", "cat", "go"}
	fmt.Println(names)

	nums := []int{-2,-1,0,1,2,3}
	fmt.Println(nums)
	floats := []float64{1.11, -3.22, 4.550940930490394992597349759347974359732040982349287598374975}
	fmt.Println(floats)

	strings := []string{"this", "is", "a", "test"}
	fmt.Println(strings)
	strings = append(strings, "go")
	fmt.Println(strings)

	// map[key]value{}
	users := map[string]string{
		"1": "user 1",
		"2": "user 2",
		"3": "user 3",
		"4": "user 4",
	}
	fmt.Println(users)
	fmt.Println(users["1"])

}