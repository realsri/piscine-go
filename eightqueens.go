// package main
package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	var board [8]int
	row(board, 0)
}

/*
func main() {
	var board [8]int
	row(board, 0)
}*/

func col_diag(board [8]int, r int, col int) bool {
	for i := 0; i < r; i++ {
		// column check
		if board[i] == col {
			return false
		}

		// diagonal check
		if abs(board[i]-col) == abs(i-r) {
			return false
		}
	}
	return true
}

func row(board [8]int, r int) {
	if r == 8 {
		printBoard(board)
		return
	}

	for col := 0; col < 8; col++ {
		if col_diag(board, r, col) {
			board[r] = col
			row(board, r+1)
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func printBoard(board [8]int) {
	for i := 0; i < 8; i++ {
		z01.PrintRune(rune(board[i] + '1'))
	}
	z01.PrintRune('\n')
}
