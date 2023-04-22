// MIT License
// Author: Umesh Patil, Neosemantix, Inc.

package main

import (
	"fmt"
	"math"
)

// Given n non-negative integers representing an elevation map where the width
// of each bar is 1, compute how much water it is able to trap after raining.
//
// Example
// Input: [0,1,0,2,1,0,1,3,2,1,2,1]
// Output: 6
//
// https://leetcode.com/explore/interview/card/amazon/76/array-and-strings/2975/
//
// Basic idea is if you hypothetically assume max height value in the first and
// last cell, we get the 'union' volume of the water. We should find the
// intersection since in reality both ends do not have max height walls.
func trap(height []int) int {
	len := len(height)
	leftMax := make([]int, len)
	rightMax := make([]int, len)
	lx := 0
	rx := 0
	for i, v := range height {
		if v > lx {
			lx = v
		}
		leftMax[i] = lx
		j := len - (i + 1)
		w := height[j]
		if w > rx {
			rx = w
		}
		rightMax[j] = rx
	}
	ans := 0
	for k, h := range height {
		min := math.Min(float64(leftMax[k]), float64(rightMax[k]))
		ans = (int(min) - h) + ans
	}
	return ans
}

func main() {
	input := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Printf("%d\n", trap(input))
}
