package main

import (
	"fmt"
	"strconv"
)

// Write a function that takes a string and returns a new string that replaces every character with the number of duplicates and the character itself, deleting the extra duplications. The letters are from the latin alphabet list only. Any other character, symbols, shall not be tested.

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}

func ZipString(s string) string {
	i := 0
	result := ""

	for i < len(s) {
		char := s[i]
		count := 1
		for i+1 < len(s) && char == s[i+1] {
			count++
			i++
		}
		result += strconv.Itoa(count) + string(char)
		i++
	}
	return result
}
