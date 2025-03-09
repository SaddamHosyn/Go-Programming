package main

import (
	"fmt"
)

//Write a function called RepeatAlpha that takes a string and displays it repeating each alphabetical character as many times as its alphabetical index.

//   'a' becomes 'a', 'b' becomes 'bb', 'e' becomes 'eeeee', etc...

func RepeatAlpha(s string) string {
	result := ""
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			index := int(char - 'a' + 1)
			for i := 0; i < index; i++ {
				result += string(char)
			}
		} else if char >= 'A' && char <= 'Z' {
			index := int(char - 'A' + 1)
			for i := 0; i < index; i++ {
				result += string(char)
			}
		} else {
			// Ensure non-alphabetic characters are included in the result as-is
			result += string(char)
		}
	}
	return result
}

func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}
