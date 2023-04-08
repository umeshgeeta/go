/*
 * https://leetcode.com/problems/remove-element/description/
 *
 * Given an integer array nums and an integer val, remove all occurrences of val in nums in-place.
 * The order of the elements may be changed. Then return the number of elements in nums which are not equal to val.
 *
 * Consider the number of elements in nums which are not equal to val be k, to get accepted,
 * you need to do the following things:
 * Change the array nums such that the first k elements of nums contain the elements which are not equal to val.
 * The remaining elements of nums are not important as well as the size of nums.
 * Return k.
 *
 * Copyright (c) 2020. Neosemantix, Inc.
 * Author: Umesh Patil
 */

package main

import (
	"fmt"
)

func main() {

	nums := []int{4, 5}
	fmt.Printf("%d\n", removeElement(nums, 5))
	fmt.Println(nums)

	nums = []int{1}
	fmt.Printf("%d\n", removeElement(nums, 1))
	fmt.Println(nums)

	nums = []int{}
	fmt.Printf("%d\n", removeElement(nums, 0))
	fmt.Println(nums)

	nums = []int{3, 3}
	fmt.Printf("%d\n", removeElement(nums, 5))
	fmt.Println(nums)

	nums = []int{3, 3}
	fmt.Printf("%d\n", removeElement(nums, 3))
	fmt.Println(nums)

	nums = []int{3, 2, 2, 3}
	fmt.Printf("%d\n", removeElement(nums, 2))
	fmt.Println(nums)

	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Printf("%d\n", removeElement(nums, 2))
	fmt.Println(nums)

	nums = []int{2}
	fmt.Printf("%d\n", removeElement(nums, 3))
	fmt.Println(nums)
}

func removeElement(nums []int, val int) int {
	size := len(nums)
	// added so that for an empty array we do not get answer 1
	if size == 0 {
		return 0
	}
	if size == 1 {
		if nums[0] == val {
			return 0
		}
		return 1
	}
	// size is 2 or more
	swapped := false
	i := 0
	j := size - 1
	for i < j {
		for nums[i] != val && i < j {
			i++
		}
		for nums[j] == val && i < j {
			j--
		}
		if i <= j {
			if nums[i] == val && nums[j] != val {
				nums[i] = nums[j]
				nums[j] = val
				swapped = true
			}
		}
	}
	if swapped {
		return i
	}
	if nums[i] == val {
		return i
	}
	return i + 1
}
