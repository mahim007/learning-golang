package main

import (
	"fmt"
	"unicode/utf8"
)

func strings_runes_main() {
	// const s = "สวัสดี"
	const s = "mahim"
	fmt.Println(s)
	fmt.Println("len:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%d -> %x\n", i, s[i])
	}

	fmt.Println("Rune count: ", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\n Using DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}