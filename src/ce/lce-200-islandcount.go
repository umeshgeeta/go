/*
 * https://leetcode.com/problems/number-of-islands/description/
 *
 * Given an m x n 2D binary grid grid which represents a map of '1's (land) and
 * '0's (water), return the number of islands.
 *
 * An island is surrounded by water and is formed by connecting adjacent lands horizontally or
 * vertically. You may assume all four edges of the grid are all surrounded by water.
 *
 */

package main

import "fmt"

func numIslands(grid [][]byte) int {
	islandCount := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 || grid[i][j] == 49 { // 49 is int value for "1" which is the input
				islandCount++
				mark(grid, i, j)
			}
		}
	}
	return islandCount
}

func mark(grid [][]byte, i int, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) {
		return
	}
	if grid[i][j] == 0 || grid[i][j] == 48 { // 48 is int value for "0"
		return
	}
	grid[i][j] = 0
	mark(grid, i-1, j)
	mark(grid, i+1, j)
	mark(grid, i, j-1)
	mark(grid, i, j+1)
}

func main() {
	grid1 := [][]byte{
		[]byte{1, 1, 1, 1, 0},
		[]byte{1, 1, 0, 1, 0},
		[]byte{1, 1, 0, 0, 0},
		[]byte{0, 0, 0, 0, 0},
	}
	nc1 := numIslands(grid1)
	fmt.Printf("number of islands found: %d\n", nc1) // expected 1 island

	grid2 := [][]byte{
		[]byte{1, 1, 0, 0, 0},
		[]byte{1, 1, 0, 0, 0},
		[]byte{0, 0, 1, 0, 0},
		[]byte{0, 0, 0, 1, 1},
	}
	nc2 := numIslands(grid2)
	fmt.Printf("number of islands found: %d\n", nc2) // expected 3 island

}
