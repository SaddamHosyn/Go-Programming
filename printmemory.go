package main

import (
	"github.com/01-edu/z01"
)

/*
Write a function that takes (arr [10]byte), and displays the memory as in the example.

After displaying the memory the function must display all the ASCII graphic characters. The non printable characters must be replaced by a dot.

The ASCII graphic characters are any characters intended to be written, printed, or otherwise displayed in a form that can be read by humans, present on the ASCII encoding.*/

// Converts a byte to its corresponding hexadecimal character

func hextoRune(s byte) rune {
	base := "0123456789abcdef"
	return rune(base[s])
}
func PrintMemory(arr [10]byte) {
	for i, char := range arr {
		firstdig := char / 16
		seconddig := char % 16
		z01.PrintRune(hextoRune(firstdig))
		z01.PrintRune(hextoRune(seconddig))

		if (i+1)%4 == 0 || i == len(arr)-1 {
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(' ')
		}
	}
	for _, char := range arr {
		if char > 32 && char < 127 {
			z01.PrintRune(rune(char))
		} else {
			z01.PrintRune('.')

		}

	}
	z01.PrintRune('\n')
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
