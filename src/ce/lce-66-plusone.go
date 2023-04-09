/*
 * https://leetcode.com/problems/plus-one/description/
 *
 * You are given a large integer represented as an integer array digits, where each digits[i] is the ith digit of the
 * integer. The digits are ordered from most significant to least significant in left-to-right order. The large integer
 * does not contain any leading 0's.
 *
 * Increment the large integer by one and return the resulting array of digits.
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
	fmt.Println(plusOne(nums))

	nums = []int{1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(plusOne(nums))

	nums = []int{9}
	fmt.Println(plusOne(nums))
}

func plusOne(digits []int) []int {
	size := len(digits)
	if size == 0 {
		// if empty, return empty for now
		return digits
	}
	result := make([]int, size+1)
	i := size - 1
	j := size
	carry := 1
	for i > -1 {
		val := digits[i] + carry
		if val < 10 {
			result[j] = val
			carry = 0
		} else {
			result[j] = val % 10
			carry = 1
		}
		j--
		i--
	}
	result[0] = carry
	if result[0] == 0 {
		return result[1:]
	}
	return result
}
