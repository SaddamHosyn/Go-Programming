package main

import (
	"fmt"
)

//Write a function that takes a string and return a string containing its first word, followed by a newline ('\n').

//A word is a sequence of characters delimited by spaces or by the start/end of the argument.

func FirstWord(s string) string {
	result := ""

	for _, char := range s {
		//char := s[i]

		if char == ' ' {
			if result != "" { // Stop if we've already found a word
				break
			}
			continue // Skip leading spaces
		}

		result += string(char) // Accumulate characters for the first word
	}

	if result == "" {
		return "\n" // Return newline for empty string or no words
	}

	return result + "\n" // Return the first word followed by a newline
}

func main() {
	fmt.Print(FirstWord("hello there"))            // Output: "hello\n"
	fmt.Print(FirstWord(""))                       // Output: "\n"
	fmt.Print(FirstWord("hello   .........  bye")) // Output: "hello\n"
}
