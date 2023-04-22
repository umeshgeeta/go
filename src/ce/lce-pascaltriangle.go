/*
 * Copyright (c) 2020. Neosemantix, Inc.
 * Author: Umesh Patil
 */

package main

import "fmt"

/*
 * LeetCode no. 119: Pascal's Triangle 2
 * https://leetcode.com/problems/pascals-triangle/description/
 *
 * Given an integer numRows, return the first numRows of Pascal's triangle.
 */
func generate(numRows int) [][]int {
	if numRows == 1 {
		return [][]int{{1}}
	}
	result := make([][]int, 0)
	rowIndex := 0
	for rowIndex < numRows {
		rowSize := rowIndex + 1
		row := make([]int, rowSize)
		row[0] = 1
		row[rowSize-1] = 1
		if rowSize > 2 {
			colIndex := 1
			for colIndex < rowSize-1 {
				row[colIndex] = result[rowIndex-1][colIndex-1] + result[rowIndex-1][colIndex]
				colIndex++
			}
		}
		result = append(result, row)
		rowIndex++
	}
	return result
}

/*
 * LeetCode problem no. 119 - Pascal Triangle 2
 * https://leetcode.com/problems/pascals-triangle-ii/description/
 *
 * Given an integer rowIndex, return the rowIndexth (0-indexed) row of the Pascal's triangle.
 *
 */
func getRow(rowIndex int) []int {
	pt := generate(rowIndex + 1)
	return pt[rowIndex]
}

func main() {
	fmt.Println(generate(1))
	fmt.Println(generate(2))
	fmt.Println(generate(3))
	fmt.Println(generate(4))
	fmt.Println(generate(5))
	fmt.Println(generate(6))
}
