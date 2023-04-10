/*
 * Copyright (c) 2020. Neosemantix, Inc.
 * Author: Umesh Patil
 */

package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%d\n", mySqrt(4))
	fmt.Printf("%d\n", mySqrt(18))
	fmt.Printf("%d\n", mySqrt(10))
	fmt.Printf("%d\n", mySqrt(16))
	fmt.Printf("%d\n", mySqrt(20))
	fmt.Printf("%d\n", mySqrt(24))
	fmt.Printf("%d\n", mySqrt(27))
	fmt.Printf("%d\n", mySqrt(36))
	fmt.Printf("%d\n", mySqrt(42))
	fmt.Printf("%d\n", mySqrt(49))
	fmt.Printf("%d\n", mySqrt(61))
	fmt.Printf("%d\n", mySqrt(77))
	fmt.Printf("%d\n", mySqrt(91))
	fmt.Printf("%d\n", mySqrt(110))
}

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	if x < 4 {
		return 1
	}
	if x < 9 {
		return 2
	}
	if x < 16 {
		return 3
	}
	halfPower := int(powersOfTwo(x) / 2)
	start := halfPower * halfPower
	sqr := start * start
	//fmt.Printf("%d %d %d\n", halfPower, start, sqr)
	for sqr > x {
		start--
		sqr = start * start
	}
	for sqr <= x {
		start++
		sqr = start * start
	}
	return start - 1
}

func powersOfTwo(x int) int {
	power := 0
	val := 1
	for val <= x {
		val = val * 2
		power++
	}
	return power
}
