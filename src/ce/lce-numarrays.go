/*
 * Copyright (c) 2020. Neosemantix, Inc.
 * Author: Umesh Patil
 */

package main

import "fmt"

/*
 * LeetCode problem no. 136: Single Number
 * https://leetcode.com/problems/single-number/description/
 *
 * Given a non-empty array of integers nums, every element appears twice except for one.
 * Find that single one.
 *
 * You must implement a solution with a linear runtime complexity and use only constant extra space.
 *
 * In the implementation we use map which can grow at the most to hald of the input array length.
 */
func singleNumber(nums []int) int {
	ln := len(nums)
	if ln == 1 {
		return nums[0]
	}
	single := make(map[int]bool)
	for _, n := range nums {
		_, ok := single[n]
		if ok {
			delete(single, n)
		} else {
			single[n] = true
		}
	}
	for n, _ := range single {
		return n
	}
	return 0 // error condition
}

/*
 * LeetCode problem no. 163: Missing ranges
 * https://leetcode.com/problems/missing-ranges/description/
 *
 * You are given an inclusive range [lower, upper] and a sorted unique integer array nums,
 * where all elements are within the inclusive range.
 *
 * A number x is considered missing if x is in the range [lower, upper] and x is not in nums.
 *
 * Return the shortest sorted list of ranges that exactly covers all the missing numbers.
 * That is, no element of nums is included in any of the ranges, and each missing number
 * is covered by one of the ranges.
 *
 */
func findMissingRanges(nums []int, lower int, upper int) [][]int {
	result := make([][]int, 0)
	numsCount := len(nums)
	switch numsCount {
	case 0:
		result = addRange(result, lower, upper)
	case 1:
		if lower <= nums[0]-1 {
			result = addRange(result, lower, nums[0]-1)
		}
		if upper >= nums[0]+1 {
			result = addRange(result, nums[0]+1, upper)
		}
	default:
		i := 0
		if nums[0] > lower {
			result = addRange(result, lower, nums[0]-1)
		}
		for i < numsCount-1 {
			if nums[i+1] > nums[i]+1 {
				// add range
				result = addRange(result, nums[i]+1, nums[i+1]-1)
			}
			i++
		}
		if nums[numsCount-1] < upper {
			result = addRange(result, nums[numsCount-1]+1, upper)
		}
	}
	return result
}

func addRange(result [][]int, start int, end int) [][]int {
	rng := make([]int, 2)
	rng[0] = start
	rng[1] = end
	return append(result, rng)
}

func main() {
	nums := []int{-1}
	result := findMissingRanges(nums, -2, -1)
	fmt.Println(result)

	//nums = []int{0, 1, 3, 50, 75}
	//result = findMissingRanges(nums, 0, 99)
	//fmt.Println(result)
}
