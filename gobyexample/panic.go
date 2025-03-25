package main

import "os"

func panic_main() {
	// panic("a problem occurred!")

	_, err := os.Create("./tmp/file")
	if err != nil {
		panic(err)
	}
}