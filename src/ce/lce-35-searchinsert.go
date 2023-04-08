/*
 * https://leetcode.com/problems/search-insert-position/description/
 *
 * Given a sorted array of distinct integers and a target value, return the index if the target is found.
 * If not, return the index where it would be if it were inserted in order.
 * You must write an algorithm with O(log n) runtime complexity.
 *
 * This is standard Binary Search, recursive implementation.
 *
 * Copyright (c) 2020. Neosemantix, Inc.
 * Author: Umesh Patil
 */

package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, 5, 6}
	fmt.Printf("%d\n", searchInsert(nums, 5))
	fmt.Println("done")

	fmt.Printf("%d\n", searchInsert(nums, 2))
	fmt.Println("done")

	fmt.Printf("%d\n", searchInsert(nums, 7))
	fmt.Println("done")

	fmt.Printf("%d\n", searchInsert(nums, 0))
	fmt.Println("done")

	nums = []int{1}
	fmt.Printf("%d\n", searchInsert(nums, 1))
	fmt.Println("done")
}

func searchInsert(nums []int, target int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}
	if size == 1 {
		if target <= nums[0] {
			return 0
		}
		return 1
	}
	return search(nums, target, 0, size-1)
}

func search(nums []int, target int, start int, end int) int {
	//fmt.Printf("%d %d \n", start, end)
	if start == end {
		if target <= nums[start] {
			return start
		} else {
			return start + 1
		}
	}
	if end == start+1 {
		if target <= nums[start] {
			return start
		} else if target <= nums[end] {
			return end
		} else {
			return end + 1
		}
	}
	middle := (start + end) / 2
	//fmt.Printf("middle: %d\n", middle)
	if target <= nums[middle] {
		return search(nums, target, start, middle)
	} else {
		return search(nums, target, middle, end)
	}
}
