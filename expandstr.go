package main

import (
	"os"

	"github.com/01-edu/z01"
)



// Write a program that takes a string and displays it with exactly three spaces between each word, with no spaces nor tabs at neither the beginning nor the end.

// The string will be followed by a newline ('\n').

// A word, in this exercise, is a sequence of visible characters.

// If the number of arguments is not 1, or if there are no word, the program displays nothing.



func main() {
	if len(os.Args) != 2 {
		return
	}

	input := os.Args[1]
	result := ""
	inWord := false

	for _, char := range input {
		//if char != ' ' && char != '\t' {
		if char > ' ' {
			result += string(char)
			inWord = true
		} else if inWord {
			// Only add exactly three spaces when transitioning out of a word
			result += "   "
			inWord = false
		}
	}

	// Remove any extra spaces at the end of result
	if len(result) > 3 && result[len(result)-3:] == "   " {
		result = result[:len(result)-3]
	}

	// Print the result
	for _, char := range result {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
