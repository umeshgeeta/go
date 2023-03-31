/*
 * https://leetcode.com/problems/bitwise-and-of-numbers-range/description/
 *
 * Given two integers left and right that represent the range [left, right],
 * return the bitwise AND of all numbers in this range, inclusive.
 *
 * Author: Umesh Patil
 */

package main

import "fmt"

func rangeBitwiseAnd(left int, right int) int {
	if left == right {
		return left
	}
	if right == left+1 {
		return left & right
	}
	mask := left
	mp := maxPowerOfTwo(left, right)
	mask = zeroOutLsb(mask, mp)
	return left & right & mask
}

func maxPowerOfTwo(left int, right int) int {
	i := 0
	p := 1
	for left+p < right {
		p = p * 2
		i++
	}
	return i
}

func zeroOutLsb(num int, k int) int {
	var mask int = (1 << k) - 1 // Creates a mask with k 1's in its most significant bits and 0's in its least significant bits
	num &= ^mask                // Sets the k least significant bits of num to zero using the mask
	return num
}

func main() {
	fmt.Println(rangeBitwiseAnd(12, 15))   // Expected output: 12
	fmt.Println(rangeBitwiseAnd(4, 7))     // Expected output: 4
	fmt.Println(rangeBitwiseAnd(21, 24))   // Expected output: 16
	fmt.Println(rangeBitwiseAnd(5, 7))     // Expected output: 4
	fmt.Println(rangeBitwiseAnd(0, 1))     // Expected output: 0
	fmt.Println(rangeBitwiseAnd(1, 3))     // Expected output: 0
	fmt.Println(rangeBitwiseAnd(162, 195)) // Expected output: 128
	fmt.Println(rangeBitwiseAnd(162, 175)) // Expected output: 160
	fmt.Println(rangeBitwiseAnd(11, 12))   // Expected output: 8
}
