/*
Backtrack solver is one of the more straightforward ways of solving a sudoku
puzzle - effectively using brute force for recursively placing a number in a
square, checking the validity of the board after placing the number, and
repeating the process for every square on the board.

@author Jon Lim
*/
package main

// Conforms to solverList map of functions, requiring a slice of int slices,
// and returns a completed slice of int slices.
func backtrackSolver(board [][]int) [][]int {
	if backtrack(board) {
		return board
	}
	return board
}

// Function backtrack looks for spots on the board that are empty, AKA if it is
// a 0 in the slice, and tries to place a number from 1 to 9 in its place before
// validating the board, and trying another number if it is not valid.
//
// If the board is valid, backtrack goes to the next empty spot on the board and
// continues to check validity of the board with every change. If the board is
// invalid, the value is reset back to 0, or empty square.
func backtrack(board [][]int) bool {
	if !hasEmptyPlace(board) {
		return true
	}

	for rowIndex := range board {
		for colIndex, currentNum := range board[rowIndex] {
			if currentNum == 0 {
				for testNum := 1; testNum <= 9; testNum++ {
					board[rowIndex][colIndex] = testNum
					if isBoardValid(board) {
						if backtrack(board) {
							return true
						}
						board[rowIndex][colIndex] = 0
					}
					board[rowIndex][colIndex] = 0
				}

				return false
			}
		}
	}
	return false
}

func hasEmptyPlace(board [][]int) bool {
	for rowIndex := range board {
		for _, num := range board[rowIndex] {
			if num == 0 {
				return true
			}
		}
	}
	return false
}

// Used for validating rows, columns, and boxes. Requires some pre-processing
// by having consumer create the slice of integers that represent the row,
// column, or box, but is a fairly re-usable method!
func isNumSetValid(numSet []int) bool {
	numberBoard := map[int]bool{}
	for _, num := range numSet {
		// View 0 as blank space
		if num == 0 {
			continue
		}

		if _, ok := numberBoard[num]; ok {
			return false
		}
		numberBoard[num] = true
	}
	return true
}

// Function isBoardValid does the validation of each row, each column, and every
// 3x3 box on the board, using isNumSetValid function.
func isBoardValid(board [][]int) bool {
	// Validate Rows
	for index := range board {
		isRowValid := isNumSetValid(board[index])
		if !isRowValid {
			return false
		}
	}

	// Validate Columns
	for index := range board {
		colNums := []int{}
		for col := 0; col < 9; col++ {
			colNums = append(colNums, board[col][index]) // Creates slice of values from volumn
		}
		isColValid := isNumSetValid(colNums)
		if !isColValid {
			return false
		}
	}

	// Validate Boxes
	boxRows := [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	}
	for _, boxRow := range boxRows {
		// First 3 Cols
		boxNums := []int{}
		for col := 0; col < 3; col++ {
			boxNums = append(boxNums, board[boxRow[0]][col])
			boxNums = append(boxNums, board[boxRow[1]][col])
			boxNums = append(boxNums, board[boxRow[2]][col])
		}
		isBoxValid := isNumSetValid(boxNums)
		if !isBoxValid {
			return false
		}

		// Second 3 Cols
		boxNums = []int{}
		for col := 3; col < 6; col++ {
			boxNums = append(boxNums, board[boxRow[0]][col])
			boxNums = append(boxNums, board[boxRow[1]][col])
			boxNums = append(boxNums, board[boxRow[2]][col])
		}
		isBoxValid = isNumSetValid(boxNums)
		if !isBoxValid {
			return false
		}

		// Third 3 Cols
		boxNums = []int{}
		for col := 6; col < 9; col++ {
			boxNums = append(boxNums, board[boxRow[0]][col])
			boxNums = append(boxNums, board[boxRow[1]][col])
			boxNums = append(boxNums, board[boxRow[2]][col])
		}
		isBoxValid = isNumSetValid(boxNums)
		if !isBoxValid {
			return false
		}
	}

	return true
}
