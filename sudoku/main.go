package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]

	if len(arguments) != 9 {
		fmt.Println("Error") // number of arguments is not 9
		return
	}

	for _, arg := range arguments {
		if len(arg) != 9 {
			fmt.Println("Error") // length of arg is not 9
			return
		}
	}

	if hasInvalidCharacters(arguments) {
		fmt.Println("Error") // Input has invalid characters
		return
	}

	if hasDuplicates(arguments) {
		fmt.Println("Error")
		return
	}

	table := [9][9]rune{}
	table = fillTable(table, arguments)
	if isSolve(&table) {
		for y := 0; y < 9; y++ {
			for x := 0; x < 9; x++ {
				if x != 8 {
					z01.PrintRune(rune(table[y][x]))
					z01.PrintRune(' ')
				} else {
					z01.PrintRune(rune(table[y][x]))
				}
			}
			z01.PrintRune(10)
		}
	} else {
		fmt.Println("Error")
	}
}

func hasInvalidCharacters(args []string) bool {
	for _, arg := range args {
		for _, value := range arg {
			if value < '1' || value > '9' {
				if value != '.' {
					return true
				}
			}
		}
	}
	return false
}

// Check for duplicate values in rows, columns, and 3x3 sub-grids
func hasDuplicates(args []string) bool {
	board := [9][9]rune{}
	board = fillTable(board, args)

	// Check rows and columns
	for i := 0; i < 9; i++ {
		row := make(map[rune]bool)
		col := make(map[rune]bool)
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				if row[board[i][j]] {
					return true
				}
				row[board[i][j]] = true
			}
			if board[j][i] != '.' {
				if col[board[j][i]] {
					return true
				}
				col[board[j][i]] = true
			}
		}
	}

	// Check 3x3 sub-grids
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			subGridSet := make(map[rune]bool)
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					val := board[i+k][j+l]
					if val != '.' {
						if subGridSet[val] {
							return true
						}
						subGridSet[val] = true
					}
				}
			}
		}
	}

	return false
}

// fills with input arguments
func fillTable(table [9][9]rune, args []string) [9][9]rune {
	for i := range args {
		for j := range args[i] {
			table[i][j] = rune(args[i][j])
		}
	}
	return table
}

// counts remaining empty cells
func isDots(table *[9][9]rune) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if table[i][j] == '.' {
				return true
			}
		}
	}
	return false
}

// check if it fits
func isValid(table *[9][9]rune, x int, y int, z rune) bool {
	// check double int
	for i := 0; i < 9; i++ {
		if z == table[i][x] {
			return false
		}
	}
	for j := 0; j < 9; j++ {
		if z == table[y][j] {
			return false
		}
	}
	// box check
	a := x / 3
	b := y / 3
	for k := 3 * a; k < 3*(a+1); k++ {
		for l := 3 * b; l < 3*(b+1); l++ {
			if z == table[l][k] {
				return false
			}
		}
	}
	return true
}

// backtracking solver
func isSolve(table *[9][9]rune) bool {
	if !isDots(table) {
		return true
	}
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if table[y][x] == '.' {
				for z := '1'; z <= '9'; z++ {
					if isValid(table, x, y, z) {
						table[y][x] = z
						if isSolve(table) {
							return true
						}
					}
					table[y][x] = '.'
				}
				return false
			}
		}
	}
	return false
}
