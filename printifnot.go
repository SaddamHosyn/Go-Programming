package main

import (
	"fmt"
)

//Write a function that takes a string as an argument and returns the letter G if the argument length is less than 3, otherwise returns Invalid Input followed by a newline \n. If it's an empty string return G followed by a newline \n.

func PrintIfNot(str string) string {
	if len(str) == 0 || len(str) < 3 {
		return "G\n"
	}
	return "Invalid Input\n"
}

func main() {
	fmt.Print(PrintIfNot("abcdefz")) // Output: "Invalid Input\n"
	fmt.Print(PrintIfNot("abc"))     // Output: "Invalid Input\n"
	fmt.Print(PrintIfNot(""))        // Output: "G\n"
	fmt.Print(PrintIfNot("14"))      // Output: "G\n"
}
