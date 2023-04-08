/*
 * Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k,
 * and nums[i] + nums[j] + nums[k] == 0.
 *
 * Notice that the solution set must not contain duplicate triplets.
 *
 * https://leetcode.com/problems/3sum/description/
 *
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

	nums = []int{-1, 0, 1, 0}
	result = threeSum(nums)
	fmt.Println(result)
}

var soln map[int]map[int]bool
var tripletCount int
var addressed map[string]bool
var kmap map[int]map[int]bool

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
		vals = make(map[int]bool)
	} else {
		_, ok2 := vals[second]
		if ok2 {
			return
		}
		// it is not present
	}
	vals[second] = true
	tripletCount++
	soln[first] = vals
}

func buildKey(numsi int, numsj int) string {
	if numsi < numsj {
		return fmt.Sprintf("%d#%d", numsi, numsj)
	}
	return fmt.Sprintf("%d#%d", numsj, numsi)
}

func buildKMap(nums []int) {
	kmap = make(map[int]map[int]bool)
	n := len(nums)
	for k := 0; k < n; k++ {
		numk := nums[k]
		m, ok := kmap[numk]
		if !ok {
			m = make(map[int]bool)
		}
		m[k] = true
		kmap[numk] = m
	}
}

func isValFromOtherIndex(val int, indexi int, indexj int) bool {
	m, present := kmap[val]
	if present {
		// value is present
		// next we have to check that it also comes from some other index other than the passed ones
		someOneElseSupplies := false
		for indexKey, _ := range m {
			if indexKey != indexi && indexKey != indexj {
				someOneElseSupplies = true
				break
			}
		}
		present = someOneElseSupplies
	}
	// the value itself is not availabe
	return present
}

func isAddressed(numsi int, numsj int) (string, bool) {
	key := buildKey(numsi, numsj)
	_, ok := addressed[key]
	return key, ok
}

func threeSum(nums []int) [][]int {
	// initialize
	soln = make(map[int]map[int]bool)
	tripletCount = 0
	addressed = make(map[string]bool)
	size := len(nums)
	buildKMap(nums)

	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			key, present := isAddressed(nums[i], nums[j])
			//fmt.Printf("(%d) %d (%d) %d %s\n", i, nums[i], j, nums[j], key)
			if !present {
				sum := nums[i] + nums[j]
				val := 0 - sum
				if isValFromOtherIndex(val, i, j) {
					orderAdd(nums[i], nums[j])

					addressed[key] = true
					addressed[buildKey(nums[i], val)] = true
					addressed[buildKey(val, nums[j])] = true
				}
			}
		}
	}
	result := make([][]int, 0)
	for first, fa := range soln {
		for second, _ := range fa {
			result = append(result, []int{first, second, -1 * (first + second)})
		}
	}
	return result
}
