/*
 * MIT License
 * Copyright (c) 2023. Neosemantix, Inc.
 * Author: Umesh Patil
 */

// MIT License
// Author: Umesh Patil, Neosemantix, Inc.

package main

import (
	"fmt"
	"math/big"
)

func main() {
	n := 45
	fmt.Printf("climbStairs: %d for n: %d\n", climbStairs(n), n)
}

func climbStairs(n int) int {
	max2Steps := n / 2
	twoSteps := 0
	//numWays := 0
	numWays := big.NewInt(0)
	for twoSteps <= max2Steps {
		oneSteps := n - (twoSteps * 2)
		//p := comb(twoSteps, oneSteps+twoSteps)
		p := nCr(int64(oneSteps+twoSteps), int64(twoSteps))
		fmt.Printf("twoSteps: %d oneSteps: %d comb: %d\n", twoSteps, oneSteps, p)
		//numWays += p
		numWays.Add(numWays, p)
		twoSteps++
	}
	fmt.Println(numWays)
	return int(numWays.Int64())
}

func factorial(n int) int {
	i := 1
	result := 1
	for i <= n {
		result = result * i
		i++
	}
	return result
}

func factorialDiv(t int, n int) int {
	i := 0
	result := 1
	for i < t {
		result = result * (n - i)
		i++
	}
	return result
}

func comb(t int, n int) int {
	nm := factorialDiv(t, n)
	tn := factorial(t)
	fmt.Printf("nm: %d tn: %d\n", nm, tn)
	return nm / tn
}

func nCr(n, r int64) *big.Int {
	minVal := n - r
	if r < minVal {
		minVal = r
	}
	if minVal == 0 {
		return big.NewInt(1)
	}

	result := big.NewInt(1)
	for i := int64(1); i <= minVal; i++ {
		numerator := big.NewInt(n - i + 1)
		denominator := big.NewInt(i)
		factor := new(big.Int).Div(numerator, denominator)
		result.Mul(result, factor)
	}

	return result
}
