package piscine

import "fmt"

const N = 9

// IsValid checks if it's valid to place num in board[row][col]
func IsValid(board [N][N]int, row, col, num int) bool {
	for x := 0; x < N; x++ {
		if board[row][x] == num || board[x][col] == num || board[row-row%3+x/3][col-col%3+x%3] == num {
			return false
		}
	}
	return true
}

// SolveSudoku uses backtracking to solve the Sudoku puzzle
func SolveSudoku(board [N][N]int) bool {
	row, col, found := FindEmpty(board)
	if !found {
		return true
	}
	for num := 1; num <= N; num++ {
		if IsValid(board, row, col, num) {
			board[row][col] = num
			if SolveSudoku(board) {
				return true
			}
			board[row][col] = 0
		}
	}
	return false
}

// FindEmpty finds an empty cell in the board (returns row, col, found)
func FindEmpty(board [N][N]int) (int, int, bool) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if board[i][j] == 0 {
				return i, j, true
			}
		}
	}
	return -1, -1, false
}

// PrintBoard prints the Sudoku board
func PrintBoard(board [N][N]int) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf("%d ", board[i][j])
		}
		fmt.Println()
	}
}
