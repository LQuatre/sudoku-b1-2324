package main

import (
	"fmt"
	"os"
)

func isSafe(sudoku *[9][9]int, row, col, n int) bool {
	for i := 0; i < 9; i++ {
		if sudoku[row][i] == n || sudoku[i][col] == n {
			return false
		}
	}
	row -= row % 3
	col -= col % 3
	for i := row; i < row+3; i++ {
		for j := col; j < col+3; j++ {
			if sudoku[i][j] == n {
				return false
			}
		}
	}
	return true
}

func solveSudoku(sudoku *[9][9]int) bool {
	var row, col int
	var found bool
	for i, r := range sudoku {
		for j, n := range r {
			if n == 0 {
				row = i
				col = j
				found = true
				break
			}
		}
		if found {
			break
		}
	}
	if !found {
		return true
	}
	for n := 1; n <= 9; n++ {
		if isSafe(sudoku, row, col, n) {
			sudoku[row][col] = n
			if solveSudoku(sudoku) {
				return true
			}
			sudoku[row][col] = 0
		}
	}
	return false
}

func main() {
	args := os.Args[1:]
	if len(args) != 9 {
		fmt.Println("Error")
		return
	}
	var sudoku [9][9]int
	for i, arg := range args {
		if len(arg) != 9 {
			fmt.Println("Error")
			return
		}
		for j, r := range arg {
			if r == '.' {
				sudoku[i][j] = 0
			} else if r >= '1' && r <= '9' {
				sudoku[i][j] = int(r - '0')
			} else {
				fmt.Println("Error")
				return
			}
		}
	}
	if !solveSudoku(&sudoku) {
		fmt.Println("Error")
		return
	}
	countLine := 0
	for _, row := range sudoku {
		// ajouter _ toute les 3 cases vers le bas
		if countLine%3 == 0 {
			fmt.Printf("\033[34m|-----------------------|\033[0m\n")
		}
		for i, n := range row {
			if i%3 == 0 {
				fmt.Printf("\033[34m| \033[0m")
			}
			for j := 0; j < 9; j++ {
				for k := 0; k < 9; k++ {
					if args[j][k] != '.' && j == countLine && k == i {
						fmt.Printf("\033[33m%d\033[0m ", n)
						break
					} else if j == countLine && k == i {
						fmt.Printf("\033[32m%d\033[0m ", n)
						break
					}
				}
			}
			if i == 8 {
				fmt.Printf("\033[34m|\033[0m")
			}
		}
		fmt.Println()
		if countLine == 8 {
			// fmt.Println("|-----------------------|") orange
			fmt.Printf("\033[34m|-----------------------|\033[0m\n")
		}
		countLine++
	}
}
