// MIT License
// Author: Umesh Patil, Neosemantix, Inc.

package main

import (
	"fmt"
	"strconv"
)

// LeetCode problem no. 93: Restore IP Addresses
// https://leetcode.com/problems/restore-ip-addresses/description/
//
// A valid IP address consists of exactly four integers separated by single dots. Each integer is between 0 and 255
// (inclusive) and cannot have leading zeros.
//
// For example, "0.1.2.201" and "192.168.1.1" are valid IP addresses, but "0.011.255.245", "192.168.1.312" and
// "192.168@1.1" are invalid IP addresses.
// Given a string s containing only digits, return all possible valid IP addresses that can be formed by inserting
// dots into s. You are not allowed to reorder or remove any digits in s. You may return the valid IP addresses
// in any order.
//
// Constraints:
// 1 <= s.length <= 20
// s consists of digits only.
func restoreIpAddresses(s string) []string {
	slen := len(s)
	if slen < 4 || slen > 12 {
		return []string{}
	}
	var digits []int
	for _, r := range s {
		adigit, _ := strconv.Atoi(string(r))
		// since it is given that all are digits, we ignore the error
		digits = append(digits, adigit)
	}
	dlen := len(digits)
	var result []string
	for fd := 1; fd < 4; fd++ {
		vf, fnum := validatePart(digits[:fd])
		if vf {
			for sd := 1; sd < 4; sd++ {
				if fd+sd > dlen {
					continue
				}
				sf, snum := validatePart(digits[fd : fd+sd])
				if sf {
					for td := 1; td < 4; td++ {
						if fd+sd+td > dlen || dlen-(fd+sd+td) < 1 {
							continue
						}
						tf, tnum := validatePart(digits[fd+sd : fd+sd+td])
						ff, ffnum := validatePart(digits[fd+sd+td:])
						if tf && ff {
							s := strconv.Itoa(fnum) + "." +
								strconv.Itoa(snum) + "." +
								strconv.Itoa(tnum) + "." +
								strconv.Itoa(ffnum)
							result = append(result, s)
						}
					}
				}

			}
		}
	}
	return result
}

func validatePart(ds []int) (bool, int) {
	dl := len(ds)
	if dl > 0 && dl < 4 {
		if ds[0] == 0 {
			if dl == 1 {
				return true, 0
			}
			return false, -1 // else it is leading '0' case...
		}
		num := buildNum(ds)
		if num > -1 && num < 256 {
			return true, num
		}
	}
	return false, -1
}

func buildNum(ds []int) int {
	num := 0
	for _, d := range ds {
		num = num*10 + d
	}
	return num
}

func main() {
	s := "25525511135"
	fmt.Println(restoreIpAddresses(s))

	s = "0000"
	fmt.Println(restoreIpAddresses(s))

	s = "101023"
	fmt.Println(restoreIpAddresses(s))
}
