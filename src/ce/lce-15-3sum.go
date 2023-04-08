/*
 * Copyright (c) 2020. Neosemantix, Inc.
 * Author: Umesh Patil
 */

package main

import "fmt"

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	result := threeSum(nums)
	fmt.Println(result)

	nums = []int{0, 1, 1}
	result = threeSum(nums)
	fmt.Println(result)

	nums = []int{0, 0, 0}
	result = threeSum(nums)
	fmt.Println(result)

	nums = []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}
	result = threeSum(nums)
	fmt.Println(result)
}

var soln map[int][]int
var tripletCount int

func orderAdd(first int, second int) {
	// always use 2 smallest
	addPair(smallestTwo(second, first, 0-(first+second)))
}

func smallestTwo(f int, s int, t int) (int, int) {
	smallest := f
	smaller := s
	if s < smallest {
		smallest = s
		smaller = f
	}
	if t < smallest {
		return t, smallest
	}
	if t < smaller {
		return smallest, t
	}
	return smallest, smaller
}

func addPair(first int, second int) {
	vals, ok := soln[first]
	if !ok {
		// it is not present
		vals = []int{}
	} else {
		for _, v := range vals {
			if v == second {
				// already present, no need to add
				return
			}
		}
		// it is not present
	}
	vals = append(vals, second)
	tripletCount++
	soln[first] = vals
}

func threeSum(nums []int) [][]int {
	// initialize
	soln = make(map[int][]int)
	tripletCount = 0

	size := len(nums)
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			sum := nums[i] + nums[j]
			for k := j + 1; k < size; k++ {
				if nums[k] == (-1 * sum) { // i != k && j != k &&
					orderAdd(nums[i], nums[j])
				}
			}
		}
	}
	result := make([][]int, 0)
	for first, fa := range soln {
		for _, second := range fa {
			result = append(result, []int{first, second, -1 * (first + second)})
		}
	}
	return result
}
