package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {

		z01.PrintRune('1')
		z01.PrintRune('\n')
		return
	}

	str1 := os.Args[1]
	str2 := os.Args[2]

	i := 0
	j := 0

	for i < len(str1) && j < len(str2) {
		if str1[i] == str2[j] {
			i++
		}
		j++
	}

	if i == len(str1) {
		z01.PrintRune('1')
	} else {
		z01.PrintRune('0')
	}
	z01.PrintRune('\n')

}
