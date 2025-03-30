package main

import "fmt"

func string_main() {
	a := `Say "hello" to go`
	fmt.Println(a)

	b := `Go is expressive, concise, clean, and efficient.
Its concurrency mechanisms make it easy to write programs
that get the most out of multi-core and networked machines,
while its {novel} type system enables flexible and modular\n
program (construction). Go 'compiles' quickly to machine code\r
yet has the [convenience] of garbage collection and the power\t
of run-time reflection. It's a fast, "statically" typed,
compiled language that feels like a dynamically typed,
interpreted language.`
	fmt.Println(b)

	c := "say \"hello\" to go"
	fmt.Println(c)
	// 	digitalocean\working-with-strings\strings.go:20:21: syntax error: unexpected name to at end of statement
	// digitalocean\working-with-strings\strings.go:20:28: newline in string
	// digitalocean\working-with-strings\strings.go:21:6: undefined: fmt.println

	fmt.Println("sammy" + "shark")
	fmt.Println("sammy" + "shark")

	fmt.Println("Sammy says, \"I like to use the `fmt` package.\"")
}