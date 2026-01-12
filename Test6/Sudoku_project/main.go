package main

import (
	"fmt"
	"os"
)

func isValid(grid [9][9]int, row, col, num int) bool {
	// Check row
	for c := 0; c < 9; c++ {
		if grid[row][c] == num {
			return false
		}
	}

	// Check column
	for r := 0; r < 9; r++ {
		if grid[r][col] == num {
			return false
		}
	}

	// Check 3x3 box
	startRow := (row / 3) * 3
	startCol := (col / 3) * 3
	for r := startRow; r < startRow+3; r++ {
		for c := startCol; c < startCol+3; c++ {
			if grid[r][c] == num {
				return false
			}
		}
	}

	return true
}

func solve(grid *[9][9]int) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if grid[row][col] == 0 {
				for num := 1; num <= 9; num++ {
					if isValid(*grid, row, col, num) {
						grid[row][col] = num
						if solve(grid) {
							return true
						}
						grid[row][col] = 0
					}
				}
				return false
			}
		}
	}
	return true
}

func main() {
	if len(os.Args) != 10 {
		fmt.Println("Error")
		return
	}

	var grid [9][9]int

	for i := 0; i < 9; i++ {
		row := os.Args[i+1]
		if len(row) != 9 {
			fmt.Println("Error")
			return
		}

		for j := 0; j < 9; j++ {
			ch := row[j]
			if ch == '.' {
				grid[i][j] = 0
			} else if ch >= '1' && ch <= '9' {
				num := int(ch - '0')
				if !isValid(grid, i, j, num) {
					fmt.Println("Error")
					return
				}
				grid[i][j] = num
			} else {
				fmt.Println("Error")
				return
			}
		}
	}

	if !solve(&grid) {
		fmt.Println("Error")
		return
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(grid[i][j])
			if j < 8 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
