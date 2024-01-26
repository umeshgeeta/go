/*
You own a resturant, and you would like to know how many tables you need
Given an array of reservations, where a reservation is a start time and an end time
ex: [0, 20], determine how many tables you need

Example 1:
[[10,30],[0,20]] : 2 Tables

Example 2:
[[0, 10], [12,17], [19,22]] : 1 Table
*/

package main

import "fmt"


func main() {
	// 0 is start, 10 is end
  /* input := [][]int{{0,10}, {5,14}, {12,17}}
	fmt.Println(tableCount(input)) */

	// input := [][]int{{0, 10}, {12,17}, {19,22}}
	// fmt.Println(tableCount(input))
	input := [][]int{{0,10}, {2, 5}, {8,12}}
	fmt.Println(tableCount(input))

}


func tableCount(reservations [][]int) int {
	tableEndTimes := make([]int, 0)
	for _, timePair := range reservations {
		tableCount := len(tableEndTimes)
		if tableCount == 0 {
			// first table
			tableEndTimes = append(tableEndTimes, timePair[1])
		} else {
			findTable := false
			j := 0
			for j < tableCount && !findTable {
				crtEndTime := tableEndTimes[j]
				if timePair[0] > crtEndTime {
					// we can sit this reservation on this tableCount
					tableEndTimes[j] = timePair[1]
					findTable = true
				}
				j++
			}
			if !findTable {
				// we did not find a table, let us add
				tableEndTimes = append(tableEndTimes, timePair[1])
			}
		}
	}
	return len(tableEndTimes)
}