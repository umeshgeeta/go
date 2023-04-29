/*
 * MIT License
 * Copyright (c) 2023. Neosemantix, Inc.
 * Author: Umesh Patil
 */

package main

import (
	"fmt"
	"math"
)

/**
 *
 * LeetCode problem no. 1011:  Capacity To Ship Packages Within D Days
 * https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/description/
 *
 * A conveyor belt has packages that must be shipped from one port to another within days days.
 * The ith package on the conveyor belt has a weight of weights[i]. Each day, we load the ship
 * with packages on the conveyor belt (in the order given by weights). We may not load more weight
 * than the maximum weight capacity of the ship.
 *
 * Return the least weight capacity of the ship that will result in all the packages on the
 * conveyor belt being shipped within days days.
 *
 * Constraints given:
 *
 * 1 <= days <= weights.length <= 5 * 104
 * 1 <= weights[i] <= 500
 *
 */
func shipWithinDays(weights []int, days int) int {
	minDiff := math.MaxInt
	minWeight := math.MaxInt
	highestWeight := 0
	sum := 0
	diffList := make([]int, len(weights))
	for i, w := range weights {
		if i > 0 {
			d := absint(weights[i] - weights[i-1])
			if d != 0 && d < minDiff {
				minDiff = d
			}
			diffList[i] = d
		}
		if w > highestWeight {
			highestWeight = w
		}
		if w < minWeight {
			minWeight = w
		}
		sum += w
	}
	diffList[0] = minWeight
	if minDiff == math.MaxInt {
		// all are equal
		minDiff = minWeight
	} else {
		minDiff = gcdList(diffList)
	}
	capacity := highestWeight
	found := false
	for !found {
		inhd := inHowManyDays(weights, capacity)
		if inhd <= days {
			found = true
		} else {
			capacity += minDiff
		}
	}
	return capacity
}

func inHowManyDays(weights []int, capacity int) int {
	days := 0
	currentSum := 0
	for _, w := range weights {
		if (currentSum + w) <= capacity {
			currentSum += w
		} else {
			days++
			currentSum = w
		}
	}
	if currentSum > 0 {
		days++
	}
	return days
}

func absint(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// gcd calculates the greatest common divisor using the Euclidean algorithm.
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// gcdList calculates the greatest common divisor of a list of integers.
func gcdList(nums []int) int {
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		result = gcd(result, nums[i])
	}
	return result
}

func main() {

	weights := []int{147, 73, 265, 305, 191, 152, 192, 293, 309, 292, 182, 157, 381, 287, 73, 162, 313, 366, 346, 47}
	fmt.Println(shipWithinDays(weights, 10))

	//weights := []int{10, 50, 100, 100, 50, 100, 100, 100}
	//fmt.Println(shipWithinDays(weights, 5))

	//weights := []int{3, 2, 2, 4, 1, 4}
	//fmt.Println(shipWithinDays(weights, 3))
	//weights = []int{1, 2, 3, 1, 1}
	//fmt.Println(shipWithinDays(weights, 4))
}
