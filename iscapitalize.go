package main

import (
	"fmt"
)

func main() {
	fmt.Println(IsCapitalized("Hello! How are you?"))
	fmt.Println(IsCapitalized("Hello How Are You"))
	fmt.Println(IsCapitalized("Whats 4this 100K?"))
	fmt.Println(IsCapitalized("Whatsthis4"))
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(IsCapitalized(""))
}

//Write a function IsCapitalized() that takes a string as an argument and returns true if each word in the string begins with either an uppercase letter or a non-alphabetic character.

//If any of the words begin with a lowercase letter return false.
//If the string is empty return false.

func IsCapitalized(m string) bool {
	if m == "" { // Check if the string is empty
		return false // If empty, return false
	}

	// Check if the first character is lowercase
	if m[0] >= 'a' && m[0] <= 'z' {
		return false // If it is, return false (indicating the first word starts with a lowercase letter)
	}

	for i := 1; i < len(m); i++ { // Start from index 1 since we already checked index 0
		if m[i] >= 'a' && m[i] <= 'z' && m[i-1] == ' ' { // Check if current character is lowercase and follows a space
			return false // If it is, return false (indicating a lowercase start of a word)
		}
	}

	return true // If all checks are passed, return true (indicating all words are capitalized or non-alphabetic)
}
