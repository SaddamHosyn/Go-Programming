package main

import (
	"os"

	"github.com/01-edu/z01"
)

//Write a program that takes one or more arguments and that, for each argument, puts the last letter of each word in uppercase and the rest in lowercase. It displays the result followed by a newline ('\n'). If there are no argument, the program displays nothing.

func main() {
	//The outer loop handles the words (arguments).
	for _, arg := range os.Args[1:] {
		//The inner loop handles the characters in each word.
		for i, char := range arg {
			if i == len(arg)-1 || arg[i+1] == ' ' {
				// **Convert to uppercase manually for A-Z
				if char >= 'a' && char <= 'z' {
					char = char - 32
				}
				// **Convert to lowercase manually for A-Z
			} else if char >= 'A' && char <= 'Z' {
				char = char + 32
			}
			z01.PrintRune(char)

		}

		z01.PrintRune('\n')
	}

}
