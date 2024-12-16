package main

import (
	"fmt"
	"os"
)

func main() {
	hasError := false
	option := os.Args[1]
	if option != "-c" {
		os.Exit(1)
	}
	total := os.Args[2]
	totalInt := AtoiInt(total)
	args := os.Args[3:]
	if len(args) < 1 {
		fmt.Printf("Error")
		os.Exit(1)
	} else {
		for i, c := range args {
			file, err := os.Open(c)
			fileState, err2 := file.Stat()
			if err != nil {
				fmt.Printf("%v", err.Error())
				hasError = true
				fmt.Printf("%v", "\n")
			} else if err2 != nil {
				hasError = true
				fmt.Printf("%v", err.Error())
			} else {
				arr := make([]byte, fileState.Size())
				file.Read(arr)
				file.Close()
				if len(args) > 1 {
					if i > 0 {
						fmt.Printf("%v", "\n")
					}
					fmt.Printf("==> %v <==", c)
					fmt.Printf("%v", "\n")
					if totalInt > len(arr) {
						fmt.Printf("%v", string(arr))
					} else {
						fmt.Printf("%v", string(arr[len(arr)-totalInt:]))
					}
				} else {
					if totalInt > len(arr) {
						fmt.Printf("%v", string(arr))
					} else {
						fmt.Printf("%v", string(arr[len(arr)-totalInt:]))
					}
				}
			}
		}
		if hasError {
			os.Exit(1)
		} else {
			os.Exit(0)
		}
	}
}

func AtoiInt(s string) int {
	var i int = 0
	if len(s) == 0 || s == "-" || s == "+" {
		return i
	}
	for _, c := range s {
		if !((c >= '0' && c <= '9') || c == '-' || c == '+') {
			return i
		}
	}
	if s[0] == '-' && s[1] == '-' {
		return i
	}
	if s[0] == '-' && len(s) != 1 {
		s = s[1:]
		i = DoAtoi1(s)
		i = -i
	} else if s[0] == '+' && len(s) != 1 {
		s = s[1:]
		i = DoAtoi1(s)
	} else {
		i = DoAtoi1(s)
	}
	return i
}

func DoAtoi1(str string) int {
	var i int = 0
	for _, c := range str {
		if !(c >= '0' && c <= '9') {
			return 0
		}
		i = i*10 + int(c-'0')
	}
	return i
}
