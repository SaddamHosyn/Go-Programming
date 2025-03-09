/*
	Write a program that takes 3 arguments, the first argument is a string in which a letter (the 2nd argument) will be replaced by another one (the 3rd argument).

If the number of arguments is different from 3, the program displays nothing.

If the second argument is not contained in the first one (the string) then the program rewrites the string followed by a newline ('\n').
*/
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	first := os.Args[1]
	second := os.Args[2]
	third := os.Args[3]

	if len(second) != 1 || len(third) != 1 {
		return
	}

	result := ""
	for _, char := range first {
		if string(char) == second {
			result += third // this part helping to replace the letter
		} else {
			result += string(char) // helping to make a word
		}
	}

	for _, c := range result {
		z01.PrintRune(c)
	}
	// Add the newline
	z01.PrintRune('\n')
}
