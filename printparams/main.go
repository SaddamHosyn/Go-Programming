package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for _, val := range os.Args[1:] {
		for _, char := range val {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
