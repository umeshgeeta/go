/*
 * Copyright (c) 2020. Neosemantix, Inc.
 * Author: Umesh Patil
 */

package main

import (
	"fmt"
)

func main() {
	num := 3
	fmt.Printf("%s\n", intToRoman(num))

	num = 58
	fmt.Printf("%s\n", intToRoman(num))

	num = 1994
	fmt.Printf("%s\n", intToRoman(num))
}

func intToRoman(num int) string {
	roman := ""
	valLeft := num
	if valLeft >= 1000 {
		// max num is 3999, so no more than 3 Ms
		mCount := valLeft / 1000
		for i := 0; i < mCount; i++ {
			roman += "M"
		}
		valLeft -= mCount * 1000
	}
	if valLeft >= 900 {
		roman += "CM"
		valLeft -= 900
	}
	if valLeft >= 500 {
		roman += "D"
		valLeft -= 500
	}
	if valLeft >= 400 {
		roman += "CD"
		valLeft -= 400
	}
	if valLeft >= 100 {
		cCount := valLeft / 100
		for i := 0; i < cCount; i++ {
			roman += "C"
		}
		valLeft -= cCount * 100
	}
	if valLeft >= 90 {
		roman += "XC"
		valLeft -= 90
	}
	if valLeft >= 50 {
		roman += "L"
		valLeft -= 50
	}
	if valLeft >= 40 {
		roman += "XL"
		valLeft -= 40
	}
	if valLeft >= 10 {
		xCount := valLeft / 10
		for i := 0; i < xCount; i++ {
			roman += "X"
		}
		valLeft -= xCount * 10
	}
	if valLeft == 9 {
		roman += "IX"
		valLeft = 0
	}
	if valLeft >= 5 {
		roman += "V"
		valLeft -= 5
	}
	if valLeft == 4 {
		roman += "IV"
		valLeft = 0
	}
	for valLeft > 0 {
		roman += "I"
		valLeft -= 1
	}
	return roman
}
