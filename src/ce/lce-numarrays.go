/*
 * Copyright (c) 2020. Neosemantix, Inc.
 * Author: Umesh Patil
 */

package main

import (
	"fmt"
	"github.com/umeshgeeta/goshared/util"
	"math"

	"sort"
)

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

/*
 * LeetCode problem no. 18: 4Sum
 * https://leetcode.com/problems/4sum/description/
 *
 * Given an array nums of n integers, return an array of all the unique quadruplets [nums[a], nums[b], nums[c], nums[d]]
 * such that:
 *
 * - 0 <= a, b, c, d < n
 * - a, b, c, and d are distinct.
 * - nums[a] + nums[b] + nums[c] + nums[d] == target
 *
 * You may return the answer in any order.
 *
 */
func fourSum(nums []int, target int) [][]int {
	var result [][]int
	n := len(nums)
	if n < 4 {
		return result
	}
	sort.Ints(nums)
	if nums[0]+nums[1]+nums[2]+nums[3] > target {
		return result
	}
	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left, right := j+1, n-1
			for left < right {
				//fmt.Printf("i: %d j: %d left: %d right: %d\n", i, j, left, right)
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return result
}

/**
 * LeetCode problem no. 31: Next Permutation
 * https://leetcode.com/problems/next-permutation/description/
 *
 * A permutation of an array of integers is an arrangement of its members into a sequence or linear order.
 * For example, for arr = [1,2,3], the following are all the permutations of arr:
 * [1,2,3], [1,3,2], [2, 1, 3], [2, 3, 1], [3,1,2], [3,2,1].
 * The next permutation of an array of integers is the next lexicographically greater permutation of its integer.
 * More formally, if all the permutations of the array are sorted in one container according to their lexicographical
 * order, then the next permutation of that array is the permutation that follows it in the sorted container.
 * If such arrangement is not possible, the array must be rearranged as the lowest possible order
 * (i.e., sorted in ascending order).
 *
 * For example, the next permutation of arr = [1,2,3] is [1,3,2].
 * Similarly, the next permutation of arr = [2,3,1] is [3,1,2].
 * While the next permutation of arr = [3,2,1] is [1,2,3] because [3,2,1] does not have a lexicographical larger
 * rearrangement.
 *
 * Given an array of integers nums, find the next permutation of nums.
 *
 * The replacement must be in place and use only constant extra memory.
 *
 * Constraints:
 * 					1 <= nums.length <= 100
 * 					0 <= nums[i] <= 100
 */
func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2
	found := false
	for i > -1 && !found {
		if nums[i] < nums[i+1] {
			found = true
		} else {
			i--
		}
	}
	if i == -1 && !found {
		// no more permutation is possible
		reverseList(nums, 0, n-1)
		return
	}
	msIndex := findImmediateNext(nums, i)
	nums[i], nums[msIndex] = nums[msIndex], nums[i]
	sort.Ints(nums[i+1:])
}

// both start and end are indices
func reverseList(nums []int, start, end int) {
	midPoint := (end - start + 1) / 2
	for i := start; i < start+midPoint; i++ {
		nums[i], nums[end-i+start] = nums[end-i+start], nums[i]
	}
}

func findImmediateNext(nums []int, where int) int {
	n := len(nums)
	which := nums[where]
	result := math.MaxInt
	index := where
	for i := where + 1; i < n; i++ {
		if nums[i] > which && nums[i] < result {
			result, index = nums[i], i
		}
	}
	return index
}

type pair struct {
	startIndex int
	// map of fruit id and accumulated fruits so far
	fruits map[int]int
}

func newPair(start int) *pair {
	p := pair{
		startIndex: start,
		fruits:     make(map[int]int),
	}
	return &p
}

// LeetCode problem no. 904: Fruit Into Baskets
// https://leetcode.com/problems/fruit-into-baskets/description/
// You are visiting a farm that has a single row of fruit trees arranged from left to right.
// The trees are represented by an integer array fruits where fruits[i] is the type of fruit
// the ith tree produces.
//
// You want to collect as much fruit as possible. However, the owner has some strict rules
// that you must follow:
//
// You only have two baskets, and each basket can only hold a single type of fruit. There is no
// limit on the amount of fruit each basket can hold.
// Starting from any tree of your choice, you must pick exactly one fruit from every tree
// (including the start tree) while moving to the right. The picked fruits must fit in one of
// your baskets.
// Once you reach a tree with fruit that cannot fit in your baskets, you must stop.
// Given the integer array fruits, return the maximum number of fruits you can pick.
func totalFruit(fruits []int) int {
	pairs := make([]pair, 0)
	if len(fruits) > 0 {
		crtPair := newPair(0)
		crtPair.fruits[fruits[0]] = 1
		for i := 1; i < len(fruits); i++ {
			fruit := fruits[i]
			ct, present := crtPair.fruits[fruit]
			if present {
				crtPair.fruits[fruit] = ct + 1
			} else {
				if len(crtPair.fruits) == 1 {
					crtPair.fruits[fruit] = 1
				} else {
					// we have encountered a new fruit, start the new pair
					j := i - 1
					previousFruit := fruits[j]
					previousFruitCount := 0
					for j > -1 && fruits[j] == previousFruit {
						previousFruitCount++
						j--
					}
					// add the current pair
					pairs = append(pairs, *crtPair)
					// create a new pair
					crtPair = newPair(i)
					crtPair.fruits[previousFruit] = previousFruitCount
					crtPair.fruits[fruit] = 1
				}
			}
		}
		pairs = append(pairs, *crtPair)
	}
	return findMaxCount(pairs)
}

func findMaxCount(pairs []pair) int {
	maxCount := 0
	for _, p := range pairs {
		count := 0
		for _, c := range p.fruits {
			count += c
		}
		if maxCount < count {
			maxCount = count
		}
	}
	return maxCount
}

// LeetCode problem no. 907:  Sum of Subarray Minimums
// https://leetcode.com/problems/sum-of-subarray-minimums/description/
//
// Given an array of integers arr, find the sum of min(b), where b ranges over every (contiguous)
// subarray of arr. Since the answer may be large, return the answer modulo 109 + 7.
// Constraints:
//
// 1 <= arr.length <= 3 * 104
// 1 <= arr[i] <= 3 * 104
func sumSubarrayMins(arr []int) int {
	modBase := 1000000007
	sum := 0
	sz := len(arr)
	stack := util.Stack[int]{}
	for i := 0; i <= sz; i++ {
		for evaluateCondition(stack, i, arr) {
			mid, _ := stack.Pop()
			leftBoundry := -1
			if !stack.Empty() {
				leftBoundry, _ = stack.Peek()
			}
			rightBoundry := i

			count := (mid - leftBoundry) * (rightBoundry - mid) % modBase
			sum += (count * arr[mid]) % modBase
			sum = sum % modBase
		}
		stack.Push(i)
	}
	return sum
}

func evaluateCondition(stack util.Stack[int], i int, arr []int) bool {
	result := false
	if !stack.Empty() {
		peek, _ := stack.Peek()
		if i == len(arr) || arr[peek] >= arr[i] {
			result = true
		}
	}
	return result
}

// Another implementation - basic of N x N complexity of Sum of Subarray Minimum (#907).
func sumSubarrayMinsBasic(arr []int) int {
	modBase := 1000000007
	sum := 0
	sz := len(arr)
	for start := 0; start < sz; start++ {
		min := arr[start]
		for j := start; j < sz; j++ {
			if arr[j] < min {
				min = arr[j]
			}
			sum = sum + min
			sum = sum % modBase
		}
	}
	return sum
}

// LeetCode no. 1207: Unique Number of Occurrences
// Given an array of integers arr, return true if the number of occurrences
// of each value in the array is unique or false otherwise.
//
// Constraints:
//
// 1 <= arr.length <= 1000
// -1000 <= arr[i] <= 1000
func uniqueOccurrences(arr []int) bool {
	occCount := make(map[int]int)
	for _, val := range arr {
		c, ok := occCount[val]
		if !ok {
			occCount[val] = 1
		} else {
			occCount[val] = c + 1
		}
	}
	distincNumCount := len(occCount) // all repeat numbers are collapsed
	distinctOcc := make(map[int]int)
	for val, occ := range occCount {
		distinctOcc[occ] = val
	}
	return len(distinctOcc) == distincNumCount
}

// LeetCode problem no. 1010: Pairs of Songs With Total Durations Divisible by 60
// https://leetcode.com/problems/pairs-of-songs-with-total-durations-divisible-by-60/description/
// You are given a list of songs where the ith song has a duration of time[i] seconds.
//
// Return the number of pairs of songs for which their total duration in seconds is divisible by 60.
// Formally, we want the number of indices i, j such that i < j with (time[i] + time[j]) % 60 == 0.
// Constraints:
//
// 1 <= time.length <= 6 * 104
// 1 <= time[i] <= 500
func numPairsDivisibleBy60(time []int) int {
	modMap := make(map[int][]int)
	for _, v := range time {
		m := v % 60
		c, _ := modMap[m]
		c = append(c, v)
		modMap[m] = c
	}
	count := 0
	for m, p := range modMap {
		r := 60 - m
		pr, ok := modMap[r]
		if ok {
			if m != r {
				count += len(p) * len(pr)
			} else {
				count += len(p) * (len(pr) - 1)
			}
		}
	}
	p0, ok := modMap[0]
	if ok {
		count += len(p0) * (len(p0) - 1)
	}
	return count / 2
}

// LeetCode problem no. 152: Maximum Product Subarray
// https://leetcode.com/problems/maximum-product-subarray/description/
// Given an integer array nums, find a subarray
// that has the largest product, and return the product.
//
// The test cases are generated so that the answer will fit in a 32-bit integer.
// Constraints:
//
// 1 <= nums.length <= 2 * 104
// -10 <= nums[i] <= 10
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
func maxProduct(nums []int) int {
	result := nums[0]
	minSoFar := result
	maxSoFar := result
	for i := 1; i < len(nums); i++ {
		minSoFar, maxSoFar = filter(nums[i], minSoFar, maxSoFar)
		if maxSoFar > result {
			result = maxSoFar
		}
	}
	return result
}

func filter(crt, minSoFar, maxSoFar int) (int, int) {
	mn := minSoFar * crt
	mx := maxSoFar * crt
	n, x := crt, crt
	if mn < n {
		n = mn
	}
	if mx < n {
		n = mx
	}
	if mx > x {
		x = mx
	}
	if mn > x {
		x = mn
	}
	return n, x
}

// LeetCode problem no. 120: Triangle
// https://leetcode.com/problems/triangle/description/
//
// Given a triangle array, return the minimum path sum from top to bottom.
//
// For each step, you may move to an adjacent number of the row below.
// More formally, if you are on index i on the current row, you may move to
// either index i or index i + 1 on the next row.
//
// Constraints:
//
// 1 <= triangle.length <= 200
// triangle[0].length == 1
// triangle[i].length == triangle[i - 1].length + 1
// -104 <= triangle[i][j] <= 104
func minimumTotal(triangle [][]int) int {
	runningSum := make([]int, 1)
	runningSum[0] = triangle[0][0]
	triangleHeight := len(triangle)
	for i := 1; i < triangleHeight; i++ {
		row := triangle[i]
		numInRow := len(row)
		numInRunningSum := len(runningSum)
		newRunningSum := make([]int, 0)
		for j := 0; j < numInRow; j++ {
			val := triangle[i][j]
			if j == 0 {
				newRunningSum = append(newRunningSum, runningSum[0]+val)
			} else {
				m1 := runningSum[j-1] + val
				if j < numInRunningSum {
					m2 := runningSum[j] + val
					if m1 < m2 {
						newRunningSum = append(newRunningSum, m1)
					} else {
						newRunningSum = append(newRunningSum, m2)
					}
				} else {
					// we to append
					newRunningSum = append(newRunningSum, m1)
				}
			}
		}
		runningSum = newRunningSum
	}
	result := runningSum[0]
	for k := 1; k < len(runningSum); k++ {
		if runningSum[k] < result {
			result = runningSum[k]
		}
	}
	return result
}

func main() {

	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	fmt.Println(minimumTotal(triangle))

	//fmt.Println(maxProduct([]int{2, 3, -2, 4}))

	//fmt.Println(numPairsDivisibleBy60([]int{174, 188, 377, 437, 54, 498, 455, 239, 183, 347, 59, 199, 52, 488, 147, 82}))
	//fmt.Println(numPairsDivisibleBy60([]int{30, 20, 150, 100, 40}))
	//fmt.Println(numPairsDivisibleBy60([]int{60, 60, 60}))

	//fruits := []int{1, 2, 1}
	//fmt.Println(totalFruit(fruits))
	//
	//fruits = []int{0, 1, 2, 2}
	//fmt.Println(totalFruit(fruits))
	//
	//fruits = []int{1, 2, 3, 2, 2}
	//fmt.Println(totalFruit(fruits))

	//nums := []int{-1}
	//result := findMissingRanges(nums, -2, -1)
	//fmt.Println(result)

	//nums = []int{0, 1, 3, 50, 75}
	//result = findMissingRanges(nums, 0, 99)
	//fmt.Println(result)

	//nums := []int{1, 0, -1, 0, -2, 2}
	//fmt.Println(fourSum(nums, 0))

	//nums := []int{2, 2, 2, 2, 2}
	//fmt.Println(fourSum(nums, 8))

	//nums := []int{1, 2, 3}
	//nextPermutation(nums)
	//fmt.Println(nums)
	//
	//nums = []int{3, 2, 1}
	//nextPermutation(nums)
	//fmt.Println(nums)
	//
	//nums = []int{1, 1, 5}
	//nextPermutation(nums)
	//fmt.Println(nums)
	//
	//nums = []int{6, 9, 8, 5}
	//nextPermutation(nums)
	//fmt.Println(nums)
	//
	//nums = []int{4, 9, 8, 5}
	//nextPermutation(nums)
	//fmt.Println(nums)
	//
	//nums = []int{2, 3, 1, 3, 3}
	//nextPermutation(nums)
	//fmt.Println(nums)

	//arr := []int{3, 1, 2, 4}
	//fmt.Println(sumSubarrayMins(arr))
	//
	//arr = []int{11, 81, 94, 43, 3}
	//fmt.Println(sumSubarrayMins(arr))
}
