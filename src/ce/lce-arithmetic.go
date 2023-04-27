/*
 * Copyright (c) 2020. Neosemantix, Inc.
 * Author: Umesh Patil
 */

package main

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
