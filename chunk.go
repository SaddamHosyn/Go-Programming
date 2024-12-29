package main

import "fmt"

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}

//Write a function called Chunk that receives as parameters a slice, slice []int, and a number size int. The goal of this function is to chunk a slice into many sub slices where each sub slice has the length of size.

//If the size is 0 it should print a newline ('\n').

func Chunk(slice []int, size int) {
	if size == 0 {
		fmt.Println()
		return
	}

	result := [][]int{}

	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)   // This condition ensures that the program only includes elements within the bounds of the original slice and doesn’t attempt to retrieve nonexistent values.
		}
		result = append(result, slice[i:end])
	}

	fmt.Println(result)
}


