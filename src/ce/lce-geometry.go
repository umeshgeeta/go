// MIT License
// Author: Umesh Patil, Neosemantix, Inc.

package main

import (
	"fmt"
	"sort"
)

// LeetCode problem n0. 356: Line Reflection
// https://leetcode.com/problems/line-reflection/description/
//
// Given n points on a 2D plane, find if there is such a line parallel to the y-axis
// that reflects the given points symmetrically.
//
// In other words, answer whether or not if there exists a line that after reflecting
// all points over the given line, the original points' set is the same as the reflected ones.
//
// Note that there can be repeated points.
//
// Constraints:
//
// n == points.length
// 1 <= n <= 104
// -108 <= points[i][j] <= 108
func isReflected(points [][]int) bool {
	if len(points) == 1 {
		// it is trivially true
		return true
	}
	pointMap := make(map[int][]int)
	var sampleY, oneY int
	refEncounter := false
	for _, pair := range points {
		xs, present := pointMap[pair[1]]
		if !present {
			xvals := []int{pair[0]}
			pointMap[pair[1]] = xvals
			oneY = pair[1]
		} else {
			xs = append(xs, pair[0])
			pointMap[pair[1]] = xs
			refEncounter = true
			sampleY = pair[1]
		}
	}
	if refEncounter {
		sampleXs := pointMap[sampleY]
		// this array of Xs must contain more than 1 Xs
		reflection, doubleMid := checkXvals(sampleXs)
		if !reflection {
			return false
		}
		for _, xvs := range pointMap {
			rfl, dmid := checkXvals(xvs)
			if !rfl {
				return false
			}
			if dmid != doubleMid {
				return false
			}
		}
	} else {
		// we have found for all Y values, there is only one X
		// then only acceptable scenario is when all of those
		// X values are same and in which case X = same value
		// is the reflection axis
		xvals := pointMap[oneY]
		xconstval := xvals[0]
		// we need to see this value for all Ys
		for _, xs := range pointMap {
			if len(xs) != 1 || xs[0] != xconstval {
				return false
			}
		}
	}
	return true
}

func curateXs(xs []int) []int {
	sort.Ints(xs)
	result := make([]int, 0)
	result = append(result, xs[0])
	n := len(xs)
	for i := 1; i < n; i++ {
		if xs[i] != xs[i-1] {
			result = append(result, xs[i])
		}
	}
	return result
}

func checkXvals(xvals []int) (bool, int) {
	xs := curateXs(xvals)
	n := len(xs)
	if n == 1 {
		return true, 2 * xs[0]
	}
	doubleMid := xs[n-1] + xs[0]
	midIndex := n / 2
	for i := 0; i < midIndex; i++ {
		if (xs[n-1-i] + xs[i]) != doubleMid {
			return false, 0
		}
	}
	// there can be at the most only 1 point on x == doubleMid/2,
	// the middle one
	if n%2 == 1 {
		if xs[midIndex] != (2 * doubleMid) {
			return false, 0
		}
	}
	return true, doubleMid
}

func main() {

	points := [][]int{{1, 1}, {0, 1}, {-1, 1}, {0, 0}}
	fmt.Println(isReflected(points))

	points = [][]int{{10, 10}, {11, 11}, {9, 11}}
	fmt.Println(isReflected(points))

	points = [][]int{{0, 0}, {0, 0}}
	fmt.Println(isReflected(points))

	points = [][]int{{1, 1}, {-1, 1}}
	fmt.Println(isReflected(points))

	points = [][]int{{1, 1}, {-1, -1}}
	fmt.Println(isReflected(points))

	points = [][]int{{-16, 1}, {16, 1}, {16, 1}}
	fmt.Println(isReflected(points))
}
