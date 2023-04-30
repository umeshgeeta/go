// MIT License
// Author: Umesh Patil, Neosemantix, Inc.

package main

import "fmt"

const (
	stateFixed    = 1
	stateIncrease = 2
	stateDecrease = 3
)

// LeetCode problem no. 54: Spiral Matrix
// https://leetcode.com/problems/spiral-matrix/description/
//
// Given an m x n matrix, return all elements of the matrix in spiral order.
//
// Constraints:
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 10
// -100 <= matrix[i][j] <= 100
func spiralOrder(matrix [][]int) []int {
	rowCount := len(matrix)
	colCount := len(matrix[0])
	visited := make([][]bool, rowCount)
	for i := range visited {
		visited[i] = make([]bool, colCount)
	}
	visitedSofar := 0
	expectedVisits := rowCount * colCount
	i, j := 0, -1 // initial value
	output := make([]int, expectedVisits)
	istate := stateFixed
	jstate := stateIncrease
	for visitedSofar < expectedVisits {
		previousi, previousj := i, j
		i, j = nextIndices(i, j, istate, jstate)
		if i < 0 || j < 0 || i >= rowCount || j >= colCount {
			// we have hit the border
			istate, jstate = nextState(istate, jstate)
			i, j = previousi, previousj
		} else {
			if !visited[i][j] {
				visited[i][j] = true
				output[visitedSofar] = matrix[i][j]
				visitedSofar++
			} else {
				// we have hit one of the shrinking borders
				istate, jstate = nextState(istate, jstate)
				i, j = previousi, previousj
			}
		}
	}
	return output
}

func nextIndices(i, j, istate, jstate int) (int, int) {
	switch istate {
	case stateFixed:
		switch jstate {
		case stateIncrease:
			return i, j + 1
		case stateDecrease:
			return i, j - 1
		}
	case stateIncrease:
		return i + 1, j //j must be fixed
	case stateDecrease:
		return i - 1, j
	}
	return -1, -1
}

func nextState(istate, jstate int) (int, int) {
	var newiState, newjState int
	switch istate {
	case stateFixed:
		// swap
		newiState, newjState = jstate, istate

	case stateIncrease:
		newiState = stateFixed
		newjState = stateDecrease

	case stateDecrease:
		newiState = stateFixed
		newjState = stateIncrease

	}
	return newiState, newjState
}

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(spiralOrder(matrix))

	matrix = [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(spiralOrder(matrix))
}
