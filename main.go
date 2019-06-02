/*
go-sudoku is a hobby package built to play around with algorithms that can
solve sudoku puzzles of various difficulties. The idea was to create a repository
of sudoku solving methods, and make them compete against each other for time to
completion of puzzles.

@author Jon Lim

https://github.com/JonLim/go-sudoku
*/
package main

import (
	"fmt"
	"time"
)

var solverList = map[string]func([][]int) [][]int{}
var puzzleList = map[string]sudokuPuzzle{}

func addSolver(solverType string, solver func([][]int) [][]int) {
	solverList[solverType] = solver
}

func init() {
	addSolver("backtrack", backtrackSolver)

	loadSudokuPuzzles()
}

func main() {
	for index := range solverList {
		solverFn := solverList[index]
		for puzzleName := range puzzleList {
			start := time.Now()

			_ = solverFn(puzzleList[puzzleName])

			// UNCOMMENT TO SEE SOLVED BOARDS
			// for row := range solvedBoard {
			// 	fmt.Println(solvedBoard[row])
			// }

			fmt.Printf("%s - Total time to solve: %s\n", puzzleName, time.Since(start))
		}
	}
}
