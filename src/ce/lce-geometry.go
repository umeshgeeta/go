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

// LeetCode problen no. 223: Rectangle Area
// https://leetcode.com/problems/rectangle-area/description/
//
// Given the coordinates of two rectilinear rectangles in a 2D plane,
// return the total area covered by the two rectangles.
//
// The first rectangle is defined by its bottom-left corner (ax1, ay1)
// and its top-right corner (ax2, ay2).
//
// The second rectangle is defined by its bottom-left corner (bx1, by1)
// and its top-right corner (bx2, by2).
//
// Constraints:
//
// -104 <= ax1 <= ax2 <= 104
// -104 <= ay1 <= ay2 <= 104
// -104 <= bx1 <= bx2 <= 104
// -104 <= by1 <= by2 <= 104
func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	// It is given that ax1 < ax2 and bx1 < bx2.
	var cx1, cx2, cy1, cy2 int
	var overlapX, overlapY bool
	if bx1 < ax1 && ax1 < bx2 && bx2 <= ax2 {
		overlapX = true
		cx1, cx2 = ax1, bx2
	} else if bx1 >= ax1 && bx2 <= ax2 {
		overlapX = true
		cx1, cx2 = bx1, bx2
	} else if ax1 <= bx1 && bx1 < ax2 && ax2 <= bx2 {
		overlapX = true
		cx1, cx2 = bx1, ax2
	} else if bx1 < ax1 && ax2 < bx2 {
		overlapX = true
		cx1, cx2 = ax1, ax2
	}
	// else no overlap; too much undershoot or overshoot
	if by1 < ay1 && ay1 < by2 && by2 <= ay2 {
		overlapY = true
		cy1, cy2 = ay1, by2
	} else if ay1 <= by1 && by2 <= ay2 {
		overlapY = true
		cy1, cy2 = by1, by2
	} else if ay1 <= by1 && by1 < ay2 && ay2 <= by2 {
		overlapY = true
		cy1, cy2 = by1, ay2
	} else if by1 < ay1 && ay2 < by2 {
		overlapY = true
		cy1, cy2 = ay1, ay2
	}
	overlapArea := 0
	if overlapX && overlapY {
		overlapArea = (cx2 - cx1) * (cy2 - cy1)
	}
	aArea := (ax2 - ax1) * (ay2 - ay1)
	bArea := (bx2 - bx1) * (by2 - by1)
	return aArea + bArea - overlapArea
}

func main() {

	fmt.Println(computeArea(-5, 0, 0, 3, -3, -3, 3, 3))

	fmt.Println(computeArea(-2, -2, 2, 2, 3, 3, 4, 4))

	fmt.Println(computeArea(-5, -3, 0, 0, -3, -3, 3, 3))

	fmt.Println(computeArea(-3, 0, 3, 4, 0, -1, 9, 2))

	fmt.Println(computeArea(-2, -2, 2, 2, -2, -2, 2, 2))

	//points := [][]int{{1, 1}, {0, 1}, {-1, 1}, {0, 0}}
	//fmt.Println(isReflected(points))
	//
	//points = [][]int{{10, 10}, {11, 11}, {9, 11}}
	//fmt.Println(isReflected(points))
	//
	//points = [][]int{{0, 0}, {0, 0}}
	//fmt.Println(isReflected(points))
	//
	//points = [][]int{{1, 1}, {-1, 1}}
	//fmt.Println(isReflected(points))
	//
	//points = [][]int{{1, 1}, {-1, -1}}
	//fmt.Println(isReflected(points))
	//
	//points = [][]int{{-16, 1}, {16, 1}, {16, 1}}
	//fmt.Println(isReflected(points))
}
