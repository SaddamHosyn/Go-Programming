package main

import (
	"os"

	"github.com/01-edu/z01"
)

//Write a program that takes one or more arguments and that, for each argument, puts the last letter of each word in uppercase and the rest in lowercase. It displays the result followed by a newline ('\n'). If there are no argument, the program displays nothing.

func isUpper(r rune) bool {
	if r >= 'A' && r <= 'Z' {
		return true
	}
	return false
}

func isLower(r rune) bool {
	if r >= 'a' && r <= 'z' {
		return true
	}
	return false
}

func reverseStrCap(s string) string {
	result := []rune(s)
	for i := 0; i < len(result); i++ {
		if i == len(result)-1 || result[i+1] == ' ' {
			if isLower(result[i]) {
				result[i] -= 32
			}
		} else if isUpper(result[i]) {
			result[i] += 32
		}
	}

	return string(result)
}

func main() {
	if len(os.Args) > 1 {
		for _, args := range os.Args[1:] {
			result := reverseStrCap(args)

			for _, r := range result {
				z01.PrintRune(r)
			}
			z01.PrintRune('\n')
		}
	}
}
