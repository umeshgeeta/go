/*
 * Copyright (c) 2020. Neosemantix, Inc.
 * Author: Umesh Patil
 */

package main

import "fmt"

func main() {
	var strs string = "()"
	p := isValid(strs)
	fmt.Printf("%t", p)
}

func isValid(s string) bool {
	l := len(s)
	halfL := (l / 2) + 1
	stack := make([]rune, halfL)
	valid := true
	i := 0
	top := 0
	var r rune
	for i < l && valid && top < halfL {
		r = rune(s[i])
		if r == '(' || r == '[' || r == '{' {
			stack[top] = r
			top++
		} else {
			if top == 0 {
				// the very first character is not an opening bracket of any type
				return false
			}
			if (r == ')' && stack[top-1] == '(') ||
				(r == ']' && stack[top-1] == '[') ||
				(r == '}' && stack[top-1] == '{') {
				top--
				stack[top] = ' '
			} else {
				valid = false
			}
		}
		i++
	}
	if top > 0 || stack[0] != ' ' {
		valid = false
	}
	return valid
}
