/*
 * https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/
 *
 * Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique
 * element appears only once. The relative order of the elements should be kept the same.
 * Then return the number of unique elements in nums.
 *
 * Consider the number of unique elements of nums be k, to get accepted, you need to do the following things:
 * Change the array nums such that the first k elements of nums contain the unique elements in the order they were
 * present in nums initially. The remaining elements of nums are not important as well as the size of nums. Return k.
 *
 * Copyright (c) 2020. Neosemantix, Inc.
 * Author: Umesh Patil
 */

package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 1, 2}
	fmt.Printf("%d\n", removeDuplicates(nums))
	fmt.Println(nums)

	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Printf("%d\n", removeDuplicates(nums))
	fmt.Println(nums)
}

func removeDuplicates(nums []int) int {
	size := len(nums)
	if size < 2 {
		// for an empty array or an array with a single element...
		return size
	}
	here := 0
	lookout := 0
	previousVal := nums[0] - 1
	for lookout < size {
		if nums[lookout] == previousVal {
			lookout++
		} else {
			if here < lookout {
				nums[here] = nums[lookout]
			}
			previousVal = nums[here]
			lookout++
			here++
		}
	}
	return here
}
