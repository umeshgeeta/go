/*
 * Copyright (c) 2020. Neosemantix, Inc.
 * Author: Umesh Patil
 */

package main

import (
	"fmt"
	"math"
)

/*
 * LeetCode no. 2591: Distribute Money to Maximum Children
 * https://leetcode.com/problems/distribute-money-to-maximum-children/description/
 *
 * You are given an integer money denoting the amount of money (in dollars) that
 * you have and another integer children denoting the number of children that you
 * must distribute the money to.
 *
 * You have to distribute the money according to the following rules:
 *
 * - All money must be distributed.
 * - Everyone must receive at least 1 dollar.
 * - Nobody receives 4 dollars.
 *
 * Return the maximum number of children who may receive exactly 8 dollars if you
 * distribute the money according to the aforementioned rules. If there is no way
 * to distribute the money, return -1.
 *
 */
func distMoney(money int, children int) int {
	if money < children {
		return -1
	}
	if money == children {
		// everyone gets $1, but no one gets $8
		return 0
	}
	// children count is given as 2 or more and less than equal to 30
	// and we know that there is more money than children, so someone
	// will get at least $2 and more
	left := money - children
	additionalFor8Max := children * 7
	if left < additionalFor8Max {
		// the amount left does not cover everyone getting 8
		max8 := left / 7
		remainder := left % 7
		non8 := children - max8

		if remainder == 3 {
			switch non8 {
			case 0:
				// everyone got 8 we are still left with $3
				// someone takes that
				return max8 - 1
			case 1:
				// if the non $8 child gets these $3, it becomes
				// $4 so some one else from max8 needs to get a
				// amount different than 8
				return max8 - 1
			default:
				// if there are 2 or more without maxed out 8
				// we can distribute remaining $3 such than none
				// of those 2 or more will get $4.
				return max8
			}
		} else {
			// remainder is not $3, so there is no danger of
			// saddling any one child with $4
			switch non8 {
			case 0:
				if remainder == 0 {
					return max8
				} else {
					// some one has to take balance so as
					// not be $8
					return max8 - 1
				}
			default:
				// there is at least one child with not 8 and
				// will not be forced to have $3; so she can have
				// all the remaining
				return max8
			}
		}
	} else if left == additionalFor8Max {
		// everyone gets 8
		return children
	} else {
		// we have more money than everyone getting 8
		// at least one children has to amass all the remaining money
		return children - 1
	}
	return -1
}

/*
 * LeetCode problem no. 319: Bulb switcher
 * https://leetcode.com/problems/bulb-switcher/description/
 *
 * There are n bulbs that are initially off. You first turn on all the bulbs,
 * then you turn off every second bulb.
 *
 * On the third round, you toggle every third bulb (turning on if it's off or turning off
 * if it's on). For the ith round, you toggle every i bulb. For the nth round,
 * you only toggle the last bulb.
 *
 * Return the number of bulbs that are on after n rounds.
 */
func bulbSwitch(n int) int {

	// As all the bulbs are initially off, at the end only bulbs that are toggled an
	// odd number of times will remain on. Whenever we are at a round i we know we toggle
	// all bulbs having a factor i. Thus, we need to find the bulbs which have an odd number
	// of factors, as those bulbs will be toggled an odd number of times (once by each factor).
	//
	// A perfect square number has an odd number of factors, since any number's factors come in
	// pairs of two different numbers, but the square root of the number will be paired with itself.
	// Thus we just need to find how many numbers from 1 to n are perfect squares. The floor of
	// the square root of n gives us the largest number whose square is less than or equal to n.
	// Hence, sqrt(n) is our answer to this problem.
	return int(math.Floor(math.Sqrt(float64(n))))
}

// LeetCode problem no. 1137: N-th Tribonacci Number
// https://leetcode.com/problems/n-th-tribonacci-number/description/
//
// The Tribonacci sequence Tn is defined as follows:
//
// T0 = 0, T1 = 1, T2 = 1, and Tn+3 = Tn + Tn+1 + Tn+2 for n >= 0.
//
// Given n, return the value of Tn.
// Constraints:
//
// 0 <= n <= 37
// The answer is guaranteed to fit within a 32-bit integer, ie. answer <= 2^31 - 1.
func tribonacci(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 1
	default:
		tim3 := 0
		tim2 := 1
		tim1 := 1
		ti := 0
		i := 3
		for i <= n {
			ti = tim1 + tim2 + tim3
			tim3 = tim2
			tim2 = tim1
			tim1 = ti
			i++
		}
		return ti
	}
}

func main() {
	fmt.Println(tribonacci(4))
	fmt.Println(tribonacci(25))
}
