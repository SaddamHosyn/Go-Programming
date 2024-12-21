package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	c := 0
	for _, val := range args {
		if val == "01" || val == "galaxy" || val == "galaxy 01" {
			c++
		}
	}
	if c >= 1 {
		fmt.Println("Alert!!!")
	}
}
