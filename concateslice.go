package main

import "fmt"

//Write a function ConcatSlice() that takes two slices of integers as arguments and returns the concatenation of the two slices.

// ConcatSlice takes two slices of integers and returns their concatenation.
func ConcatSlice(slice1, slice2 []int) []int {
	result := []int{} // Initialize an empty slice to hold the result
	// Append elements from the first slice
	for _, char := range slice1 {
		result = append(result, char)
	}
	// Append elements from the second slice
	for _, char := range slice2 {
		result = append(result, char)
	}
	return result // Return the concatenated result
}

func main() {
	// Testing the ConcatSlice function
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{4, 5, 6}))   // Output: [1 2 3 4 5 6]
	fmt.Println(ConcatSlice([]int{}, []int{4, 5, 6, 7, 8, 9})) // Output: [4 5 6 7 8 9]
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{}))          // Output: [1 2 3]
}
