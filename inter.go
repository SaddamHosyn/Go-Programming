package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isContain(s string, r rune) bool {
	for _, char := range s {
		if char == r {
			return true
		}
	}
	return false
}

func main() {

	if len(os.Args) != 3 {
		return
	}

	str1 := os.Args[1]
	str2 := os.Args[2]

	var result string

	for _, char := range str1 {
		if isContain(str2, char) && !isContain(result, char) {
			result += string(char)
		}
	}

	for _, char := range result {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')

}
