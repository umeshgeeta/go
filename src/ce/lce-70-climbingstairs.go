/*
 * https://leetcode.com/problems/climbing-stairs/description/
 *
 * You are climbing a staircase. It takes n steps to reach the top.
 * Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
 *
 */

// MIT License
// Author: Umesh Patil, Neosemantix, Inc.

package main

import (
	"fmt"
	"math/big"
)

func main() {
	n := 6
	fmt.Printf("climbStairs: %d for n: %d\n", climbStairs(n), n)

	n = 45
	fmt.Printf("climbStairs: %d for n: %d\n", climbStairs(n), n)
}

func climbStairs(n int) int {
	max2Steps := n / 2
	twoSteps := 0
	numWays := big.NewInt(0)
	for twoSteps <= max2Steps {
		oneSteps := n - (twoSteps * 2)
		p := nCr(int64(oneSteps+twoSteps), int64(twoSteps))
		fmt.Printf("twoSteps: %d oneSteps: %d nCr: %d\n", twoSteps, oneSteps, p)
		numWays.Add(numWays, p)
		twoSteps++
	}
	fmt.Println(numWays)
	return int(numWays.Int64())
}

func nCr(n, r int64) *big.Int {
	minVal := n - r
	if r < minVal {
		minVal = r
	}
	if minVal == 0 {
		return big.NewInt(1)
	}

	numerator := big.NewInt(1)
	denominator := big.NewInt(1)
	for i := int64(1); i <= minVal; i++ {
		n := big.NewInt(n - i + 1)
		numerator.Mul(numerator, n)
		d := big.NewInt(i)
		denominator.Mul(denominator, d)
	}
	result := new(big.Int).Div(numerator, denominator)
	fmt.Printf("%d %d %d\n", numerator, denominator, result)
	return result
}
