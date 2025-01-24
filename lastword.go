package main

import (
	"fmt"
)

func main() {
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(LastWord(" lorem,ipsum "))
	fmt.Print(LastWord(" "))
}

//Write a function LastWord that takes a stringand returns its last word followed by a\n`. A word is a section of string delimited by spaces or by the start/end of the string.

func LastWord(str string) string {
	result := ""
	wordfound := false
	for i := len(str) - 1; i >= 0; i-- {
		char := str[i]
		if char != ' ' {
			result = string(char) + result
			wordfound = true
		} else if wordfound {
			break
		}
	} 
	return result + "\n"

	

}
