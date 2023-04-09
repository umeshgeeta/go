/*
 * https://leetcode.com/problems/add-binary/description/
 *
 * Given two binary strings a and b, return their sum as a binary string.
 *
 * You can use strconv function to convert into int and back to a string. But the issue with that implementation is
 * it can work as long as the string can be accommodated as int64. If the string is bigger than that, that method
 * does not work in which case you can do addition character by character.
 *
 * In Go string is collection of UTF-8 characters. There is NO char type in Go. UTF-8 character may need more than
 * one byte for a single character, say a Korean character. So when you treat string as a byte array, single individual
 * byte may represent a char or only a part of it. When one says str[i] it returns ith byte or uint8 value at the ith
 * place which is ASCI value. When one single byte is adequate to represent the character, that byte ASCI value is
 * good enough like 48 for int '0' and 49 for int '1'. Hence, we substract 48 to do the bitwise arithmetic.
 *
 * Copyright (c) 2020. Neosemantix, Inc.
 * Author: Umesh Patil
 */

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("%s\n", addBinaryAsInt("11", "1"))
	fmt.Printf("%s\n", addBinaryAsInt("1010", "1011"))

	fmt.Printf("%s\n", addBinary("0", "0"))
	fmt.Printf("%s\n", addBinary("11", "1"))
	fmt.Printf("%s\n", addBinary("1010", "1011"))
}

func addBinaryAsInt(a string, b string) string {
	binaryNuma, _ := strconv.ParseInt(a, 2, 64)
	binaryNumb, _ := strconv.ParseInt(b, 2, 64)
	return strconv.FormatInt(binaryNuma+binaryNumb, 2)
}

func addBinary(a string, b string) string {
	alen := len(a)
	blen := len(b)
	if alen == 0 {
		return b
	}
	if blen == 0 {
		return a
	}
	len := alen
	if len < blen {
		len = blen
	}
	p := len + 1
	result := make([]byte, p)
	i := 0
	carry := uint8(0)
	for i < p {
		abit := getBit(a, i, alen)
		bbit := getBit(b, i, blen)
		sum := abit + bbit + carry
		var c uint8 = 0
		switch sum {
		case 3:
			carry = 1
			c = 1
		case 2:
			carry = 1
			c = 0
		case 1:
			carry = 0
			c = 1
		case 0:
			carry = 0
			c = 0
		}
		result[p-i-1] = 48 + c
		i++
	}
	if result[0] == 48 {
		return string(result[1:])
	}
	return string(result)
}

func getBit(x string, i int, xl int) byte {
	j := xl - i - 1
	if j < 0 {
		return 0
	}
	return x[j] - 48
}
