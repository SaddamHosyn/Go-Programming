package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command( "ls", "-la", "../test" )
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(out))
}
