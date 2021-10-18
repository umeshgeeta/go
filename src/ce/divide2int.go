/*
 * MIT License
 * Copyright (c) 2021. Neosemantix, Inc.
 * Author: Umesh Patil
 */

//Given two integers dividend and divisor, divide two integers without using multiplication, division, and mod operator.
//
//Return the quotient after dividing dividend by divisor.
//
//The integer division should truncate toward zero, which means losing its fractional part.
//For example, truncate(8.345) = 8 and truncate(-2.7335) = -2.
//
//Note: Assume we are dealing with an environment that could only store integers within the 32-bit signed integer
//range: [−231, 231 − 1]. For this problem, assume that your function returns 231 − 1 when the division result overflows.

package main

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if divisor == 0 {
		return math.MaxInt32
	}
	dvnd := abs(dividend)
	dvsr := abs(divisor)
	remainder := dvnd
	quotient := 0
	for remainder > 0 {
		remainder = remainder - dvsr
		quotient++
	}
	if remainder == 0 {
		return adjustSign(dividend, divisor, quotient)
	}
	if quotient > 0 {
		quotient = quotient - 1
		return adjustSign(dividend, divisor, quotient)
	}
	return quotient
}

func adjustSign(nd int, sr int, q int) int {
	r := q
	if oppositeSign(nd, sr) {
		r = 0 - q
	}
	if r < math.MinInt32 {
		r = math.MinInt32
	}
	if r > math.MaxInt32 {
		r = math.MaxInt32
	}
	return r
}

func abs(n int) int {
	if n < 0 {
		return 0 - n
	}
	return n
}

func oppositeSign(nd int, sr int) bool {
	result := false
	if nd > 0 {
		if sr < 0 {
			result = true
		}
	} else {
		if sr > 0 {
			result = true
		}
	}
	return result
}

func printAnswers(divident int, divisor int) {
	ans := divide(divident, divisor)
	fmt.Printf("%d / %d = %d \n", divident, divisor, ans)
}

func main() {

	printAnswers(-2147483648, -1)
	printAnswers(10, 3)
	printAnswers(7, -3)
	printAnswers(0, -3)
	printAnswers(-2457, -31)
	printAnswers(-8692457, 231)
	printAnswers(56, 0)
}
