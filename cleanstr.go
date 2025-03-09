package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	str := os.Args[1]

	inWord := false
	for i, char := range str {
		if char > ' ' {
			if inWord && i != 0 && str[i-1] != ' ' && str[i-1] != '\t' && !isTrailingSpaces(str, i) {
				z01.PrintRune(' ')
			}
		} else {
			z01.PrintRune(char)
			inWord = true
		}
	}

	z01.PrintRune('$')
	z01.PrintRune('\n')
}

func isTrailingSpaces(str string, index int) bool {
	for i := index; i < len(str); i++ {
		if str[i] != ' ' && str[i] != '\t' {
			return false
		}
	}
	return true
}
