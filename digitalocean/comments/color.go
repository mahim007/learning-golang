package main

import "fmt"

const favColor string = "blue" // could have chosen any color

func main() {
	var guess string
	// create an input loop
	for {
		//ask the user to guess my favorite color
		fmt.Println("Guess my favorite color:")
		// try to read a line from the user.
		// print out the error and exit if there is one.
		if _, err := fmt.Scanln(&guess); err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		// did they guess the correct color
		if favColor == guess {
			// they guessed it
			fmt.Printf("%q is my favorite color\n", favColor)
			return
		}

		/* this code is problematic and we are not going to use it, until it's fixed
		someProblematicCode(guess)
		*/

		// wrong! have them guess again
		fmt.Printf("sorry, %q is not my favorite color. guess again\n", guess)
	}
}
