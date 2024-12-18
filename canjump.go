package main

import (
	"fmt"
)

func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(CanJump(input1))

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(CanJump(input2))

	input3 := []uint{0}
	fmt.Println(CanJump(input3))
}

func max(a, b int) int {

	if a > b {
		return a
	}

	return b
}

func CanJump(steps []uint) bool {

	if len(steps) == 0 {
		return false
	}

	maxjump := 0
	for i := 0; i < len(steps); i++ {
		if i > maxjump {
			return false
		}

		maxjump = max(maxjump, i+int(steps[i]))

		if maxjump == len(steps)-1 || len(steps) == 1 {
			return true
		}
	}

	return false
}
