/*
 * Copyright (c) 2020. Neosemantix, Inc.
 * Author: Umesh Patil
 */

package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	first := strs[0]
	fl := len(first)
	strCount := len(strs)
	i := 1
	match := true
	prefix := ""
	for i < fl+1 && match {
		prefix = first[0:i]
		for j := 1; j < strCount && match; j++ {
			if len(strs[j]) < len(prefix) {
				match = false
			} else {
				s := strs[j][0:i]
				if prefix != s {
					match = false
				}
			}
		}
		i++
	}
	l := len(prefix)
	if !match && l > 0 {
		prefix = prefix[0 : l-1]
	}
	return prefix
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	p := longestCommonPrefix(strs)
	fmt.Printf("%s", p)
}
