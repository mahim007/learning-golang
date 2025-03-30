package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Sammy Shark"
	fmt.Println(strings.ToLower(s))
	fmt.Println(s)
	fmt.Println(strings.ToUpper(s))

	fmt.Println(strings.HasPrefix(s, "Sammy"))
	fmt.Println(strings.HasSuffix(s, "Shark"))
	fmt.Println(strings.Contains(s, "Sh"))
	fmt.Println(strings.Contains(s, "sh"))
	fmt.Println(strings.Count(s, "s"))
	fmt.Println(strings.Count(s, "S"))

	openSource := "Sammy contributes to the open source"
	// \n\t\r counts as 3 chars
	fmt.Println(len(openSource))

	fmt.Println(strings.Join([]string{"mango", "berry", "banana", "cucumber"}, ", "))

	balloon := "Sammy  has  a  balloon"
	ss := strings.Split(balloon, " ")
	fmt.Printf("%q", ss)
	fmt.Println()

	data := " username password  email date"
	fields := strings.Fields(data)
	fmt.Printf("%q", fields)
	fmt.Println()

	baloon2 := "Sammy has a balloon"
	fmt.Println(strings.ReplaceAll(baloon2, "has", "had"))
}