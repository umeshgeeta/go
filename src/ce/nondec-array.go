// MIT License
// Author: Umesh Patil, Neosemantix, Inc.

package main

import "fmt"

//
// Given an array nums with n integers, your task is to check if it could become
// non-decreasing by modifying at most 1 element.
//
// We define an array is non-decreasing if nums[i] <= nums[i + 1] holds for
// every i (0-based) such that (0 <= i <= n - 2).
//
// https://leetcode.com/problems/non-decreasing-array/

var mm map[int]int

// This is bit efficient implementation, we do only 1 pass of the array.
// We take a current number and compare it with it's next value. If the next
// value is smaller than the current value, we count 1 modification is needed.
// The real challenge is to determine which value is to be changed. Our
// preference is to change the current value. But you will not be able to
// reduce the current value to the level of next number, if your previous value
// is holding the barrier. In other words, you can reduce the current value
// provided previous value is low enough. That is when you need to do the
// comparison between next value and previous value to determine if the current
// value can be reduced. If not, then next value needs to be changed and more
// importantly that value needs to be used in subsequent comparison. That is
// when we use map. We can alter the given array but that is a bad practices and
// possible loss of original values can create an issue (I have not thought
// through); but in any case generally input should not altered permanently.
// To track this 'alternative input' we use map instead of allocating duplicate
// array. There will not be that many entries in the map since after two
// modifications we exit the loop and only when there are modifications, there
// are relevant entries in the map.
func checkPossibility(nums []int) bool {
	len := len(nums)
	if len < 3 {
		return true
	} else {
		mm = make(map[int]int)
		mod := 0
		i := 0
		setVal(1, nums[1])
		for i < len-1 && mod < 2 {
			newVal := getVal(nums, i+1)
			if newVal < getVal(nums, i) {
				mod++
				if i == 0 {
					// we can change nums[0], newVal remains as is
				} else {
					if newVal < getVal(nums, i-1) {
						setVal(i+1, getVal(nums, i))
					} else {
						// we can change nums[i]
						setVal(i, getVal(nums, i-1))
					}
				}
			}
			// else it is non decreasing, continue forward
			i++
		}
		if mod < 2 {
			return true
		} else {
			return false
		}
	}
}

func setVal(i, v int) {
	mm[i] = v
}

func getVal(nums []int, i int) int {
	v, ok := mm[i]
	if ok {
		return v
	} else {
		return nums[i]
	}
}

func main() {
	nums := []int{2, 3, 3, 2, 4}
	fmt.Printf("%v\n", checkPossibility(nums))
	nums = []int{3, 4, 2, 3}
	fmt.Printf("%v\n", checkPossibility(nums))
	nums = []int{4, 2, 3}
	fmt.Printf("%v\n", checkPossibility(nums))
	nums = []int{3, 3, 2, 2}
	fmt.Printf("%v\n", checkPossibility(nums))
}
